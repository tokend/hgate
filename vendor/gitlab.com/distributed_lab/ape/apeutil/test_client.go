package apeutil

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

type Signer interface {
	SignRequest(r *http.Request) error
}

type TestClient struct {
	t      *testing.T
	ts     *httptest.Server
	signer Signer
	claims jwt.Claims
}

func (c *TestClient) SetSigner(signer Signer) *TestClient {
	c.signer = signer
	return c
}

func (c TestClient) WithClaims(claims jwt.Claims) *TestClient {
	c.claims = claims
	return &c
}

func NewClient(t *testing.T, ts *httptest.Server) *TestClient {
	return &TestClient{
		t:  t,
		ts: ts,
	}
}

func (c *TestClient) Do(method, path, body string) *http.Response {
	c.t.Helper()
	request, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.ts.URL, path), bytes.NewReader([]byte(body)))
	if err != nil {
		c.t.Fatal(err)
	}

	if c.signer != nil {
		if err := c.signer.SignRequest(request); err != nil {
			c.t.Fatal(err)
		}
	}

	if c.claims != nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.claims)
		signed, err := token.SignedString([]byte("bueno"))
		if err != nil {
			c.t.Fatal(err)
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", signed))
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		c.t.Fatal(err)
	}
	return response
}

func (c *TestClient) Get(path string) *http.Response {
	return c.Do("GET", path, ``)
}

func (c *TestClient) Put(path, body string) *http.Response {
	return c.Do("PUT", path, body)
}

func (c *TestClient) Post(path, body string) *http.Response {
	return c.Do("POST", path, body)
}

func (c *TestClient) Patch(path, body string) *http.Response {
	return c.Do("PATCH", path, body)
}

func (c *TestClient) Delete(path, body string) *http.Response {
	return c.Do("DELETE", path, body)
}
