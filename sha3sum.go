package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3" //https://godoc.org/golang.org/x/crypto/sha3
	"hash"
	"os"
	"strings"
)

var bits = flag.Int("bits", 384, "supports 224, 256, 384 and 512 bits.")
var file = flag.String("file", "-", "file to perform SHA-3.")
const (
	ERR_READ = 1
	ERR_BITS = 2
	ERR_STAT = 3
	ERR_OPEN = 4
)

func main() {
	flag.Parse()
	supported_alg := map[int]hash.Hash{
		224: sha3.New224(),
		256: sha3.New256(),
		384: sha3.New384(),
		512: sha3.New512(),
	}
	alg := sha3.New384() // Default
	if *file == "-" {    // Stdin
		scanner := bufio.NewScanner(os.Stdin)
		var stdinBuf string
		for scanner.Scan() {
			stdinBuf += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "read stdin:", err)
			os.Exit(ERR_READ)
		}
		buf := []byte(stdinBuf)
		if val, exists := supported_alg[*bits]; exists {
			alg = val
		} else {
			fmt.Fprintln(os.Stderr, "unsupported number of bits")
			os.Exit(ERR_BITS)
		}
		alg.Write(buf)
		buf = alg.Sum(nil)
		fmt.Printf("%v -\n", strings.ToLower(hex.EncodeToString(buf)))
	} else { // File
		fi, err := os.Stat(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "stat error:", err)
			os.Exit(ERR_STAT)
		}
		buf := make([]byte, 1, fi.Size())
		c, err := os.Open(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "open error:", err)
			os.Exit(ERR_OPEN)
		}
		_, err = c.Read(buf)
		if err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
			os.Exit(ERR_READ)
		}
		if val, exists := supported_alg[*bits]; exists {
			alg = val
		} else {
			fmt.Fprintln(os.Stderr, "unsupported number of bits")
			os.Exit(ERR_BITS)
		}
		alg.Write(buf)
		buf = alg.Sum(nil)
		fmt.Printf("%v %s\n", strings.ToLower(hex.EncodeToString(buf)), *file)
	}
}
