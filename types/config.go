package types

import "ireul.com/com"

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
}

// Rule represents a rule
type Rule struct {
	Match     com.Map `yaml:"match"`
	Text      string  `yaml:"text"`
	HTTPSync  string  `yaml:"http_sync"`
	HTTPAsync string  `yaml:"http_async"`
}
