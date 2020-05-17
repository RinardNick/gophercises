package main

// reading in files in go https://appdividend.com/2019/12/19/how-to-read-files-in-golang-go-file-read-example/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var storyPath string = "storyTemplate.json"
	storyBytes, err := readWholeFile(string(storyPath))
	if err != nil {
		log.Println("Failed to read in file: ", err)
	}

	var story interface{}
	err = json.Unmarshal(storyBytes, &story)
	if err != nil {
		log.Println("Failed to Unmarshall JSON: ", err)
	}

	fmt.Println(story.debate)
}

func readWholeFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error: ", err)
		return nil, err
	}
	return data, nil
}
