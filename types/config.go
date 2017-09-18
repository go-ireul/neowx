package types

import (
	"regexp"
	"strings"

	"ireul.com/com"
	"ireul.com/structs"
)

// Config represents a config.yaml file
type Config struct {
	Env      string    `yaml:"env"`
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	RedisURL string    `yaml:"redis_url"`
	Accounts []Account `yaml:"accounts"`
	Rules    []Rule    `yaml:"rules"`
}

// Account represents a MP account
type Account struct {
	Name      string `yaml:"name"`
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	OrignalID string `yaml:"orignal_id"`
	Default   bool   `yaml:"default"`
}

// Rule represents a rule
type Rule struct {
	Match     map[string]string `yaml:"match"`
	Text      string            `yaml:"text"`
	HTTPSync  string            `yaml:"http_sync"`
	HTTPAsync string            `yaml:"http_async"`
}

// Matches match a rule against a WxReq
func (r Rule) Matches(req WxReq) (bool, error) {
	if len(r.Match) > 0 {
		m := structs.Map(req)
		for k, v := range r.Match {
			v0 := com.ToStr(m[k])
			if len(v) > 2 && v[0] == '/' && v[len(v)-1] == '/' {
				// regexp
				ok, err := regexp.MatchString(v[1:len(v)-1], v0)
				if err != nil || !ok {
					return false, err
				}
			} else {
				// simple match
				if strings.ToLower(strings.TrimSpace(v)) != strings.ToLower(strings.TrimSpace(v0)) {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

// AccountByName return a Account with name
func (c Config) AccountByName(n string) Account {
	for _, a := range c.Accounts {
		if a.Name == n {
			return a
		}
	}
	return Account{}
}

// DefaultAccount return a default Account
func (c Config) DefaultAccount() Account {
	for _, a := range c.Accounts {
		if a.Default {
			return a
		}
	}
	if len(c.Accounts) > 0 {
		return c.Accounts[0]
	}
	return Account{}
}
