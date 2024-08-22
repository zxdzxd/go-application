// build word count (wc) tool cli

package wordcount

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	bytecount  *string
	filename   *string
	total_line *string
	totalword  *string
	totalchar  *string
)

func init() {
	bytecount = flag.String("c", "", "no of bytes in a file")
	filename = flag.String("f", "", "file name")
	total_line = flag.String("l", "", "no of lines in a file")
	totalword = flag.String("w", "", "no of words in a file")
	totalchar = flag.String("m", "", "no of characters in a file")
}

func Wordcount() {
	fmt.Println("read file")
	flag.Parse()
	// *filename = "wordcount/test.txt"
	if *filename == "" {
		fmt.Println("Input file name using -f flag")
	}
	data, err := os.ReadFile(*filename)
	datastring := string(data)
	if err != nil {
		fmt.Printf("unable to read file \nError- %s", err)
		return
	}
	switch {
	case *bytecount == "c":
		fmt.Println(len(data))

	case *total_line == "l":
		sliceofstring := strings.Split(datastring, "\n")

		fmt.Println(len(sliceofstring) - 1)

	case *totalword == "w":
		// fmt.Println(len(datastring))
		word_count := 0
		foundSpace := false
		for _, j := range datastring {
			if unicode.IsSpace(j) {
				foundSpace = false
			} else if !foundSpace {
				foundSpace = true
				word_count++
			}
		}
		fmt.Println(word_count)
	case *totalchar == "m":
		// char_count := 0
		total_char := strings.Split(datastring, "")

		// for _, j := range datastring {
			
		// 	// if unicode.IsSpace(j) {
		// 	// 	char_count++
		// 	// }
		// }
		fmt.Println(len(total_char))

	}

}
