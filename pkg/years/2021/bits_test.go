package y2021

import "testing"

func TestParseBits(t *testing.T) {
	t.Parallel()

	input := "9C0141080250320F1802104A08"
	packets, err := ParseBits(input)
	if err != nil {
		t.Error(err)
	}

	for _, packet := range packets {
		t.Logf("%d", packet.VersionSum())
		packet.Print(0)
	}
}
