package clientsdk

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client interface {
	// Close() error
	// DomainCreate(name string) error
	// DomainDelete(name string) error
	DomainsGet() ([]string, error)
	Close() error

	Joe() (string, error)
}

type client struct {
	url   string
	token string
}

type RpcLoginResp struct {
	Token string `json:"Token"`
}

type RpcDomain struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Logo string `json:"logo,omitempty"`
}

func login(baseUrl string, user string, pass string) string {
	url := baseUrl + "/adminlogin"
	authToken := fmt.Sprintf("Basic %s", B64Encode(user+":"+pass))
	data, err := GET(url, authToken, nil)
	if err != nil {
		log.Printf("adminlogin: %v\n", err)
		return ""
	}

	var resp RpcLoginResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		log.Printf("Bad JSON login response\n")
		return ""
	}

	if resp.Token == "" {
		log.Printf("Token missing in login response\n")
		return ""
	}

	return resp.Token
}

func logout(baseUrl string) {
	url := baseUrl + "logout"
	GET(url, "", nil)
}

func ClientOpen(url string, user string, pass string) (Client, error) {
	cc := client{url: url, token: ""}

	// Login
	authToken := login(cc.url, user, pass)
	if authToken == "" {
		return nil, fmt.Errorf("login failed")
	}

	cc.token = fmt.Sprintf("Bearer %s", authToken)

	return &cc, nil
}

func (c *client) DomainsGet() ([]string, error) {
	url := c.url + "/domains"
	resp, err := GET(url, c.token, nil)
	if err != nil {
		log.Printf("GetAllDomains: %v\n", err)
		return nil, err
	}

	var domains []RpcDomain
	err = json.Unmarshal(resp, &domains)
	if err != nil {
		log.Printf("GetAllDomains: Unmarshal %v\n", err)
		return nil, err
	}

	var names []string
	for _, domain := range domains {
		names = append(names, domain.Name)
	}

	return names, nil
}

func (c *client) DomainCreate() error {
	url := c.url + "/domains"
	resp, err := POST(url, c.token, nil)
	if err != nil {
		log.Printf("GetAllDomains: %v\n", err)
		return nil, err
	}

	var domains []RpcDomain
	err = json.Unmarshal(resp, &domains)
	if err != nil {
		log.Printf("GetAllDomains: Unmarshal %v\n", err)
		return nil, err
	}

	var names []string
	for _, domain := range domains {
		names = append(names, domain.Name)
	}

	return names, nil
}

func (c *client) Joe() (string, error) {
	log.Printf(c.url)

	return "d", nil
}

func (c *client) Close() error {
	logout(c.url)

	return nil
}
