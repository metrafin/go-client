package metrafin

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

// An authorization created by an Application
type Authorization struct {
	application *Application
	AccessToken string
}

type tokenInfoRes struct {
	Error string `json:"error"`
	Scopes []string `json:"scopes"`
	UserId string `json:"userId`
	Expires string `json:"expires"`
}

// A home address
type homeAddress struct {
	Full string `json:"full"`
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	City string `json:"city"`
	AdministrativeDivision string `json:"administrativeDivision"`
	AdministrativeRegion string `json:"administrativeRegion"`
	PostalCode string `json:"postalCode"`
	CountryCode string `json:"countryCode"`
}

// Profile information for a user
type profileRes struct {
	Error string `json:"error"`
	UserId string `json:"userId"`
	Username string `json:"username"`
	Created string `json:"created"`
	Verified struct {
		FirstName string `json:"firstName"`
		MiddleName string `json:"middleName"`
		LastName string `json:"lastName"`
		Country string `json:"country"`
		HomeAddress homeAddress `json:"homeAddress"`
		Age int `json:"age`
		Phone string `json:"phone"`
	}
}

// FetchInfo retrieves stats about authorization.
func (a *authorization) FetchInfo() (info *tokenInfoRes, err error) {
	auth := *a
	app := *auth.application

	res, err := doRequest(Request{
		Url: "https://api.metrafin.com/v1/token",
		Method: "GET",
		Data: &[]byte{},
		Headers: &map[string]string{
			"Authorization": app.PrivateToken + ":" + auth.AccessToken,
		},
	}, nil)

	if err != nil {
		return nil, err
	}

	all, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var parsed tokenInfoRes

	err = json.Unmarshal(all, &parsed)

	if err != nil {
		return nil, err
	}

	if parsed.Error != "" {
		return nil, errors.New(parsed.Error)
	}

	return &parsed, nil
}

// FetchProfile retrieves profile information of user.
func (a *authorization) FetchProfile() (profile *profileRes, err error) {
	auth := *a
	app := *auth.application

	res, err := doRequest(Request{
		Url: "https://api.metrafin.com/v1/user/profile",
		Method: "GET",
		Data: &[]byte{},
		Headers: &map[string]string{
			"Authorization": app.PrivateToken + ":" + auth.AccessToken,
		},
	}, nil)

	if err != nil {
		return nil, err
	}

	all, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var parsed profileRes

	err = json.Unmarshal(all, &parsed)

	if err != nil {
		return nil, err
	}

	if parsed.Error != "" {
		return nil, errors.New(parsed.Error)
	}

	return &parsed, nil
}
