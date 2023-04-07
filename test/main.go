package main

import (
	"encoding/json"
	"fmt"
	"gin_vue_blog_AfterEnd/model/response"
	"io/ioutil"
)

type CodeMsg map[int]string

const FilePath = "model/response/ErrCode.json"

func main() {
	jsonFile, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println(err)
	}

	var codeMsg CodeMsg
	err = json.Unmarshal(jsonFile, &codeMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(codeMsg)
	fmt.Println(response.SettingsError)
}
