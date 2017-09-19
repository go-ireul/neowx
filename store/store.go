package store

import (
	"time"

	"ireul.com/redis"
)

// Store represents a redis store for Wechat MP account
type Store struct {
	*redis.Client
}

// AccessTokenKeyPrefix
const (
	AccessTokenPrefix = "access_token."
	JSTicketPrefix    = "js_ticket."
	OAuthURLPrefix    = "oauth_url."
)

// NewStore create a store
func NewStore(c *redis.Client) *Store {
	return &Store{Client: c}
}

// SetAccessToken set access token to appID
func (s Store) SetAccessToken(appID string, token string, exp time.Duration) error {
	return s.Set(AccessTokenPrefix+appID, token, exp).Err()
}

// SetJSTicket set access token to appID
func (s Store) SetJSTicket(appID string, ticket string, exp time.Duration) error {
	return s.Set(JSTicketPrefix+appID, ticket, exp).Err()
}

// GetAccessToken get access token for appId
func (s Store) GetAccessToken(appID string) (string, error) {
	return s.Get(AccessTokenPrefix + appID).Result()
}

// GetJSTicket get access token for appId
func (s Store) GetJSTicket(appID string) (string, error) {
	return s.Get(JSTicketPrefix + appID).Result()
}
