package basicauth

// thanks @cpliakas
// https://github.com/google/go-github/issues/119#issuecomment-62276682

import (
	"net/http"
)

type Transport struct {
	Username string
	Password string
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

func (t *Transport) Client() *http.Client {
	return &http.Client{Transport: t}
}
