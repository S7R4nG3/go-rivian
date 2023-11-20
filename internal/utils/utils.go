package utils

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GraphqlQuery(headers map[string]string, url string, body []byte, log *logrus.Logger) (int, []byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		log.Fatalf("Unable to create new HTTP Request with error -- %v", err)
		return 0, []byte{}, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to make GraphQL Query request with error -- %v", err)
		return 0, []byte{}, err
	}
	defer resp.Body.Close()
	rawBody, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, rawBody, nil
}

func IsOk(statusCode int) bool {
	if statusCode >= 200 && statusCode <= 299 {
		return true
	}
	return false
}
