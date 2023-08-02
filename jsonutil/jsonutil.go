package jsonutil

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kishore-tadapaneni/utils/model"
)

func WriteFile(filepath string, data []model.Employee) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, jsonData, 0644)
}
