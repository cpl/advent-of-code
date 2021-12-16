package y2021

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func bitsDecode(in string) string {
	var builder strings.Builder
	builder.Grow(len(in) * 4)

	for _, r := range in {
		switch r {
		case '0':
			builder.WriteString("0000")
		case '1':
			builder.WriteString("0001")
		case '2':
			builder.WriteString("0010")
		case '3':
			builder.WriteString("0011")
		case '4':
			builder.WriteString("0100")
		case '5':
			builder.WriteString("0101")
		case '6':
			builder.WriteString("0110")
		case '7':
			builder.WriteString("0111")
		case '8':
			builder.WriteString("1000")
		case '9':
			builder.WriteString("1001")
		case 'A':
			builder.WriteString("1010")
		case 'B':
			builder.WriteString("1011")
		case 'C':
			builder.WriteString("1100")
		case 'D':
			builder.WriteString("1101")
		case 'E':
			builder.WriteString("1110")
		case 'F':
			builder.WriteString("1111")
		}
	}

	return builder.String()
}

func bitsAllZero(in string) bool {
	return strings.IndexRune(in, '1') == -1
}

func ParseBits(input string) ([]*BitsPacket, error) {
	binary := bitsDecode(strings.TrimSpace(input))
	fmt.Println(binary)

	idx := 0
	max := len(binary) - 1

	var packets []*BitsPacket

	for idx < max {
		if idx > max-6 {
			if bitsAllZero(binary[idx:]) {
				break
			}

			return packets, fmt.Errorf("not enough bits to begin packet")
		}

		log.Printf("decode[%4d %4d]: version=%s", idx, len(packets), binary[idx:idx+3])
		log.Printf("decode[%4d %4d]:    type=%s", idx, len(packets), binary[idx+3:idx+6])

		packet, err := bitsParsePacket(binary, &idx)
		if err != nil {
			return packets, fmt.Errorf("error parsing packet: %w", err)
		}

		packets = append(packets, packet)
		idx++
	}

	return packets, nil
}

func bitsParsePacket(binary string, idx *int) (*BitsPacket, error) {
	ptr := *idx

	var err error

	v, _ := strconv.ParseUint(binary[ptr:ptr+3], 2, 8)
	t, _ := strconv.ParseUint(binary[ptr+3:ptr+6], 2, 8)
	ptr += 6

	packet := &BitsPacket{
		Version: bitsVersion(v),
		TypeId:  bitsTypeId(t),
	}

	switch packet.TypeId {
	case bitsTypeIdLiteral:
		var builder strings.Builder
		final := false
		for !final {
			finalBit := binary[ptr]
			final = finalBit == '0'

			builder.WriteString(binary[ptr+1 : ptr+5])
			ptr += 5
		}

		valueStr := builder.String()
		builder.Reset()

		value, err := strconv.ParseUint(valueStr, 2, 64)
		if err != nil {
			return nil, err
		}
		packet.Value = value

		log.Printf("decode[%4d]: literal=%d f=%t (%s)", ptr, value, final, valueStr)
	default:
		switch binary[ptr] {
		case '1':
			countBits := 11
			count, _ := strconv.ParseUint(binary[ptr+1:ptr+1+countBits], 2, 64)
			ptr += 1 + countBits
			log.Printf("decode[%4d]: op C=%05d", ptr, count)

			packet.SubPackets = make([]*BitsPacket, count)
			for subPacket := 0; subPacket < int(count); subPacket++ {
				packet.SubPackets[subPacket], err = bitsParsePacket(binary, &ptr)
				if err != nil {
					return nil, fmt.Errorf("failed parsing subpacket %d of %d: %w", subPacket, count, err)
				}
			}

		case '0':
			lengthBits := 15
			length, _ := strconv.ParseUint(binary[ptr+1:ptr+1+lengthBits], 2, 64)
			ptr += 1 + lengthBits
			log.Printf("decode[%4d]: op L=%05d", ptr, length)

			internalPtr := 0
			packet.SubPackets = make([]*BitsPacket, 0)
			for internalPtr < int(length) {
				subPacket, err := bitsParsePacket(binary[ptr:ptr+int(length)], &internalPtr)
				if err != nil {
					return nil, fmt.Errorf("failed parsing subpacket %d of %d: %w", internalPtr, length, err)
				}

				packet.SubPackets = append(packet.SubPackets, subPacket)
			}

			ptr += int(length)
		}

		switch packet.TypeId {
		case bitsTypeIdSum:
			for _, subPacket := range packet.SubPackets {
				packet.Value += subPacket.Value
			}
		case bitsTypeIdProduct:
			mul := uint64(1)
			for _, subPacket := range packet.SubPackets {
				mul *= subPacket.Value
			}
			packet.Value = mul
		case bitsTypeIdMin:
			min := uint64(math.MaxUint64)
			for _, subPacket := range packet.SubPackets {
				if subPacket.Value < min {
					min = subPacket.Value
				}
			}
			packet.Value = min
		case bitsTypeIdMax:
			max := uint64(0)
			for _, subPacket := range packet.SubPackets {
				if subPacket.Value > max {
					max = subPacket.Value
				}
			}
			packet.Value = max
		case bitsTypeIdGT:
			cmp0 := packet.SubPackets[0]
			cmp1 := packet.SubPackets[1]

			if cmp0.Value > cmp1.Value {
				packet.Value = 1
			}
		case bitsTypeIdLT:
			cmp0 := packet.SubPackets[0]
			cmp1 := packet.SubPackets[1]

			if cmp0.Value < cmp1.Value {
				packet.Value = 1
			}
		case bitsTypeIdEQ:
			cmp0 := packet.SubPackets[0]
			cmp1 := packet.SubPackets[1]

			if cmp0.Value == cmp1.Value {
				packet.Value = 1
			}
		default:
			return packet, fmt.Errorf("unknown packet type")
		}
	}

	*idx = ptr

	return packet, nil
}
