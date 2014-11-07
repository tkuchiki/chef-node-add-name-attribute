package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func DecodeJson(filename string) interface{} {
	fileByte, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	dec := json.NewDecoder(bytes.NewReader(fileByte))

	var jsonData interface{}
	dec.Decode(&jsonData)

	return jsonData
}

func EncodeJson(jsonData interface{}) []byte {
	jsonByte, err := json.Marshal(jsonData.(map[string]interface{}))

	if err != nil {
		fmt.Println(err)
	}

	var jsonIndentedBuf bytes.Buffer
	json.Indent(&jsonIndentedBuf, jsonByte, "", "    ")

	return jsonIndentedBuf.Bytes()
}
