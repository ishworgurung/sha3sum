package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

const (
	stdinType = iota
	fileType
	binary   = "./builds/usr/bin/sha3sum"
	testFile = "./little_fox"
)

var (
	tData = map[int]string{
		224: "a33ba86cd6898df92126e3c6a5a8afe8bb90c39a63937637734e47e5",
		256: "bfd4491e0b904d781c2ef9c446bfa4504935fff746a253f68dd44b03fe1ab464",
		384: "7bc2a58662ff2bfc2f78f2858579c71308de082229aefbb71520b6898b9cf9a58822316402bf467febb909fe8662f3d9",
		512: "24766877e14d07ade089295ad091e1c6ebcdfbd5362e4baf9a5c8819e4329937a62d9c1c13fbfa47b43513ca9a04e1f20184547f95c4035f69caaed2002eb482",
	}
	nBits = []int{224, 256, 384, 512}
)

func cliHelper(nbits int, in string, inType int) error {
	nBitArg := fmt.Sprintf("-bits=%d", nbits)
	var cmd *exec.Cmd
	if inType == fileType {
		fileArg := fmt.Sprintf("-file=%s", in)
		cmd = exec.Command(binary, nBitArg, fileArg)
	} else {
		cmd = exec.Command(binary, nBitArg, "-file=-")
		cmd.Stdin = strings.NewReader(in)
	}
	cmdOut, err := cmd.Output()
	if err != nil {
		return err
	}
	if len(cmdOut) == 0 {
		return errors.New("empty command output")
	}
	got := strings.Split(string(cmdOut), " ")[0]
	algorithm := fmt.Sprintf("SHA-3(%d)", nbits)
	if got != tData[nbits] {
		errs := fmt.Sprintf("%s expected: '%s', got: '%s'",
			algorithm, tData[nbits], got)
		return errors.New(errs)
	}
	return nil
}

func TestSha3SumStdin(t *testing.T) {
	for _, b := range nBits {
		t.Logf("running %d bits\n", b)
		if err := cliHelper(b, "Little fox jumped over the tree (or something like that).", stdinType); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSha3SumFile(t *testing.T) {
	for _, b := range nBits {
		t.Logf("running %d bits\n", b)
		if err := cliHelper(b, testFile, fileType); err != nil {
			t.Fatal(err)
		}
	}
}
