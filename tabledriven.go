package main
import (
	"os"
	"fmt"
	"io/ioutil"
)

const CRCPOLY = 0xedb88320
const INITXOR = 0xffffffff
const FINALXOR = 0xffffffff

func main() {

	buffer, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("buffer len = %d\n", len(buffer))

	make_crc_table()

	crc := crc32_tabledriven(buffer)

	fmt.Printf("%d\t%08x\n", len(buffer), crc)

	os.Exit(0)
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

