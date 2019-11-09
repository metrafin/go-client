package metrafin

import (
	"errors"
	"io/ioutil"
	"encoding/json"
)

type Application struct {
	PrivateToken string
}

type createAccessTokenReq struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type createAccessTokenRes struct {
	Error *string `json:"error"`
	AccessToken string `json:"accessToken"`
}

type resolveUserReq struct {
	ResolveBy string `json:"resolveBy"`
	Value string `json:"value"`
}

type resolveUserRes struct {
	error *string `json:"error"`
	UserId string `json:"userId"`
	Username string `json:"username"`
}

func (a *Application) Auth (by string, value string) (auth *Authorization, err error) {
	if by == "authorizationCode" {
		data := createAccessTokenReq{
			AuthorizationCode: value,
		}

		serialized, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		resp, err := doRequest(Request{
			Url: "https://api.metrafin.com/v1/createAccessToken",
			Method: "POST",
			Data: &serialized,
			Headers: &map[string]string{
				"Authorization": (*a).PrivateToken,
			},
		}, nil)

		if err != nil {
			return nil, err
		}

		all, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		var parsed createAccessTokenRes

		err = json.Unmarshal(all, &parsed)

		if err != nil {
			return nil, err
		}

		if parsed.Error != nil {
			return nil, errors.New(*parsed.Error)
		}

		return &Authorization{
			Application: a,
			AccessToken: parsed.AccessToken,
		}, nil
	} else if by == "accessToken" {
		return &Authorization{
			Application: a,
			AccessToken: value,
		}, nil
	} else {
		return nil, errors.New("Unknown authorization method \"" + by + "\"")
	}
}

func (a *Application) ResolveUser (resolveBy string, value string) (result *resolveUserRes, err error) {
	data := resolveUserReq{
		ResolveBy: resolveBy,
		Value: value,
	}

	if data.ResolveBy != "username" && data.ResolveBy != "userId" {
		return nil, errors.New("Cannot resolve by '" + data.ResolveBy + "'")
	}

	serialized, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	resp, err := doRequest(Request{
		Url: "https://api.metrafin.com/v1/resolveUser",
		Method: "POST",
		Data: &serialized,
		Headers: &map[string]string{
			"Authorization": (*a).PrivateToken,
		},
	}, nil)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	all, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var res resolveUserRes

	err = json.Unmarshal(all, &res)

	if err != nil {
		return nil, err
	}

	if res.error != nil {
		return nil, errors.New(*res.error)
	} else {
		return &res, nil
	}
}
