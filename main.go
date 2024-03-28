package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func createFolder(folderName string) (string, error) {

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err := os.Mkdir(folderName, 0755) // 0755 permissions allow read & execute for everyone and full control for the owner
		if err != nil {
			fmt.Printf("Failed to create directory: %s", err)
			return "", err
		}
	} else if err != nil {
		return "", err
	}
	return folderName, nil
}

func splitFileIntoChunks(filePath string, folderName string, chunkSize int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, chunkSize)
	for chunkNum := 1; ; chunkNum++ {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			return err // Handle other errors
		}

		// Generate new filename based on original name and chunk number
		newFilename := fmt.Sprintf("%s_%d%s", filepath.Base(fileInfo.Name()), chunkNum, filepath.Ext(fileInfo.Name()))
		newFilePath := filepath.Join(folderName, newFilename)
		fmt.Printf("DEBUG: newfilename: %s\n", newFilename)
		fmt.Printf("DEBUG: newfilename: %s\n", newFilePath)
		// Write the chunk to a new file
		err = writeChunkToFile(newFilePath, buffer[:bytesRead])
		if err != nil {
			return err
		}
	}
	return nil
}

// writeChunkToFile writes a chunk of data to a file specified by filePath.
func writeChunkToFile(filePath string, data []byte) error {
	chunkFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer chunkFile.Close()

	_, err = chunkFile.Write(data)
	return err
}

// need to split file into 256mb file size
// take arg with file name

func main() {
	const requiredArgs = 2

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s file-name file-size\n", os.Args[0])
		fmt.Printf("Error: This program requires %d arguments, but you provided %d.\n", requiredArgs, len(os.Args)-1)
		os.Exit(1) // Exit with a non-zero status to indicate an error
	}

	filenameArg := os.Args[1]
	folderName := strings.Split(filenameArg, ".")[0]
	fmt.Println(filenameArg)
	fileSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("File fize must be an integer: 128, 256, 512, etc.")
	}

	convertMBtoBytes := fileSize * 1024 * 1024

	folderName, err = createFolder(folderName)
	if err != nil {
		fmt.Println("Error creating folder name")
	}

	err = splitFileIntoChunks(filenameArg, folderName, convertMBtoBytes)
	if err != nil {
		fmt.Printf("Error splitting file: %v\n", err)
		return
	}

}
