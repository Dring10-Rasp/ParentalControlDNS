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

type SuccessItem struct {
	Item string
}

type ErrorItem struct {
	Item  string
	Error string
}

type Processed struct {
	Success []SuccessItem
	Errors  []ErrorItem
}

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

type Groups struct {
	Name          string
	Comment       *string
	Enabled       bool
	Id            int
	Date_added    int64
	Date_modified int64
}

type Clients struct {
	Client        string
	Comment       *string
	Groups        int
	Id            int64
	Date_added    int64
	Date_modified int64
	name          *string
}

type UknownClients struct {
	Hwaddr    *string
	MacVendor *string
	LastQuery int
	Addresses *string
	Names     *string
}

type Lists struct {
	Address        string
	Type           string
	Comment        *string
	Groups         []int
	Enabled        bool
	ID             int
	DateAdded      int64
	DateModified   int64
	DateUpdated    int64
	Number         int
	InvalidDomains int
	AbpEntries     int
	Status         int
}

type ListResponse struct {
	Lists     []Lists
	Processed *Processed
	Took      float64
}

type post_groups struct {
	Groups    []Groups
	Processed *Processed
	Took      float64
}

type get_groups struct {
	Groups []Groups
	Took   float64
}

type post_clients struct {
	Clients   []Clients
	Processed *Processed
	Took      float64
}

type get_clients struct {
	Clients []Clients
	Took    float64
}

type get_clients_suggestrions struct {
	Clients []UknownClients
	Took    float64
}

type get_lists struct {
	Lists []Lists
	Took  float64
}

type Domains struct {
	Domain        string
	Unicode       string
	Type          string
	Kind          string
	Comment       *string
	Groups        []int
	Enabled       bool
	Id            int64
	Date_added    int64
	Date_modified int64
}

type get_domains struct {
	Domains []Domains
	Took    float64
}

type post_domains struct {
	Domains   []Domains
	Processed *Processed
	Took      float64
}

func parse_load[T any](load []byte) (*T, error) {
	var response T
	err := json.Unmarshal(load, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
