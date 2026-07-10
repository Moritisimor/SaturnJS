package httplib

func Get(url string) (map[string]any, error) {
	return Request(url, "GET", "", map[string]string{})
}
