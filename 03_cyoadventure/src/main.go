package main

// reading in files in go https://appdividend.com/2019/12/19/how-to-read-files-in-golang-go-file-read-example/
// handling JSON in go https://blog.golang.org/json

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

	m := story.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of some other type")
		}
	}
}

func readWholeFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error: ", err)
		return nil, err
	}
	return data, nil
}
