[![Build Status](https://drone.io/github.com/badbug0101/sha3sum/status.png)](https://drone.io/github.com/badbug0101/sha3sum/latest)

## About

This tool provides a sha3sum facility using golang.org/x/crypto/sha3. The details of can be found at https://godoc.org/golang.org/x/crypto/sha3. You can refer to https://godoc.org/golang.org/x/crypto/sha3 for more information and refer to Keccak web site (http://keccak.noekeon.org/) for detailed information on SHA-3.

## Usage

* Help using `-h`:

  ```
  # sha3sum -h
  Usage of sha3sum:
  -bits int
    	supports 224, 256, 384 and 512 bits. (default 384)
  -file string
    	file to perform SHA-3. (default "-")
  ```

* Calculate SHA-3 hash from stdin (default if `-file=arg` is not provided):

  ```
  # sha3sum -bits=256 -file=-
  Little fox jumped over the tree (or something like that).
  bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464 -
  OR
  # sha3sum -bits=256
  Little fox jumped over the tree (or something like that).
  bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464 -
  OR
  # echo -n "Little fox jumped over the tree (or something like that)." | ./builds/usr/bin/sha3sum -bits=256
  bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464 -
  ```

* Calculate SHA-3 hash of a file:

  ```
  # /usr/bin/sha3sum -bits=512 -file=/bin/ls
  db20c1839983506f01c03ac3f876f99b9110c9ac6dd02b48789a48d8b0e04fa3d382e15ab1afe27595f4583a190309bfd7daefb30a4dcd6e78e7dfbd43909a0e /bin/ls
  ```

## Build

To build it, install the `go` SDK from http://golang.org and run `make`. Please note that old builds are provided under `old-builds` but I don't maintain them anywhere.

## Contributing

This is a very small utility; and as such if you would like features added in, feel free to submit pull requests.
