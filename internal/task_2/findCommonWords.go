package task2

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FindCommonWords(outputFilename string, inputFilenames ...string) error{
	words := make(map[string]int)

	for _, file := range inputFilenames{
		f, rc := os.Open(file)

		if rc != nil{
			return ErrOpenFile
		}
		defer f.Close()

		var word string
		var err error

		for ; err != io.EOF;{
			_, err = fmt.Fscan(f, &word)
			if err == nil{
				words[word] += 1
			}
		}
	}

	file, rc := os.Create(outputFilename)

	if rc != nil{
		return ErrOpenFile
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for word, count := range words{
		if count == len(inputFilenames){
			writer.WriteString(word)
			writer.WriteByte('\n')
		}
	}

	writer.Flush()
	return  nil
}
