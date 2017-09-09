package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"hash"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

var (
	supportedAlgorithmBits = map[int]hash.Hash{
		224: sha3.New224(),
		256: sha3.New256(),
		384: sha3.New384(),
		512: sha3.New512(),
	}
	bits = flag.Int("bits", 384, "support 224, 256, 384 and 512 bits")
	file = flag.String("file", "-", "input file")
)

//computeSha3 calculates the SHA-3 hash of buf using (n)bits of SHA-3 hashing algorithm and returns the result as a byte slice
func computeSha3(buf []byte, bits int) ([]byte, error) {
	if bits != 224 && bits != 256 && bits != 384 && bits != 512 {
		return nil, errors.New("unsupported number of bits")
	}
	if algo, ok := supportedAlgorithmBits[bits]; ok {
		_, err := algo.Write(buf)
		if err != nil {
			return nil, err
		}
		return algo.Sum(nil), nil
	}
	return nil, nil
}

func main() {
	flag.Parse()

	if *file == "-" {
		buf := bytes.Buffer{}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		sum, err := computeSha3(buf.Bytes(), *bits)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s -\n", strings.ToLower(hex.EncodeToString(sum)))
		return
	}

	fi, err := os.Stat(*file)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, fi.Size())
	c, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	n, err := c.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if int64(n) != fi.Size() {
		log.Fatal("file size mismatches read size")
	}
	sum, err := computeSha3(buf, *bits)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %s\n", strings.ToLower(hex.EncodeToString(sum)), *file)
}
