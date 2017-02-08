test: matchfile crc32
	./matchfile stdfile README.md > bytes
	cat README.md bytes > matched
	./crc32 matched stdfile

matchfile: matchfile.go
	go build matchfile.go

crc32: crc32.go
	go build crc32.go
