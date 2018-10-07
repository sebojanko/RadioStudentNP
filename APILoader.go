package main

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const LINK string = "http://www.radiostudent.hr/wp-admin/admin-ajax.php?action=rsplaylist_api"

func GetJSONFromLink() []byte {
	response, err := http.Get(LINK)
	if err != nil {
		panic("Have not been able to fetch data from RS.")
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Have not been able to read data from body.")
	}

	return responseData
}

func SaveJSONToDataStruct(jsonFromLink []byte) []map[string]interface{} {
	argsMap := make(map[string]interface{})

	err := json.Unmarshal(jsonFromLink, &argsMap)
	if err != nil {
		panic(err)
	}
	var songs []map[string]interface{}
	song := make(map[string]interface{})

	for _, v := range argsMap["rows"].([]interface{}) {
		for k2, v2 := range v.(map[string]interface{}) {
			song[k2] = v2.(string)
		}
		songs = append(songs, song)
		song = make(map[string]interface{})
	}

	return songs
}
