package main

import (
	"log"
	"testing"

	"github.com/logos106/clientsdk-go/clientsdk"
)

func TestJoe(t *testing.T) {
	// Call ClientOpen
	baseUrl := "http://localhost:8000/api/v1"
	user := "admin@powerdomain"
	pass := "password"
	cc, err := clientsdk.ClientOpen(baseUrl, user, pass)
	if err != nil {
		t.Fatalf(`Error while running ClientOpen. %v`, err)
	}

	// Call DomainsGet
	names, err := cc.DomainsGet()
	if err != nil {
		t.Fatalf(`Error while running DomainsGet. %v`, err)
	}

	// Call DomainsGet
	err = cc.DomainCreate()
	if err != nil {
		t.Fatalf(`Error while running DomainCreate. %v`, err)
	}

	// Call Close
	err = cc.Close()
	if err != nil {
		t.Fatalf(`Error while running Close. %v`, err)
	}

	log.Println(names)
}
