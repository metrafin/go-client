// Package metrafin provides structured means for accessing the Metrafin Public API and authenticating Metrafin users.
package metrafin

import (
	"encoding/json"
	"errors"
)

// An Application for the Metrafin API
type Application struct {
	PrivateToken string
}

type createAccessTokenReq struct {
	AuthorizationCode string `json:"authorizationCode"`
}

type createAccessTokenRes struct {
	Error       string `json:"error"`
	AccessToken string `json:"accessToken"`
}

type resolveUserReq struct {
	ResolveBy string `json:"resolveBy"`
	Value     string `json:"value"`
}

// ResolveUserRes represents responses from a documented endpoint: https://github.com/metrafin/documentation#post-v1resolveuser
type ResolveUserRes struct {
	Error    string `json:"error"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
}

// Auth creates a new Authorization by "authorizationCode" or "accessToken".
func (a *Application) Auth(by string, value string) (auth *Authorization, err error) {
	if by == "authorizationCode" {
		data := createAccessTokenReq{
			AuthorizationCode: value,
		}

		serialized, err := json.Marshal(data)

		if err != nil {
			return nil, err
		}

		parsed := &createAccessTokenRes{}

		err = doRequest(request{
			URL:    "https://api.metrafin.com/v1/createAccessToken",
			Method: "POST",
			Data:   &serialized,
			Headers: &map[string]string{
				"Authorization": (*a).PrivateToken,
			},
		}, nil, parsed)

		if err != nil {
			return nil, err
		}

		if parsed.Error != "" {
			return nil, errors.New(parsed.Error)
		}

		return &Authorization{
			application: a,
			AccessToken: parsed.AccessToken,
		}, nil
	} else if by == "accessToken" {
		return &Authorization{
			application: a,
			AccessToken: value,
		}, nil
	} else {
		panic("Unknown authorization method \"" + by + "\"")
	}
}

// ResolveUser resolves a user by either "username" or "userId".
func (a *Application) ResolveUser(resolveBy string, value string) (result *ResolveUserRes, err error) {
	data := resolveUserReq{
		ResolveBy: resolveBy,
		Value:     value,
	}

	if data.ResolveBy != "username" && data.ResolveBy != "userId" {
		panic("Cannot resolve by '" + data.ResolveBy + "'")
	}

	serialized, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	parsed := &ResolveUserRes{}

	err = doRequest(request{
		URL:    "https://api.metrafin.com/v1/resolveUser",
		Method: "POST",
		Data:   &serialized,
		Headers: &map[string]string{
			"Authorization": (*a).PrivateToken,
		},
	}, nil, parsed)

	if err != nil {
		return nil, err
	}

	if parsed.Error != "" {
		return nil, errors.New(parsed.Error)
	}

	return parsed, nil
}
