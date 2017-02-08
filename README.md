# Match the CRC32 value of files

To make my [Almost-Narchissist](https://github.com/bediger/Self-replicating-go) program,
I needed to be able to match the CRC32 value of a given file, the source of the
Almost-Narcissist itself.  I could only find a PHP CRC32-matcher, and it didn't
seem to work.

So I wrote a CRC32 value matcher.

I used math from:

    Reversing CRC - Theory and Practice
    HU Berlin Public Report SAR-PR-2006-05
    Martin Stigge, Henryk Plotz, Wolf Muller, Jens-Peter Redlich

I pretty much followed their appendices to get to the final matcher.

## Files

* `stdfile`: 9 bytes, "123456789", CRC32 cbf43926
* `stdfile2`: 4 bytes, "asd\n", CRC32 152ddece
* `bitoriented.go` - Code from Appendix A.2, transliterated into Go
* `tabledriven.go` - Appendix A.3, transliterated from C to Go
* `crc32.go` - calculate CRC32 using standard package `hash/crc32`
* `matchfile.go` - give back bytes to append to a file to match another file's CRC32. Appendix A.5
* `crc32.php` - calculate CRC32 using PHP's library routine, just to double check.

To create a file with the same CRC32 value as another file:

    $ go build matchfile.go
    $ ./matchfile file_to_match some_file > bytes
    $ cat bytes >> some_file
    $ go build crc32.go
    $ ./crc32 file_to_match some_file

To test the code:

    $ make test

It should build a file with a CRC32 that matches that of `stdfile`.
