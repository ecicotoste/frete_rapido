package ws

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func WSRequest(metodo string, apiUrl string, payloadBytes []byte) ([]byte, error) {

	req, err := http.NewRequest(metodo, apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		strError := fmt.Sprintf("RequestApiMcc StatusCode: %d --> %s", res.StatusCode, string(body))
		return body, errors.New(strError)
	}

	return body, nil

}
