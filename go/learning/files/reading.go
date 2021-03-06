package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// Learning to read files in go.

func GetTestFilePath() string {
	// return the rooted path name of current working directory
	curDir, _ := os.Getwd()
	return path.Join(curDir, "sample.txt")
}

// BasicReading is about getting started with reading files in go.
// Reads the test file using os.ReadFile and later with os.Open
func BasicReading() {
	fmt.Println("========= Reading files ==========")

	testFile := GetTestFilePath()

	// files can be read in many different ways
	// To simply read all data in a file in one shot,
	// os.ReadFile() comes in handy. It returns the data
	// in a byte slice.
	fmt.Println("==> Reading using os.ReadFile")
	dataBytes, err := os.ReadFile(testFile)
	checkErr(err)

	// converting the byte slice into a string
	data := string(dataBytes)
	fmt.Printf("data: %v\n", data)

	fmt.Println("==> Reading using file opened with os.Open")
	file, err := os.Open(testFile)
	checkErr(err)
	// don't forget to close the file handle
	// it is idiomatic in golang to use defer right
	// after a successful open
	defer file.Close()

	// os.Read() reads up to len(buf) given to Read()
	// since we we reused the byte slice (dataBytes),
	// file.Read() read all contents of the file.
	readLength, err := file.Read(dataBytes)
	checkErr(err)
	fmt.Printf("readLength: %v\n", readLength)

	// Read the file in chunks
	fmt.Printf("Reading from %v\n", "sample2.txt")
	buf := make([]byte, 64)
	samplefile, err := os.Open("sample2.txt")
	if err != nil {
		checkErr(err)
	}
	defer samplefile.Close()

	for {
		n, err := samplefile.Read(buf)
		if err != nil {
			if err != io.EOF {
				checkErr(err)
			}
			break
		}
		// Print upto the number of bytes read.
		fmt.Printf("%s", buf[:n])
	}
	fmt.Println()
}

func ReadUsingIOUtil() {
	fmt.Println("========= Reading files using ioutil functions ==========")

	file, err := os.Open(GetTestFilePath())
	checkErr(err)
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	checkErr(err)

	data := string(rawData)
	fmt.Printf("data: %v\n", data)
}
