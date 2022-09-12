package proxy

import (
	"net/http"
)

func NewUrlRedirector(inputUrl, redirectUrl string) *UrlRedirector {
	return &UrlRedirector{
		inputUrl:    inputUrl,
		redirectUrl: redirectUrl,
	}
}

// UrlRedirector is a RequestModifier that redirect all requests from
// inputUrl/* -> redirectUrl/*
type UrlRedirector struct {
	inputUrl    string
	redirectUrl string
}

// UrlRedirector rewrites any requests that match the pattern inputUrl/*
// to redirectUrl/*
func (r UrlRedirector) ModifyRequest(req *http.Request) error {
	if req.URL.Hostname() == r.inputUrl {
		req.Host = r.redirectUrl
	}
	return nil
}
