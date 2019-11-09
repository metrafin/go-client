package metrafin

import (
	"net/http"
	"bytes"
)

type Request struct {
	Url string
	Method string
	Data *[]byte
	Headers *map[string]string
}

func doRequest(request Request, client *http.Client) (resp *http.Response, error error) {
	innerReq, err := http.NewRequest(request.Method, request.Url, bytes.NewReader(*request.Data))

	if err != nil {
		return nil, err
	}

	if request.Headers != nil {
		headers := *request.Headers

		for key, value := range headers {
			innerReq.Header.Set(key, value)
		}
	}

	if client == nil {
		client = &http.Client{}
	}

	resp, err = client.Do(innerReq)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
