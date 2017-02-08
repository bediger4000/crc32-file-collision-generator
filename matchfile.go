package main
import (
	"os"
	"fmt"
	"io/ioutil"
)

const CRCPOLY = 0xedb88320
const INITXOR = 0xffffffff
const FINALXOR = 0xffffffff
const CRCINV = 0x5b358fd3

func main() {

	buffer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	make_crc_table()

	crc_to_match := crc32_tabledriven(buffer)

	fmt.Printf("%d\t%08x\n", len(buffer), crc_to_match)

	buffer, err = ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}

	crc_to_make_match := crc32_tabledriven(buffer)

	fmt.Printf("%d\t%08x\n", len(buffer), crc_to_make_match)

	bytes_to_match := fix_crc(crc_to_match, crc_to_make_match)

	fmt.Printf("Bytes to match: %08x\n", bytes_to_match)

	var i uint32
	for i = 0; i < 4; i++ {
		b := (bytes_to_match >> (i*8)) & 0xff
		fmt.Printf("%02x ", b)
		buffer = append(buffer, byte(b))
	}
	fmt.Printf("\n")

	matching_crc := crc32_tabledriven(buffer)
	fmt.Printf("%d\t%08x\n", len(buffer), matching_crc)
	
	os.Exit(0)
}

func fix_crc(crc_to_match uint32, crc_to_make_match uint32) (uint32) {
	var new_content uint32

	crc_to_match ^= FINALXOR

	for i :=0; i < 32; i++ {
		if (new_content & 1) != 0 {
			new_content = (new_content >> 1) ^ CRCPOLY
		} else {
			new_content >>= 1
		}

		if (crc_to_match & 1) != 0 {
			new_content ^= CRCINV
		}
		crc_to_match >>= 1
	}

	new_content ^= (crc_to_make_match ^ FINALXOR)

	return new_content
}

var crc_table [256]uint32

func make_crc_table() {
	var n uint32

	for n =0; n < 256; n++ {
		c := n
		for k := 0; k < 8; k++ {
			if (c & 1) != 0 {
				c = CRCPOLY ^ (c >> 1)
			} else {
				c >>= 1
			}
		}
		crc_table[n] = c
	}
}

func crc32_tabledriven(buffer []byte) (uint32) {
	var crcreg uint32 = INITXOR

	// I think range buffer makes the len(buffer) important:
	// buffer has to have values up to buffer[len(buffer)} for
	// this to be correct.
	for _, c := range buffer {
		b := uint32(c)
		crcreg = (crcreg >> 8) ^ crc_table[((crcreg ^ b) & 0xff)]
	}

	return crcreg ^ FINALXOR
}

