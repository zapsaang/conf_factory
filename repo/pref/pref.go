package pref

type RenameRule struct {
	Action  string `yaml:"action"`
	Express string `yaml:"express"`
	Content string `yaml:"content"`
}

type Subscription struct {
	UpdateInterval int64  `yaml:"update_interval"`
	Source         string `yaml:"source"`
	Fitler         string `yaml:"filter"`
	Sort           string `yaml:"sort"`
	// Proxy          string       `yaml:"proxy"`
	Nodes       []string     `yaml:"nodes"`
	Rename      []RenameRule `yaml:"rename"`
	InsertEmoji bool         `yaml:"insert_emoji"`
}

type Template struct {
	Platform string `yaml:"platform"`
	Source   string `yaml:"source"`
	// Proxy    string `yaml:"proxy"`
	KV []string `yaml:"kv"`
}

type RuleSet struct {
	UpdateInterval int64  `yaml:"update_interval"`
	Source         string `yaml:"source"`
	Action         string `yaml:"action"`
	Type           string `yaml:"type"`
}

type ProxyGroup struct {
	Interval  int64  `yaml:"interval"`
	Tolerance int64  `yaml:"tolerance"`
	Timeout   int64  `yaml:"timeout"`
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	Fitler    string `yaml:"filter"`
	TestURL   string `yaml:"test_url"`
}

type Preference struct {
	Version       int64
	Subscriptions []Subscription
	Templates     []Template
	RuleSets      []RuleSet
	ProxyGroups   []ProxyGroup
}
