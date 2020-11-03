package main

import (
	"fmt"
	"os"
)

//import (
//  "fmt"
//  "os"
//)

func WriteORFsToFile(orfs []ORF, outFilename string) {
	outfile, err := os.Create(outFilename)

	if err != nil {
		panic("Error: couldn't create file.")
	}

	//range over ORFs and write each one to file in a nice format. That is, one ORF per line.
	for i := range orfs {
		fmt.Fprint(outfile, orfs[i].startingPosition, orfs[i].startingPosition+orfs[i].length)
		// let's also indicate the strand
		if orfs[i].revComp == true {
			fmt.Fprint(outfile, "-")
		} else {
			fmt.Fprint(outfile, "+")
		}
		// print a new line
		fmt.Fprintln(outfile)
	}

	outfile.Close()
}
