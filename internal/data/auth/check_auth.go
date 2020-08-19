package auth

import (
	"context"
	"firebase/internal/entity/auth"
	"firebase/pkg/errors"
	"firebase/pkg/httpclient"
	"net/http"
)

// Data ...
type Data struct {
	client  *httpclient.Client
	baseURL string
}

// New ...
func New(client *httpclient.Client, baseURL string) Data {
	d := Data{
		client:  client,
		baseURL: baseURL,
	}
	return d
}

// CheckAuth ...
func (d Data) CheckAuth(ctx context.Context, code string) (auth.Auth, error) {
	var auth auth.Auth
	var endpoint = d.baseURL + "/checkrights"
	var body = map[string]interface{}{
		"code": code,
	}

	type ctxKey string
	_token := ctx.Value(ctxKey("_token")).(string)

	headers := make(http.Header)
	headers.Set("Authorization", _token)
	headers.Set("Content-Type", "application/json")

	_, err := d.client.PostJSON(ctx, endpoint, headers, body, &auth)
	if err != nil {
		return auth, errors.Wrap(err, "[DATA][CheckAuth]")
	}

	return auth, nil
}
