package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

const (
	SHA3Bin      = "./builds/usr/bin/sha3sum"
	SHA3TestFile = "./little_fox"
)

func errx(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func SHA3sumStdin(nbits int, txt string) {
	nbits_str := fmt.Sprintf("-bits=%d", nbits)
	cmd := exec.Command(SHA3Bin, nbits_str, "-file=-")
	cmd.Stdin = strings.NewReader(txt)
	got, err := cmd.Output()
	errx(err)
	want := map[int]string{
		224: "a33ba86cd6898df92126e3c6a5a8afe8bb90c39a63937637734e47e5 -",
		256: "bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464 -",
		384: "7bc2a58662ff2bfc2f78f2858579c71308de082229aefbb71520b6898b9cf9a58822316402bf467febb909fe8662f3d9 -",
		512: "24766877e14d07ade089295ad091e1c6ebcdfbd5362e4baf9a5c8819e4329937a62d9c1c13fbfa47b43513ca9a04e1f20184547f95c4035f69caaed2002eb482 -",
	}
	got_str := fmt.Sprintf("%s", got)
	got_str = got_str[:len(got_str)-1]
	algo := fmt.Sprintf("SHA-3(%d)", nbits)
	if got_str != want[nbits] {
		log.Fatalf("%s: want: '%s', got: '%s'", algo, want[nbits], got_str)
	}
}

func TestSha3sumStdin224(t *testing.T) {
	SHA3sumStdin(224, "Little fox jumped over the tree (or something like that).")
}

func TestSha3sumStdin256(t *testing.T) {
	SHA3sumStdin(256, "Little fox jumped over the tree (or something like that).")
}

func TestSha3sumStdin384(t *testing.T) {
	SHA3sumStdin(384, "Little fox jumped over the tree (or something like that).")
}

func TestSha3sumStdin512(t *testing.T) {
	SHA3sumStdin(512, "Little fox jumped over the tree (or something like that).")
}

func SHA3sumFile(nbits int, testfile string) {
	nbits_str := fmt.Sprintf("-bits=%d", nbits)
	testfile_str := fmt.Sprintf("-file=%s", testfile)
	//fmt.Println(SHA3Bin, nbits_str, testfile_str)
	cmd := exec.Command(SHA3Bin, nbits_str, testfile_str)
	got, err := cmd.Output()
	errx(err)
	want := map[int]string{
		224: fmt.Sprintf("a33ba86cd6898df92126e3c6a5a8afe8bb90c39a63937637734e47e5 %s", testfile),
		256: fmt.Sprintf("bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464 %s", testfile),
		384: fmt.Sprintf("7bc2a58662ff2bfc2f78f2858579c71308de082229aefbb71520b6898b9cf9a58822316402bf467febb909fe8662f3d9 %s", testfile),
		512: fmt.Sprintf("24766877e14d07ade089295ad091e1c6ebcdfbd5362e4baf9a5c8819e4329937a62d9c1c13fbfa47b43513ca9a04e1f20184547f95c4035f69caaed2002eb482 %s", testfile),
	}
	got_str := fmt.Sprintf("%s", got)
	got_str = got_str[:len(got_str)-1]
	algo := fmt.Sprintf("SHA-3(%d)", nbits)
	if got_str != want[nbits] {
		log.Fatalf("%s: want: '%s', got: '%s'", algo, want[nbits], got_str)
	}
}

func TestSha3sumFile224(t *testing.T) {
	SHA3sumFile(224, SHA3TestFile)
}

func TestSha3sumFile256(t *testing.T) {
	SHA3sumFile(256, SHA3TestFile)
}

func TestSha3sumFile384(t *testing.T) {
	SHA3sumFile(384, SHA3TestFile)
}

func TestSha3sumFile512(t *testing.T) {
	SHA3sumFile(512, SHA3TestFile)
}
