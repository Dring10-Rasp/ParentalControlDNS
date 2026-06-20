package main

import (
	"encoding/json"
)

const (
	OK                = 200
	CONTENT_CREATED   = 201
	NO_CONTENT        = 204
	BAD_REQUEST       = 400
	UNAUTHORIZED      = 401
	REQUEST_FAILED    = 402
	FORBIDDEN         = 403
	NOT_FOUND         = 404
	TOO_MANY_REQUESTS = 429
	SERVER_ERRORS     = 500
)

type Session struct {
	Valid    bool
	Totp     bool
	Sid      *string
	Csrf     *string
	Validity int
	Message  *string
}
type get_auth struct {
	Session Session
	Took    float64
}

type History struct {
	TimeStamp float64
	Total     int
	Cached    int
	Blocked   int
	Forwarded int
}
type history_database struct {
	History []History
	Took    float64
}

func parse_load[T any](load []byte) (*T, error) {
	var response T
	err := json.Unmarshal(load, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
