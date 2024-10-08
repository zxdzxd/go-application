// build word count (wc) tool cli

package wordcount

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
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
	flag.Parse()
	// *filename = "wordcount/test.txt"
	var stdata *os.File
	var filestat fs.FileInfo
	var data []byte
	var err error
	if *filename == "" {
		stdata = os.Stdin
		filestat, err = stdata.Stat()
		if err != nil {
			fmt.Printf("unable to get std fileinfo \nError- %s", err)
			return
		}
		if filestat.Size() > 0 {
			data, err = io.ReadAll(stdata)
			if err != nil {
				fmt.Printf("unable to read stdin file \nError- %s", err)
				return
			}
		}
	} else {
		data, err = os.ReadFile(*filename)
	}

	datastring := string(data)
	if err != nil {
		fmt.Printf("unable to read file \nError- %s", err)
		return
	}
	switch {
	case *bytecount == "c":
		fmt.Println(byte_count(data))

	case *total_line == "l":
		fmt.Println(totalLine(datastring))

	case *totalword == "w":
		fmt.Println(total_word(datastring))

	case *totalchar == "m":
		total_char := strings.Split(datastring, "")
		fmt.Println(len(total_char))
	default:
		fmt.Println(totalLine(datastring), total_word(datastring), byte_count(data))

	}

}

func byte_count(data []byte) int {
	return len(data)
}

func totalLine(datastring string) int {
	sliceofstring := strings.Split(datastring, "\n")
	return len(sliceofstring) - 1
}

func total_word(datastring string) int {
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
	return word_count
}
