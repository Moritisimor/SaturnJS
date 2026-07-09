package netlib

import (
	"io"
	"net/http"
	"strings"
)

func httpGet(url string) (map[string]any, error) {
	response := map[string]any{}
	res, err := http.Get(url)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	headers := [][]string{}
	for key, values := range res.Header {
		val := strings.Join(values, ", ")
		headers = append(headers, []string{key, val})
	}

	response["body"] = string(body)
	response["status"] = res.StatusCode
	response["headers"] = headers
	response["getHeader"] = res.Header.Get
	response["ip"] = res.Request.RemoteAddr

	return response, nil
}
