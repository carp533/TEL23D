package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory path as an argument.")
		return
	}

	dirPath := os.Args[1]
	// dirPath := "."

	files, err := getFileList(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(files)

	// f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println("Datei konnte nicht gelesen werden")
	// 	return
	// }

	// _, err = f.Write([]byte("Test"))
	// if err != nil {
	// 	fmt.Println("Datei konnte nicht geschrieben werden")
	// 	return
	// }

	// defer f.Close()
	// subdirs, err := getSubdirs(".")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(subdirs)

}

func getSubdirs(dirPath string) ([]string, error) {
	var subdirs []string
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subdirs = append(subdirs, file.Name())
		}
	}

	return subdirs, nil
}

func getFileList(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, file := range files {
		result = append(result, file.Name())
	}
	return result, err
}
