package httplib

func Post(
	url string,
	body string,
	headers map[string]string,
	contentType string,
) (map[string]any, error) {
	finalHeaders := headers
	finalHeaders["Content-Type"] = contentType

	return Request(url, "POST", body, finalHeaders)
}
