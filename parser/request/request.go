package request

import (
	"io"
	"net/http"
)

func GetBody(url string) (io.Reader, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		return nil, err
	}

	return resp.Body, nil
}
