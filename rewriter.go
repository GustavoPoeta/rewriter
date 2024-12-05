package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Receive the file's name and open it
func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if err != nil {
		panic(err)
	}

	return file, nil
}

// Receives a file and reads and writes its content of the file line by line in an array (slice)
func readToSlc(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)

	var fileContent []string

	fileContent = make([]string, 0)

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileContent, nil
}

// Receives a map where key is the index in the slice that needs to be modified and the value is the new line
func modifyFileArr(fileArr []string, newLinesMap map[int]string) []string {

	// loops newLinesMap and if the index is valid:
	//it modifies the content of the specified line with its corresponding value
	for i, newLine := range newLinesMap {
		if i >= 0 && i < len(fileArr) {
			fileArr[i] = newLine
		}
	}

	return fileArr
}

func main() {
	file, err := openFile("dummy.txt")

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := readToSlc(file)

	if err != nil {
		log.Fatal(err)
	}

	newLinesMap := make(map[int]string)

	newLinesMap[1] = "não e você"

	modifiedArr := modifyFileArr(fileContent, newLinesMap)

	fmt.Println(modifiedArr)
}
