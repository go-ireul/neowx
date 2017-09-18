package types

// Config represents a config.yaml file
type Config struct {
	Env      string    `yaml:"env"`
	Host     string    `yaml:"host"`
	Port     int       `yaml:"port"`
	RedisURL string    `yaml:"redis_url"`
	Accounts []Account `yaml:"accounts"`
}

// Account represents a MP account
type Account struct {
	Name      string `yaml:"name"`
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	OrignalID string `yaml:"orignal_id"`
	Default   bool   `yaml:"default"`
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
