# Matching CRC32 of files
From:
    Reversing CRC - Theory and Practice
    HU Berlin Public Report SAR-PR-2006-05
    Martin Stigge, Henryk Plotz, Wolf Muller, Jens-Peter Redlich

## Files

* `stdfile`: 9 bytes, "123456789", CRC32 cbf43926
* `stdfile2`: 4 bytes, "asd\n", CRC32 152ddece
* `bitoriented.go` - Code from Appendix A.2, transliterated into Go
