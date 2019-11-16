package metrafin

import (
	"errors"
)

// Authorization represents an authorization created by an Application
type Authorization struct {
	application *Application
	AccessToken string
}

// TokenInfoRes represents responses from a documented endpoint: https://github.com/metrafin/documentation#get-v1token
type TokenInfoRes struct {
	Error   string   `json:"error"`
	Scopes  []string `json:"scopes"`
	UserID  string   `json:"userId"`
	Expires string   `json:"expires"`
}

// A home address
type homeAddress struct {
	Full                   string `json:"full"`
	Line1                  string `json:"line1"`
	Line2                  string `json:"line2"`
	City                   string `json:"city"`
	AdministrativeDivision string `json:"administrativeDivision"`
	AdministrativeRegion   string `json:"administrativeRegion"`
	PostalCode             string `json:"postalCode"`
	CountryCode            string `json:"countryCode"`
}

// ProfileRes represents responses from a documented endpoint: https://github.com/metrafin/documentation#get-v1userprofile
type ProfileRes struct {
	Error    string `json:"error"`
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Created  string `json:"created"`
	Verified struct {
		FirstName   string      `json:"firstName"`
		MiddleName  string      `json:"middleName"`
		LastName    string      `json:"lastName"`
		Country     string      `json:"country"`
		HomeAddress homeAddress `json:"homeAddress"`
		Age         int         `json:"age"`
		Phone       string      `json:"phone"`
	}
}

// FetchInfo retrieves stats about authorization.
func (a *Authorization) FetchInfo() (info *TokenInfoRes, err error) {
	auth := *a
	app := *auth.application

	parsed := &TokenInfoRes{}

	err = doRequest(request{
		URL:    "https://api.metrafin.com/v1/token",
		Method: "GET",
		Data:   &[]byte{},
		Headers: &map[string]string{
			"Authorization": app.PrivateToken + ":" + auth.AccessToken,
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

// FetchProfile retrieves profile information of user.
func (a *Authorization) FetchProfile() (profile *ProfileRes, err error) {
	auth := *a
	app := *auth.application

	parsed := &ProfileRes{}

	err = doRequest(request{
		URL:    "https://api.metrafin.com/v1/user/profile",
		Method: "GET",
		Data:   &[]byte{},
		Headers: &map[string]string{
			"Authorization": app.PrivateToken + ":" + auth.AccessToken,
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
