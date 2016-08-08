package main
import "testing"

/**** STDIN ****/
func TestSha3sumStdin224(*testing.T) {
	// Pipe some values into a opened stdin to subprocess `sha3sum -bits 224 -file -`
}


func TestSha3sumStdin256(*testing.T) {
	// Pipe some values into a opened stdin to subprocess `sha3sum -bits 256 -file -`
}


func TestSha3sumStdin384(*testing.T) {
	// Pipe some values into a opened stdin to subprocess `sha3sum -bits 384 -file -`
}

func TestSha3sumStdin512(*testing.T) {
	// Pipe some values into a opened stdin to subprocess `sha3sum -bits 384 -file -`
}
/************/

/**** FILES ****/
func TestSha3sumFile224(*testing.T) {
	// Pipe the stdout of the subprocess `sha3sum -bits 224 -file testfile`

}

func TestSha3sumFile256(*testing.T) {
	// Pipe the stdout of the subprocess `sha3sum -bits 256 -file testfile`

}

func TestSha3sumFile384(*testing.T) {
	// Pipe the stdout of the subprocess `sha3sum -bits 384 -file testfile`

}

func TestSha3sumFile512(*testing.T) {
	// Pipe the stdout of the subprocess `sha3sum -bits 512 -file testfile`

}
/************/
