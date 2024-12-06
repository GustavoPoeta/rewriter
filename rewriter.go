package main

import (
	"bufio"
	"log"
	"os"
)

// Receive the file's name and open it
func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if os.IsNotExist(err) {
		file, err = os.Create(fileName)
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

// Receives a file and reads and writes its content of the file line by line in an array (slice)
func readToSlc(file *os.File) ([]string, error) {
	scanner := bufio.NewScanner(file)

	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	var fileContent []string

	fileContent = make([]string, 0)

	for scanner.Scan() { // while there are tokens to scan
		fileContent = append(fileContent, scanner.Text()) // append it to the fileContent array
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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

// Removes the file content and rewrites it with the changes given
func writeFile(newContent []string, file *os.File) (*os.File, error) {

	// remove file content
	err := file.Truncate(0)

	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriter(file)

	// for each line in newContent, write it and go to the next line
	for _, line := range newContent {
		_, err := writer.WriteString(line + "\n")

		if err != nil {
			return nil, err
		}
	}

	// makes the changes
	err = writer.Flush()
	if err != nil {
		return nil, err
	}

	return file, nil
}
