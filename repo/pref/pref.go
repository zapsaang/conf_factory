package pref

import "github.com/BurntSushi/toml"

type Server struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`

	Token  string `toml:"token"`
	Domain string `toml:"domain"`

	LogLevel string `toml:"log_level"`
	LogDir   string `toml:"log_dir"`

	WorkDir   string `toml:"work_dir"`
	HTTPProxy string `toml:"http_proxy"`
}

type RenameRule struct {
	Action  string `toml:"action"`
	Express string `toml:"express"`
	Content string `toml:"content"`
}

type Subscription struct {
	UpdateInterval int64        `toml:"update_interval"`
	Source         string       `toml:"source"`
	Fitler         string       `toml:"filter"`
	Sort           string       `toml:"sort"`
	Nodes          []string     `toml:"nodes"`
	Rename         []RenameRule `toml:"rename"`
	Emoji          []RenameRule `toml:"emoji"`
	RenameFile     string       `toml:"rename_file"`
	EmojiFile      string       `toml:"emoji_file"`
}

type Template struct {
	Platform string `toml:"platform"`
	Source   string `toml:"source"`
}

type RuleSet struct {
	UpdateInterval int64  `toml:"update_interval"`
	Source         string `toml:"source"`
	Action         string `toml:"action"`
	Type           string `toml:"type"`
}

type ProxyGroup struct {
	Interval  int64  `toml:"interval"`
	Tolerance int64  `toml:"tolerance"`
	Timeout   int64  `toml:"timeout"`
	Name      string `toml:"name"`
	Type      string `toml:"type"`
	Fitler    string `toml:"filter"`
	TestURL   string `toml:"test_url"`
}

type Preference struct {
	Version      uint8  `toml:"version"`
	ServerConfig Server `toml:"server"`

	Templates []Template `toml:"templates"`

	Rename     []RenameRule `toml:"rename"`
	Emoji      []RenameRule `toml:"emoji"`
	RenameFile string       `toml:"rename_file"`
	EmojiFile  string       `toml:"emoji_file"`

	Subscriptions []Subscription `toml:"subcriptions"`
	RuleSets      []RuleSet      `toml:"rule_sets"`
	ProxyGroups   []ProxyGroup   `toml:"proxy_groups"`
}

func UnmarshalPref(buf []byte) (*Preference, error) {
	p := &Preference{}
	if err := toml.Unmarshal(buf, p); err != nil {
		return nil, err
	}
	return p, nil
}
