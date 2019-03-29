package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var inFile = flag.String("in", "", "input file (defaults to stdin)")
var outFile = flag.String("out", "", "output file (defaults to stdout)")

var unitSeparator = byte(0x1F)   // ascii unit separator
var recordSeparator = byte(0x1E) // ascii record separator

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *inFile == "" && len(args) == 1 {
		inFile = &args[0]
	}

	var err error
	input := os.Stdin
	output := os.Stdout

	if *inFile != "" {
		input, err = os.Open(*inFile)
		if err != nil {
			log.Fatalf("Could not open input file: %s", err)
		}
		defer input.Close()
	} else {
		info, err := os.Stdin.Stat()
		if err == nil {
			mode := info.Mode()
			if mode&os.ModeCharDevice != 0 {
				fmt.Fprintf(os.Stderr, "Reading from stdin\n")
			}
		}
	}

	r := csv.NewReader(input)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		for i, f := range rec {
			output.WriteString(f)
			if i < len(rec)-1 {
				output.Write([]byte{unitSeparator})
			}
		}

		output.Write([]byte{recordSeparator})
	}
}
