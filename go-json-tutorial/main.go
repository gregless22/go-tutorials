package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	pl = fmt.Println
)

func main() {

	jsonFile, err := os.Open("users.json")
	//handle the error
	if err != nil {
		pl(err)
	}

	pl("successfully opened file")

	defer jsonFile.Close()

	// get the byte areray
	byt, err := ioutil.ReadAll(jsonFile)

	//create the users array
	var users Users

	json.Unmarshal(byt, &users)

	pl(users)

}
