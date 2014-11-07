package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GoWalk(location string) chan string {
	ch := make(chan string)
	go func() {
		err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			ch <- path
			return nil
		})

		if err != nil {
			fmt.Println(err)
		}

		defer close(ch)
	}()
	return ch
}

func WriteJSON(filepath string, jsonData []byte) {
	ioutil.WriteFile(filepath, jsonData, 0644)
}

func FileNotExists(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

func FileExists(name string) bool {
	return !FileNotExists(name)
}
