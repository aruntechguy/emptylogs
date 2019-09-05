package main

import (
	"fmt"
	"os"
	//"path/filepath"
	//"encoding/json"
	//"io/ioutil"
)

type Fileinfo struct {
	Filepath   string `json:"filepath"`
	Backup     bool   `json:"backup"`
	Recurrence int    `json:"recurrence"`
}

type Filelist struct {
	Files []Fileinfo `json:"filelist"`
}

func main() {

	// Open input file
	jsonFile, err := os.Open("emptylogs_filelist.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened input file")

	defer jsonFile.Close()

	// Test input file

	// each file path in each routine

	// check if path is valid
	// backup before emptying? options in json file
	// empty file
}
