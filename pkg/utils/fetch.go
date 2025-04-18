package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res, target)
	if err != nil {
		return err
	}

	return nil
}
