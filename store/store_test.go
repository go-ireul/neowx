package store

import (
	"testing"
	"time"

	"ireul.com/redis"
)

func TestStoreAccessToken(t *testing.T) {
	s := NewStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	err := s.SetAccessToken("hello", "a", time.Second*10)
	if err != nil {
		t.Error(err)
	}
	k, err := s.GetAccessToken("hello")
	if err != nil {
		t.Error(err)
	}
	if k != "a" {
		t.Error("not equals " + k + "!=" + "a")
	}
}

func TestStoreJSTicket(t *testing.T) {
	s := NewStore(redis.NewClient(&redis.Options{Addr: "localhost:6379"}))
	err := s.SetJSTicket("hello", "a", time.Second*10)
	if err != nil {
		t.Error(err)
	}
	k, err := s.GetJSTicket("hello")
	if err != nil {
		t.Error(err)
	}
	if k != "a" {
		t.Error("not equals " + k + "!=" + "a")
	}
}
