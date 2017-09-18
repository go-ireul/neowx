package types

// RuleMatcher a matcher for rule
type RuleMatcher func(m WxReq, cfg Config) bool

// Rule represents a rule
type Rule struct {
	Fn        RuleMatcher
	Text      string
	HTTPSync  string
	HTTPAsync string
}
