package ydb

import (
	"context"
	"errors"
)

var (
	// ErrCredentialsNoCredentials may be returned by Credentials implementations to
	// make driver act as if there no Credentials at all. That is, driver will
	// not send any token meta information during request.
	ErrCredentialsNoCredentials = errors.New("ydb: credentials: no credentials")
)

// Credentials is an interface that contains options used to authorize a
// client.
type Credentials interface {
	Token(context.Context) (string, error)
}

// CredentialsFunc is an adapter to allow the use of ordinary functions as
// Credentials.
type CredentialsFunc func(context.Context) (string, error)

// Token implements Credentials.
func (f CredentialsFunc) Token(ctx context.Context) (string, error) {
	return f(ctx)
}

// AuthTokenCredentials implements Credentials interface with static
// authorization parameters.
type AuthTokenCredentials struct {
	AuthToken string
}

// Token implements Credentials.
func (a AuthTokenCredentials) Token(_ context.Context) (string, error) {
	return a.AuthToken, nil
}

type multiCredentials struct {
	cs []Credentials
}

func (m *multiCredentials) Token(ctx context.Context) (token string, err error) {
	for _, c := range m.cs {
		token, err = c.Token(ctx)
		if err == nil {
			return
		}
	}
	if err == nil {
		err = ErrCredentialsNoCredentials
	}
	return
}

// MultiCredentials creates Credentials which represents multiple ways of
// obtaining token.
// Its Token() method proxies call to the underlying credentials in order.
// When first successful call met, it returns. If there are no successful
// calls, it returns last error.
func MultiCredentials(cs ...Credentials) Credentials {
	all := make([]Credentials, 0, len(cs))
	for _, c := range cs {
		if m, ok := c.(*multiCredentials); ok {
			all = append(all, m.cs...)
		} else {
			all = append(all, c)
		}
	}
	return &multiCredentials{
		cs: all,
	}
}
