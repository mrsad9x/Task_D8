package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("D:/sampleMessage.out")
	if err != nil {
		log.Fatalf("could not open the file: %v", err)
	}
	defer file.Close()

	log.Println("time start")

	// we just reset the offset. because we read this file once
	// imagine the cursor is in the end of the file so we have to get back to the first line and read it again

	file.Seek(0, 0)
	readWithReadLine(file)
	log.Println("time end.")
}

// Read with Readline function

func read(r *bufio.Reader) ([]byte, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return ln, err
}

func readWithReadLine(file *os.File) {
	reader := bufio.NewReader(file)
	var line []byte
	for {
		_, err := read(reader)
		line, _ = read(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("a real error happened here: %v\n", err)
		}
		//fmt.Println(string(line))
		//fmt.Println(len(line))
	}
	fmt.Println(line)

}
