package metrafin

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Authorization struct {
	Application *Application
	AccessToken string
}

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

type profileRes struct {
	Error *string `json:"error"`
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

func (a *Authorization) FetchProfile() (profile *profileRes, err error) {
	auth := *a
	app := *auth.Application

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

	if parsed.Error != nil {
		return nil, errors.New(*parsed.Error)
	}

	return &parsed, nil
}
