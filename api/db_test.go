package api

import (
	"testing"
)

func TestGetAGI(t *testing.T) {
	User = "postgres"
	Password = "1234"

	agi, err := GetAGI("x", true, false)
	t.Log(agi.Name, "%", err)

	agi, err = GetAGI("jmcvetta/neoism", true, false)
	t.Log(agi.Name, "%", err)
}