package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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

func EncodeJson(jsonData interface{}, filename string) []byte {
	newJson := make(map[string]interface{})

	for key, value := range jsonData.(map[string]interface{}) {
		newJson[key] = value
	}

	if _, ok := newJson["name"]; !ok {
		filenameWithoutExt := strings.Replace(filepath.Base(filename), filepath.Ext(filename), "", -1)
		newJson["name"] = filenameWithoutExt
	}

	jsonByte, err := json.Marshal(newJson)

	if err != nil {
		fmt.Println(err)
	}

	var jsonIndentedBuf bytes.Buffer
	json.Indent(&jsonIndentedBuf, jsonByte, "", "    ")

	return jsonIndentedBuf.Bytes()
}
