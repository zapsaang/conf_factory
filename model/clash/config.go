package clash

import (
	"bytes"

	"gopkg.in/yaml.v3"

	"github.com/zapsaang/conf_factory/utils/consts"
)

type RawConfig struct {
	Port                 *int              `yaml:"port,omitempty"`
	SocksPort            *int              `yaml:"socks-port,omitempty"`
	RedirPort            *int              `yaml:"redir-port,omitempty"`
	TproxyPort           *int              `yaml:"tproxy-port,omitempty"`
	MixedPort            *int              `yaml:"mixed-port,omitempty"`
	Authentication       []string          `yaml:"authentication,omitempty"`
	AllowLan             *bool             `yaml:"allow-lan,omitempty"`
	BindAddress          *string           `yaml:"bind-address,omitempty"`
	Mode                 *string           `yaml:"mode,omitempty"`
	LogLevel             *string           `yaml:"log-level,omitempty"`
	IPV6                 *bool             `yaml:"ipv6,omitempty"`
	ExternalController   *string           `yaml:"external-controller,omitempty"`
	ExternalUI           *string           `yaml:"external-ui,omitempty"`
	Secret               *string           `yaml:"secret,omitempty"`
	InterfaceName        *string           `yaml:"interface-name,omitempty"`
	RoutingMark          *int              `yaml:"routing-mark,omitempty"`
	Hosts                Hosts             `yaml:"hosts,omitempty"`
	Profile              Profile           `yaml:"profile,omitempty"`
	DNS                  DNSConfig         `yaml:"dns,omitempty"`
	Experimental         Experimental      `yaml:"experimental,omitempty"`
	CFWLatencyTimeout    *int              `yaml:"cfw-latency-timeout,omitempty"`
	CFWLatencyURL        *string           `yaml:"cfw-latency-url,omitempty"`
	CFWConnBreakStrategy ConnBreakStrategy `yaml:"cfw-conn-break-strategy,omitempty"`
	CFWBypass            []string          `yaml:"cfw-bypass,omitempty"`
	Proxies              Proxies           `yaml:"proxies,omitempty"`
	ProxyGroup           ProxyGroups       `yaml:"proxy-groups,omitempty"`
	ProxyProviders       ProxyProviders    `yaml:"proxy-providers,omitempty"`
	Script               Script            `yaml:"script,omitempty"`
	RuleProviders        RuleProviders     `yaml:"rule-providers,omitempty"`
	Inbounds             []Inbound         `yaml:"inbounds,omitempty"`
	Tunnels              []Tunnel          `yaml:"tunnels,omitempty"`
	Rule                 []string          `yaml:"rules,omitempty"`
}

type Profile struct {
	StoreSelected bool `yaml:"store-selected"`
	StoreFakeIP   bool `yaml:"store-fake-ip"`
	Tracing       bool `yaml:"tracing"`
}

type Experimental struct {
	UDPFallbackMatch  bool `yaml:"udp-fallback-match"`
	SniffTLSSNI       bool `yaml:"sniff-tls-sni"`
	IgnoreResolveFail bool `yaml:"ignore-resolve-fail"`
}

type ConnBreakStrategy struct {
	Proxy   string `yaml:"proxy"`
	Profile bool   `yaml:"profile"`
	Mode    bool   `yaml:"mode"`
}

func UnmarshalRawConfig(buf []byte) (*RawConfig, error) {
	rawCfg := &RawConfig{}
	buf = bytes.ReplaceAll(buf, consts.BytesTabCharacter, consts.BytesIndent4Space)
	if err := yaml.Unmarshal(buf, rawCfg); err != nil {
		return nil, err
	}

	return rawCfg, nil
}

func MarshalRawConfig(cfg *RawConfig) ([]byte, error) {
	return yaml.Marshal(cfg)
}
