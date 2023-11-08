package api

import (
	"io"
	"log"
	"net/http"
)

func CallFactsAPI(path string) (string, error) {
	url := "https://cat-fact.herokuapp.com/" + path
	log.Println("URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Response: ", string(body))

	return string(body), nil
}
