package clash

type DNSConfig struct {
	Enable            *bool              `yaml:"enable,omitempty"`
	IPV6              *bool              `yaml:"ipv6,omitempty"`
	PreferH3          *bool              `yaml:"prefer-h3,omitempty"`
	Listen            *string            `yaml:"listen,omitempty"`
	DefaultNameserver []string           `yaml:"default-nameserver,omitempty"`
	EnhancedMode      *string            `yaml:"enhanced-mode,omitempty"`
	FakeIPRange       *string            `yaml:"fake-ip-range,omitempty"`
	UseHosts          *bool              `yaml:"use-hosts,omitempty"`
	SearchDomains     []string           `yaml:"search-domains,omitempty"`
	FakeIPFilter      []string           `yaml:"fake-ip-filter,omitempty"`
	Nameserver        []string           `yaml:"nameserver,omitempty"`
	Fallback          []string           `yaml:"fallback,omitempty"`
	FallbackFilter    FallbackFilter     `yaml:"fallback-filter,omitempty"`
	NameserverPolicy  *map[string]string `yaml:"nameserver-policy,omitempty"`
}

type FallbackFilter struct {
	GeoIP     *bool    `yaml:"geoip,omitempty"`
	GeoIPCode *string  `yaml:"geoip-code,omitempty"`
	IPCIDR    []string `yaml:"ipcidr,omitempty"`
	Domain    []string `yaml:"domain,omitempty"`
}
