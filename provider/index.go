package provider

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	. "gaming-service/model"
)

func Request(method string, url string, body io.Reader, headers map[string]string) (HttpResponse, error) {
	client := &http.Client{}
	response := HttpResponse{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		return response, err
	}
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
	res, err := client.Do(request)

	if err != nil || res.StatusCode != http.StatusOK {
		response.StatusCode = res.StatusCode
		if err == nil {
			err = errors.New("API 請求失敗")
		}
		return response, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		return response, err
	}
	response.StatusCode = http.StatusOK
	response.Response = resBody
	return response, nil
}
