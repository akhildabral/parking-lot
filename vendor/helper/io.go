package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFileToBytes(filePath string) ([]byte, error) {
	//Get config file's absolute path
	configPath, _ := filepath.Abs(filePath)

	//Read Commands config
	file, err := os.Open(configPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	return byteValue, nil
}
