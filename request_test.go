package metrafin

import (
	"testing"
)

type errorRes struct {
	Error string `json:"error"`
}

func TestRequest(t *testing.T) {
	out := errorRes{}

	err := doRequest(request{
		URL:    "https://api.metrafin.com",
		Method: "GET",
		Data:   &[]byte{},
		Headers: &map[string]string{
			"Host": "api.metrafin.com",
		},
	}, nil, &out)

	if err != nil {
		t.Error(err)
		return
	}
}
