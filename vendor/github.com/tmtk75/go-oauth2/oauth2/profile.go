package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"

	xoauth2 "golang.org/x/oauth2"
)

// Profile gives some methods to retrieve major values from profile data.
type Profile interface {
	Token() *xoauth2.Token
	Name() string
	Nickname() string
	AvatarURL() string
}

// GetProfileData retrieves profile data using an endpoint and a token string
func GetProfileData(endpoint, token string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	ctype, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch ctype {
	case "application/json", "text/javascript":
		var data map[string]interface{}
		json.Unmarshal(b, &data)
		return data, nil
	}
	return map[string]interface{}{}, fmt.Errorf("unknown Content-Type: %v", ctype)
}
