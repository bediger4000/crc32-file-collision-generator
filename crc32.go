package main

import (
	"fmt"
	"hash/crc32"
	"os"
)

func main() {
	var all []byte
	var count int
	buffer := make([]byte, 1024)
	for _, fname := range os.Args[1:] {
		fin, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Problem opening %q for read: %v\n",
				fname, err)
			os.Exit(1)
		}
		for {
			rdsz, err2 := fin.Read(buffer)
			if rdsz == 0 {
				break
			}
			if err2 != nil {
				fmt.Fprintf(os.Stderr, "Read problem on %s: %v\n", fname, err2)
				os.Exit(3)
			}
			count += rdsz
			all = append(all, buffer...)
		}
		fin.Close()
	}
	fmt.Printf("len(all) = %d, read %d bytes\n", len(all), count)
	var cksum uint32 = crc32.ChecksumIEEE(all[0:count])
	fmt.Printf("IEEE CRC32: %x\n", cksum)

	os.Exit(0)
}
