# Matching CRC32 of files
From:
    Reversing CRC - Theory and Practice
    HU Berlin Public Report SAR-PR-2006-05
    Martin Stigge, Henryk Plotz, Wolf Muller, Jens-Peter Redlich

## Files

* `stdfile`: 9 bytes, "123456789", CRC32 cbf43926
* `stdfile2`: 4 bytes, "asd\n", CRC32 152ddece
* `bitoriented.go` - Code from Appendix A.2, transliterated into Go
* `tabledriven.go` - Appendix A.3, transliterated from C to Go
* `crc32.go` - calculate CRC32 using standard package `hash/crc32`
* `matchfile.go` - give back bytes to append to a file to match another file's CRC32
