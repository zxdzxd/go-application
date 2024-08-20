// build word count (wc) tool cli

package wordcount

import (
	"flag"
	"fmt"
	"os"
)

var (
	bytecount *string
	filename  *string
)

func init() {
	bytecount = flag.String("c", "", "no of bytes in a file")
	filename = flag.String("f", "", "file name")
}

func Wordcount() {
	fmt.Println("read file")
	flag.Parse()
	// *filename = "wordcount/test.txt"
	if *filename == "" {
		fmt.Printf("Input file name using -f flag")
	}
	data, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Printf("unable to read file \nError- %s", err)
		return
	}
	

	if *bytecount == "c" {
		fmt.Println(len(data))
	} else {
		fmt.Printf("count %s\n", *bytecount)
	}

}
