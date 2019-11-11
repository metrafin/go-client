package metrafin

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

type Request struct {
	Url string
	Method string
	Data *[]byte
	Headers *map[string]string
}

func doRequest(request Request, client *http.Client, output interface{}) (error error) {
	innerReq, err := http.NewRequest(request.Method, request.Url, bytes.NewReader(*request.Data))

	if err != nil {
		return err
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

	resp, err := client.Do(innerReq)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(all, &output)

	if err != nil {
		return err
	}

	return nil
}
