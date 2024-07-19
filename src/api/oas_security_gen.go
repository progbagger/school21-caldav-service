// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleAuthorization handles Authorization security.
	HandleAuthorization(ctx context.Context, operationName string, t Authorization) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityAuthorization(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t Authorization
	const parameterName = "Authorization"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleAuthorization(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// Authorization provides Authorization security value.
	Authorization(ctx context.Context, operationName string) (Authorization, error)
}

func (s *Client) securityAuthorization(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.Authorization(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"Authorization\"")
	}
	req.Header.Set("Authorization", t.APIKey)
	return nil
}
