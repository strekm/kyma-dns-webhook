package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-acme/lego/v3/providers/dns"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type dnsReq struct {
	Domain  string `json:"domain"`
	Token   string `json:"token"`
	KeyAuth string `json:"keyauth"`
}

func getDNSReq(reqBody io.ReadCloser) (*dnsReq, error) {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading request body")
	}

	var req dnsReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		return nil, errors.Wrapf(err, "while decoding request body")
	}

	return &req, nil
}

func PresentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		provider, err := dns.NewDNSChallengeProviderByName("gcloud")

		if err != nil {
			http.Error(w, fmt.Sprintf("could not get gcloud provider: %v", err), http.StatusServiceUnavailable)
		}

		req, err := getDNSReq(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not get request body: %v", err), http.StatusBadRequest)
		}

		err = provider.Present(req.Domain, req.Token, req.KeyAuth)
		if err != nil {
			http.Error(w, fmt.Sprintf("present req failed: %v", err), http.StatusServiceUnavailable)
		}

		return
	})
}

func CleanupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		provider, err := dns.NewDNSChallengeProviderByName("gcloud")
		if err != nil {
			http.Error(w, fmt.Sprintf("could not get gcloud provider: %v", err), http.StatusServiceUnavailable)
		}

		req, err := getDNSReq(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not get request body: %v", err), http.StatusBadRequest)
		}

		err = provider.CleanUp(req.Domain, req.Token, req.KeyAuth)
		if err != nil {
			http.Error(w, fmt.Sprintf("present req failed: %v", err), http.StatusServiceUnavailable)
		}

		return
	})
}
