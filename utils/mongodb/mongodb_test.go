package mongodb

import (
	"github.com/tot0p/env"
	"testing"
)

// TestMongoDBPing tests the ping function of the mongodb database
func TestMongoDBPing(t *testing.T) {
	err := env.LoadPath("../../.env")
	if err != nil {
		t.Error(err)
	}
	err = NewMongoDB(env.Get("URI_MONGODB"))
	if err != nil {
		t.Error(err)
	}
	err = DB.Ping()
	if err != nil {
		t.Error(err)
	}
	err = DB.Disconnect()
	if err != nil {
		t.Error(err)
	}
}
