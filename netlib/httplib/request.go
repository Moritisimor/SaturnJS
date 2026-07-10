package httplib

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func Request(
	serverUrl string,
	method string,
	body string,
	headers map[string]string,
) (map[string]any, error) {
	req, err := http.NewRequest(method, serverUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return map[string]any{}, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return map[string]any{}, err
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return map[string]any{}, err
	}

	resHeaders := [][]string{}
	for key, values := range res.Header {
		val := strings.Join(values, ", ")
		resHeaders = append(resHeaders, []string{key, val})
	}

	return map[string]any{
		"body":      string(resBody),
		"status":    res.StatusCode,
		"headers":   resHeaders,
		"getHeader": res.Header.Get,
		"ip":        res.Request.RemoteAddr,
	}, nil
}
