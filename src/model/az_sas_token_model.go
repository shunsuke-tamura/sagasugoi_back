package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type generateSasTokenFunctionRes struct {
	SasKey string `json:"sasKey"`
	Url    string `json:"url"`
}

func GeterateSASToken() (generateSasTokenFunctionRes, error) {
	method := "GET"
	url := os.Getenv("GENERATE_SAS_TOKEN_FUNCTION_URL")
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("NewRequest err=%s", err.Error())
		return generateSasTokenFunctionRes{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Client.Do err=%s", err.Error())
		return generateSasTokenFunctionRes{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll err=%s", err.Error())
		return generateSasTokenFunctionRes{}, err
	}

	sasToken := generateSasTokenFunctionRes{}
	err = json.Unmarshal(body, &sasToken)
	if err != nil {
		log.Fatalf("json.Unmarshal err=%s", err.Error())
	}
	return sasToken, nil
}
