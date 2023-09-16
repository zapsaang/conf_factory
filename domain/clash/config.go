package clash

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/zapsaang/conf_factory/utils/consts"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port               int                      `yaml:"port,omitempty"`
	SocksPort          int                      `yaml:"socks-port,omitempty"`
	RedirPort          int                      `yaml:"redir-port,omitempty"`
	TproxyPort         int                      `yaml:"tproxy-port,omitempty"`
	MixedPort          int                      `yaml:"mixed-port,omitempty"`
	RoutingMark        int                      `yaml:"routing-mark,omitempty"`
	AllowLan           bool                     `yaml:"allow-lan,omitempty"`
	IPV6               bool                     `yaml:"ipv6,omitempty"`
	BindAddress        string                   `yaml:"bind-address,omitempty"`
	Mode               string                   `yaml:"mode,omitempty"`
	LogLevel           string                   `yaml:"log-level,omitempty"`
	ExternalController string                   `yaml:"external-controller,omitempty"`
	ExternalUI         string                   `yaml:"external-ui,omitempty"`
	Secret             string                   `yaml:"secret,omitempty"`
	InterfaceName      string                   `yaml:"interface-name,omitempty"`
	Hosts              map[string]string        `yaml:"hosts,omitempty"`
	ProxyProviders     map[string]ProxyProvider `yaml:"proxy-providers,omitempty"`
	Profile            Profile                  `yaml:"profile,omitempty"`
	Experimental       Experimental             `yaml:"experimental,omitempty"`
	DNS                DNSConfig                `yaml:"dns,omitempty"`
	Proxies            []Proxy                  `yaml:"proxies,omitempty"`
	ProxyGroup         []ProxyGroup             `yaml:"proxy-groups,omitempty"`
	Inbounds           []Inbound                `yaml:"inbounds,omitempty"`
	Tunnels            []Tunnel                 `yaml:"tunnels,omitempty"`
	Rule               []string                 `yaml:"rules,omitempty"`
	Authentication     []string                 `yaml:"authentication,omitempty"`
}

type Profile struct {
	StoreSelected bool `yaml:"store-selected,omitempty"`
	StoreFakeIP   bool `yaml:"store-fake-ip,omitempty"`
}

type Experimental struct {
	UDPFallbackMatch bool `yaml:"udp-fallback-match,omitempty"`
}

type DNSConfig struct {
	Listen            string            `yaml:"listen,omitempty"`
	FakeIPRange       string            `yaml:"fake-ip-range,omitempty"`
	EnhancedMode      string            `yaml:"enhanced-mode,omitempty"`
	DefaultNameserver []string          `yaml:"default-nameserver,omitempty"`
	SearchDomains     []string          `yaml:"search-domains,omitempty"`
	FakeIPFilter      []string          `yaml:"fake-ip-filter,omitempty"`
	Nameserver        []string          `yaml:"nameserver,omitempty"`
	Fallback          []string          `yaml:"fallback,omitempty"`
	Enable            bool              `yaml:"enable,omitempty"`
	IPV6              bool              `yaml:"ipv6,omitempty"`
	UseHosts          bool              `yaml:"use-hosts,omitempty"`
	FallbackFilter    FallbackFilter    `yaml:"fallback-filter,omitempty"`
	NameserverPolicy  map[string]string `yaml:"nameserver-policy,omitempty"`
}

type FallbackFilter struct {
	GeoIP     bool     `yaml:"geoip,omitempty"`
	GeoIPCode string   `yaml:"geoip-code,omitempty"`
	IPCIDR    []string `yaml:"ipcidr,omitempty"`
	Domain    []string `yaml:"domain,omitempty"`
}

type Proxy struct {
	Port      int       `yaml:"port,omitempty"`
	AlterId   int       `yaml:"alterId,omitempty"`
	Version   int       `yaml:"version,omitempty"`
	Udp       bool      `yaml:"udp,omitempty"`
	TLS       bool      `yaml:"tls,omitempty"`
	Name      string    `yaml:"name,omitempty"`
	Type      string    `yaml:"type,omitempty"`
	Server    string    `yaml:"server,omitempty"`
	Cipher    string    `yaml:"cipher,omitempty"`
	Password  string    `yaml:"password,omitempty"`
	UUID      string    `yaml:"uuid,omitempty"`
	Username  string    `yaml:"username,omitempty"`
	SNI       string    `yaml:"sni,omitempty"`
	Obfs      string    `yaml:"obfs,omitempty"`
	Protocol  string    `yaml:"protocol,omitempty"`
	PSK       string    `yaml:"psk,omitempty"`
	Network   string    `yaml:"network,omitempty"`
	Plugin    string    `yaml:"plugin,omitempty"`
	PluginOps PluginOps `yaml:"plugin-opts,omitempty"`
	AlPN      []string  `yaml:"alpn,omitempty"`
}

type PluginOps struct {
	MaxEarlyData    int               `yaml:"max-early-data,omitempty"`
	Mode            string            `yaml:"mode,omitempty"`
	Host            string            `yaml:"host,omitempty"`
	Path            string            `yaml:"path,omitempty"`
	EarlyDataHeader string            `yaml:"early-data-header-name,omitempty"`
	ServerName      string            `yaml:"servername,omitempty"`
	TLS             bool              `yaml:"tls,omitempty"`
	SkipCertVerify  bool              `yaml:"skip-cert-verify,omitempty"`
	Mux             bool              `yaml:"mux,omitempty"`
	Headers         map[string]string `yaml:"headers,omitempty"`
	WSOptions       WSOptions         `yaml:"ws-opts,omitempty"`
	H2Options       H2Options         `yaml:"h2-opts,omitempty"`
	HttpOptions     HttpOptions       `yaml:"http-opts,omitempty"`
	GrpcOptions     GrpcOptions       `yaml:"grpc-opts,omitempty"`
	ObfsOptions     ObfsOptions       `yaml:"obfs-opts,omitempty"`
}

type WSOptions struct {
	Path    string            `yaml:"path,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

type H2Options struct {
	Host []string `yaml:"host,omitempty"`
	Path string   `yaml:"path,omitempty"`
}

type HttpOptions struct {
	Method  string              `yaml:"method,omitempty"`
	Path    []string            `yaml:"path,omitempty"`
	Headers map[string][]string `yaml:"headers,omitempty"`
}

type GrpcOptions struct {
	GrpcServiceName string `yaml:"grpc-service-name,omitempty"`
}

type ObfsOptions struct {
	Mode string `yaml:"mode,omitempty"`
	Host string `yaml:"host,omitempty"`
}

type ProxyGroup struct {
	Interval      int      `yaml:"interval,omitempty"`
	Tolerance     int      `yaml:"tolerance,omitempty"`
	RoutingMark   int      `yaml:"routing-mark,omitempty"`
	Name          string   `yaml:"name,omitempty"`
	Type          string   `yaml:"type,omitempty"`
	URL           string   `yaml:"url,omitempty"`
	Strategy      string   `yaml:"strategy,omitempty"`
	Filter        string   `yaml:"filter,omitempty"`
	InterfaceName string   `yaml:"interface-name,omitempty"`
	Proxies       []string `yaml:"proxies,omitempty"`
	Use           []string `yaml:"use,omitempty"`
	DisableUDP    bool     `yaml:"disable-udp,omitempty"`
	Lazy          bool     `yaml:"lazy,omitempty"`
}

type HealthCheck struct {
	Enable   bool   `yaml:"enable,omitempty"`
	Interval int    `yaml:"interval,omitempty"`
	Lazy     bool   `yaml:"lazy,omitempty"`
	URL      string `yaml:"url,omitempty"`
}

type ProxyProvider struct {
	Type        string      `yaml:"type,omitempty"`
	URL         string      `yaml:"url,omitempty"`
	Interval    int         `yaml:"interval,omitempty"`
	Path        string      `yaml:"path,omitempty"`
	HealthCheck HealthCheck `yaml:"health-check,omitempty"`
}

type InboundType string

const (
	InboundTypeSocks  InboundType = "socks"
	InboundTypeRedir  InboundType = "redir"
	InboundTypeTproxy InboundType = "tproxy"
	InboundTypeHTTP   InboundType = "http"
	InboundTypeMixed  InboundType = "mixed"
)

var supportInboundTypes = map[InboundType]bool{
	InboundTypeSocks:  true,
	InboundTypeRedir:  true,
	InboundTypeTproxy: true,
	InboundTypeHTTP:   true,
	InboundTypeMixed:  true,
}

type inbound struct {
	Type          InboundType `json:"type" yaml:"type,omitempty"`
	BindAddress   string      `json:"bind-address" yaml:"bind-address,omitempty"`
	IsFromPortCfg bool        `json:"-" yaml:"-,omitempty"`
}

func parseInbound(alias string) (*inbound, error) {
	u, err := url.Parse(alias)
	if err != nil {
		return nil, err
	}
	listenerType := InboundType(u.Scheme)
	return &inbound{
		Type:        listenerType,
		BindAddress: u.Host,
	}, nil
}

type Inbound inbound

func (i *Inbound) UnmarshalYAML(unmarshal func(any) error) error {
	var tp string
	if err := unmarshal(&tp); err != nil {
		var inner inbound
		if err := unmarshal(&inner); err != nil {
			return err
		}

		*i = Inbound(inner)
	} else {
		inner, err := parseInbound(tp)
		if err != nil {
			return err
		}

		*i = Inbound(*inner)
	}

	if !supportInboundTypes[i.Type] {
		return fmt.Errorf("not support inbound type: %s", i.Type)
	}
	_, portStr, err := net.SplitHostPort(i.BindAddress)
	if err != nil {
		return fmt.Errorf("bind address parse error. addr: %s, err: %w", i.BindAddress, err)
	}
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil || port == 0 {
		return fmt.Errorf("invalid bind port. addr: %s", i.BindAddress)
	}
	return nil
}

type tunnel struct {
	Network []string `yaml:"network,omitempty"`
	Address string   `yaml:"address,omitempty"`
	Target  string   `yaml:"target,omitempty"`
	Proxy   string   `yaml:"proxy,omitempty"`
}

type Tunnel tunnel

func (t *Tunnel) UnmarshalYAML(unmarshal func(any) error) error {
	var tp string
	if err := unmarshal(&tp); err != nil {
		var inner tunnel
		if err := unmarshal(&inner); err != nil {
			return err
		}

		*t = Tunnel(inner)
		return nil
	}

	// parse udp/tcp,address,target,proxy
	parts := lo.Map(strings.Split(tp, ","), func(s string, _ int) string {
		return strings.TrimSpace(s)
	})
	if len(parts) != 4 {
		return fmt.Errorf("invalid tunnel config %s", tp)
	}
	network := strings.Split(parts[0], "/")

	// validate network
	for _, n := range network {
		switch n {
		case "tcp", "udp":
		default:
			return fmt.Errorf("invalid tunnel network %s", n)
		}
	}

	address := parts[1]
	target := parts[2]
	for _, addr := range []string{address, target} {
		if _, _, err := net.SplitHostPort(addr); err != nil {
			return fmt.Errorf("invalid tunnel target or address %s", addr)
		}
	}

	*t = Tunnel(tunnel{
		Network: network,
		Address: address,
		Target:  target,
		Proxy:   parts[3],
	})
	return nil
}

func UnmarshalConfig(buf []byte) (*Config, error) {
	rawCfg := &Config{}
	buf = bytes.ReplaceAll(buf, consts.BytesTabCharacter, consts.BytesIndent4Space)
	if err := yaml.Unmarshal(buf, rawCfg); err != nil {
		return nil, err
	}

	return rawCfg, nil
}
