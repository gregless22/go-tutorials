package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	content := "Hello from go"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()

	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("All done with file, characters %v\n", ln)

	bytes := []byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)

	// read in the new file
	fileName := "./hello.txt"
	readContent, err := ioutil.ReadFile(fileName)
	result := string(readContent)
	fmt.Printf("Read content in %v\n", result)

	// walking a directory structure
	root, _ := filepath.Abs(".")
	fmt.Println("processing path", root)

	err = filepath.Walk(root, processPath)

	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func processPath(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
	}
	return nil
}
