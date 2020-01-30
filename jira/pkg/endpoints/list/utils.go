package list

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func extendRequestWithAuthenticatedUser(user string, token string) func(req *http.Request) (*http.Request, error) {
	return func(req *http.Request) (*http.Request, error) {
		authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, token))))
		req.Header.Add("Authorization", authHeader)
		return req, nil
	}
}
