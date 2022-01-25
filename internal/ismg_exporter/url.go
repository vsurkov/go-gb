package ismgExporter

import (
	"fmt"
	urlx "github.com/goware/urlx"
	"net/url"
)

type URL struct {
	Value string
}

type URLError struct {
	url     string
	message string
}

func newURLError(url string) *URLError {
	return &URLError{url, "not valid URL value"}
}

func (err *URLError) Error() string {
	return "'" + err.url + "' " + err.message
}

func (toTest *URL) isValid() error {
	_, err := url.ParseRequestURI(toTest.Value)
	if err != nil {
		return fmt.Errorf("%v: %v", toTest.Value, err)
	}

	u, err := urlx.Parse(toTest.Value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return newURLError(toTest.Value)
	}

	return nil
}

func (URL) Create(url string) (*URL, error) {
	u := &URL{url}
	err := u.isValid()

	if err != nil {
		return nil, err
	}
	return u, nil
}
