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

	crc := crc32_bitoriented(buffer, len(buffer))

	fmt.Printf("%d\t%08x\n", len(buffer), crc)

	os.Exit(0)
}

func crc32_bitoriented(buffer []byte, length int) (uint32) {
	var crcreg uint32 = INITXOR

	for _, c := range buffer {
		b := uint32(c)
		for i := 0; i < 8; i++ {
			if ((crcreg^b) & 1) != 0 {
				crcreg = (crcreg >> 1) ^ CRCPOLY
			} else {
				crcreg >>= 1
			}
			b >>= 1
		}
	}

	return crcreg ^ FINALXOR
}
