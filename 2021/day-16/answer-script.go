package main

import (
	"fmt"
	"os"
	"strings"
)

var Conversions map[byte]int = map[byte]int{
	'0': 0b0000,
	'1': 0b0001,
	'2': 0b0010,
	'3': 0b0011,
	'4': 0b0100,
	'5': 0b0101,
	'6': 0b0110,
	'7': 0b0111,
	'8': 0b1000,
	'9': 0b1001,
	'A': 0b1010,
	'B': 0b1011,
	'C': 0b1100,
	'D': 0b1101,
	'E': 0b1110,
	'F': 0b1111,
}

const (
	VersionLength   = 3
	TypeLength      = 3
	LiteralPacketId = 4
)

type Packet struct {
	version int
	typeId  int
	value   int
	body    []Packet
}

func getRange(start, end int, nibbles string) int {
	result := 0
	for start < end {
		result <<= 1
		nibbleIndex := start / 4
		offset := 3 - (start % 4)
		bit := (Conversions[nibbles[nibbleIndex]] >> offset) & 1
		result |= int(bit)
		start++
	}
	return result
}

func parse(start int, nibbles string) (Packet, int) {
	packet := Packet{}
	packet.version = getRange(start, start+3, nibbles)
	start += 3
	packet.typeId = getRange(start, start+3, nibbles)
	start += 3
	if packet.typeId == LiteralPacketId {
		keepReading := true
		value := 0
		for keepReading {
			value <<= 4
			next := getRange(start, start+5, nibbles)
			start += 5
			keepReading = (next & 0b10000) != 0
			value |= next & 0b1111
		}
		packet.value = value
		return packet, start
	} else {
		lengthTypeId := getRange(start, start+1, nibbles)
		start++
		if lengthTypeId == 0 {
			childrenLengthInBits := getRange(start, start+15, nibbles)
			start += 15
			childPacket, finishedAt := parse(start, nibbles)
			packet.body = append(packet.body, childPacket)
			for finishedAt-start < childrenLengthInBits {
				childPacket, finishedAt = parse(finishedAt, nibbles)
				packet.body = append(packet.body, childPacket)
			}
			start += childrenLengthInBits
			return packet, start
		} else {
			subpacketsCount := getRange(start, start+11, nibbles)
			start += 11
			finishedAt := start
			for subpacketsCount > 0 {
				childPacket, newFinishedAt := parse(finishedAt, nibbles)
				packet.body = append(packet.body, childPacket)
				finishedAt = newFinishedAt
				subpacketsCount--
			}
			return packet, finishedAt
		}
	}
}

func versionSum(packet Packet) int {
	res := packet.version
	for i := 0; i < len(packet.body); i++ {
		res += versionSum(packet.body[i])
	}
	return res
}

func eval(packet Packet) int {
	typeId := packet.typeId
	if typeId == 0 {
		sum := 0
		for i := 0; i < len(packet.body); i++ {
			sum += eval(packet.body[i])
		}
		return sum
	} else if typeId == 1 {
		prod := eval(packet.body[0])
		for i := 1; i < len(packet.body); i++ {
			prod *= eval(packet.body[i])
		}
		return prod
	} else if typeId == 2 {
		min := eval(packet.body[0])
		for i := 1; i < len(packet.body); i++ {
			res := eval(packet.body[i])
			if res < min {
				min = res
			}
		}
		return min
	} else if typeId == 3 {
		max := eval(packet.body[0])
		for i := 1; i < len(packet.body); i++ {
			res := eval(packet.body[i])
			if res > max {
				max = res
			}
		}
		return max
	} else if typeId == 4 {
		return packet.value
	} else if typeId == 5 {
		if len(packet.body) != 2 {
			panic(fmt.Errorf("greater-than packet %v doesn't have 2 children", packet))
		}
		a := eval(packet.body[0])
		b := eval(packet.body[1])
		if a > b {
			return 1
		} else {
			return 0
		}
	} else if typeId == 6 {
		if len(packet.body) != 2 {
			panic(fmt.Errorf("less-than packet %v doesn't have 2 children", packet))
		}
		a := eval(packet.body[0])
		b := eval(packet.body[1])
		if a < b {
			return 1
		} else {
			return 0
		}
	} else if typeId == 7 {
		if len(packet.body) != 2 {
			panic(fmt.Errorf("equal-to packet %v doesn't have 2 children", packet))
		}
		a := eval(packet.body[0])
		b := eval(packet.body[1])
		if a == b {
			return 1
		} else {
			return 0
		}
	} else {
		panic(fmt.Errorf("unknown packet type %d. packet: %v", typeId, packet))
	}
}

func run() int {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	nibbles := strings.TrimSpace(string(data))
	packet, _ := parse(0, nibbles)

	fmt.Println(versionSum(packet))
	fmt.Println(eval(packet))
	return 0
}

func main() {
	os.Exit(run())
}
