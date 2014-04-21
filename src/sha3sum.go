// This file provides a SHA3 sum tool using go.crypto/sha3 in similar vein to other
// tools such as md5sum, shasum etc. The internals of the Keccak-f function are
// computed in go.crypto/keccakf.go used by go.crypto/sha3.
// Refer to go.crypto/sha3 reference implementation for more information and refer
// to Keccak web site (http://keccak.noekeon.org/) for detailed information on SHA3.

package main

import (
	"bufio"
	"code.google.com/p/go.crypto/sha3"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"os"
	"strings"
)

func main() {
	// User flags
	var bits = flag.Int("bits", 384, "Supports 224, 256, 384(default) and 512 bits.")
	var file = flag.String("file", "", "File to sha3sum (default is standard input)")
	var docheck = flag.Bool("check", false, "Check sha3 hashes against a file")
	var docheckfile = flag.String("check-file", "", "File to check the sha3 hashes against")

	flag.Parse()
	supported_alg := map[int]hash.Hash{
		224: sha3.NewKeccak224(),
		256: sha3.NewKeccak256(),
		384: sha3.NewKeccak384(),
		512: sha3.NewKeccak512(),
	}
	alg := sha3.NewKeccak384() // default
	if *docheck == true && *docheckfile != "" {
		_, cferr := os.Stat(*docheckfile)
		if cferr != nil {
			fmt.Fprintln(os.Stderr, "stat error:", cferr)
			os.Exit(-1)
		}
		// cfbuf := make([]byte, 1, cfile.Size())
		c, err := os.Open(*docheckfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "open error:", err)
			os.Exit(-1)
		}
		// _, err = c.Read(cfbuf)
		scanner := bufio.NewScanner(c)
		fmt.Println(scanner)
		var sha3line string
		for scanner.Scan() {
			sha3line = scanner.Text()
			sha3, sha3file := strings.Split(sha3line, " ")
			verify_sha3sum(sha3, sha3file, alg)
		}

	} else {
		if *file == "" {
			scanner := bufio.NewScanner(os.Stdin)
			var stdinBuf string
			for scanner.Scan() {
				stdinBuf += scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
				os.Exit(-1)
			}
			buf := []byte(stdinBuf)
			if val, exists := supported_alg[*bits]; exists {
				alg = val
			} else {
				fmt.Fprintln(os.Stderr, "unsupported number of bits")
				os.Exit(-1)
			}
			alg.Write(buf)
			buf = alg.Sum(nil)
			fmt.Printf("%v -\n", strings.ToLower(hex.EncodeToString(buf)))
		} else {
			fi, err := os.Stat(*file)
			if err != nil {
				fmt.Fprintln(os.Stderr, "stat error:", err)
				os.Exit(-1)
			}
			buf := make([]byte, 1, fi.Size())
			c, err := os.Open(*file)
			if err != nil {
				fmt.Fprintln(os.Stderr, "open error:", err)
				os.Exit(-1)
			}
			_, err = c.Read(buf)
			if err != nil {
				fmt.Fprintln(os.Stderr, "read error:", err)
				os.Exit(-1)
			}
			if val, exists := supported_alg[*bits]; exists {
				alg = val
			} else {
				fmt.Fprintln(os.Stderr, "unsupported number of bits")
				os.Exit(-1)
			}
			alg.Write(buf)
			buf = alg.Sum(nil)
			sha3sum(alg, *file)
			fmt.Printf("%v %s\n", strings.ToLower(hex.EncodeToString(buf)), *file)
		}
	}
}
