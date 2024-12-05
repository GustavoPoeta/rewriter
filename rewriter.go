package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if err != nil {
		panic(err)
	}

	return file, nil
}

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

func modifyFileArr(fileArr []string, newLinesMap map[int]string) []string {

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
