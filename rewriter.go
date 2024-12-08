package rewriter

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// OpenFile opens a file with the given name in read-write mode.
// If the file does not exist, it will be created.
// Returns the opened file and an error, which is nil on success.
func OpenFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if os.IsNotExist(err) {
		file, err = os.Create(fileName)
	} else if err != nil {
		return nil, err
	}

	return file, nil
}

// ReadToSlc receives a file, and creates a slice that will contain the content
// of the file, line by line. Then it starts reading and appending the lines inside the slice.
// It will return the slice and an error that, if successful, will be nil.
func ReadToSlc(file *os.File) ([]string, error) {

	fileInfo, err := os.Stat(file.Name())

	if os.IsNotExist(err) || fileInfo.Size() <= 0 {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)

	_, err = file.Seek(0, 0)
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

// ModifyFileArr receives an array containing the lines the file to modify,
// and a map with the new ones where [index of the line to modify]new line. It will loop the map and if the index is valid
// it replaces the original line with the new. It returns the same file's array,
// but modified, and an error that, if successful, is nil.
func ModifyFileArr(fileArr []string, newLinesMap map[int]string) ([]string, error) {

	if len(newLinesMap) <= 0 {
		return []string{}, fmt.Errorf("modification failed: the lenght of the given map is 0 or negative")
	}

	// loops newLinesMap and if the index is valid:
	//it modifies the content of the specified line with its corresponding value
	for i, newLine := range newLinesMap {
		if i >= 0 && i < len(fileArr) {
			fileArr[i] = newLine
		} else if i >= len(fileArr) {
			fileArr = append(fileArr, newLine)
		}
	}

	return fileArr, nil
}

// WriteFile Removes the file content and rewrites it with the changes given
// WriteFile overwrites the content of the given file with the provided lines.
// The file's existing content is truncated, and the new lines from newContent
// are written to it, each followed by a newline character.
// Changes are flushed to disk before the function returns.
// Returns the file and an error, which is nil on success.
func WriteFile(newContent []string, file *os.File) (*os.File, error) {

	fileInfo, err := os.Stat(file.Name())

	if os.IsNotExist(err) || fileInfo.Size() <= 0 {
		return nil, err
	}

	if len(newContent) <= 0 {
		return file, fmt.Errorf("file was not modificated as there aren't changes to be made")
	}

	// remove file content
	err = file.Truncate(0)
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
