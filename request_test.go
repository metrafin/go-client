package metrafin

import (
	"testing"
)

func TestRequest (t *testing.T) {
	resp, err := doRequest(Request{
		Url: "https://api.metrafin.com",
		Method: "GET",
		Data: &[]byte{},
		Headers: &map[string]string{
			"Host": "api.metrafin.com",
		},
	}, nil)

	if err != nil {
		t.Error(err)
		return
	}

	if resp.StatusCode != 0 {
		return
	}
}
