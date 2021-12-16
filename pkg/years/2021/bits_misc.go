package y2021

import (
	"fmt"
	"strings"
)

const (
	bitsMaskVersion = 0b11100000
	bitsMaskTypeId  = 0b00011100
)

type BitsPacket struct {
	Version bitsVersion
	TypeId  bitsTypeId

	Value  uint64
	Length int

	SubPackets []*BitsPacket
}

type bitsVersion uint8

type bitsTypeId uint8

const (
	bitsTypeIdSum     bitsTypeId = 0x00
	bitsTypeIdProduct bitsTypeId = 0x01
	bitsTypeIdMin     bitsTypeId = 0x02
	bitsTypeIdMax     bitsTypeId = 0x03
	bitsTypeIdLiteral bitsTypeId = 0x04
	bitsTypeIdGT      bitsTypeId = 0x05
	bitsTypeIdLT      bitsTypeId = 0x06
	bitsTypeIdEQ      bitsTypeId = 0x07
)

func (t bitsTypeId) IsOperator() bool {
	return t != bitsTypeIdLiteral
}

func (packet *BitsPacket) Print(depth int) {
	fmt.Printf("%s%+v\n", strings.Repeat("\t", depth), packet)
	for _, subPacket := range packet.SubPackets {
		subPacket.Print(depth + 1)
	}
}

func (packet *BitsPacket) VersionSum() int {
	sum := int(packet.Version)

	for _, subPacket := range packet.SubPackets {
		sum += subPacket.VersionSum()
	}

	return sum
}
