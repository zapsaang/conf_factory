package clash

import (
	"fmt"

	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/utils/consts"
	"gopkg.in/yaml.v3"
)

type Proxy struct {
	Name   string `yaml:"name"`
	Type   string `yam:"type"`
	Server string `yaml:"server"`

	// basic
	TFO         *bool   `yaml:"tfo,omitempty"`
	MPTCP       *bool   `yaml:"mptcp,omitempty"`
	Interface   *string `yaml:"interface-name,omitempty"`
	RoutingMark *int    `yaml:"routing-mark,omitempty"`
	IPVersion   *string `yaml:"ip-version,omitempty"`
	DialerProxy *string `yaml:"dialer-proxy,omitempty"`

	// ss

	// Port   int    `yaml:"port"`
	// Password          string         `yaml:"password"`
	Cipher            *string       `yaml:"cipher,omitempty"`
	UDP               *bool         `yaml:"udp,omitempty"`
	Plugin            *string       `yaml:"plugin,omitempty"`
	PluginOpts        PluginOptions `yaml:"plugin-opts,omitempty"`
	UDPOverTCP        *bool         `yaml:"udp-over-tcp,omitempty"`
	UDPOverTCPVersion *int          `yaml:"udp-over-tcp-version,omitempty"`
	ClientFingerprint *string       `yaml:"client-fingerprint,omitempty"`

	// ssr
	// Name          string `yaml:"name"`
	// Server        string `yaml:"server"`
	// Port          int    `yaml:"port"`
	// Password      string `yaml:"password"`
	// Cipher        string `yaml:"cipher"`
	// Obfs          string `yaml:"obfs"`
	ObfsParam *string `yaml:"obfs-param,omitempty"`
	// Protocol      string `yaml:"protocol"`
	ProtocolParam *string `yaml:"protocol-param,omitempty"`
	// UDP           bool   `yaml:"udp,omitempty"`

	// sock5
	// Name           string `yaml:"name"`
	// Server         string `yaml:"server"`
	// Port           int    `yaml:"port"`
	UserName *string `yaml:"username,omitempty"`
	Password *string `yaml:"password,omitempty"`
	TLS      *bool   `yaml:"tls,omitempty"`
	// UDP            bool   `yaml:"udp,omitempty"`
	SkipCertVerify *bool   `yaml:"skip-cert-verify,omitempty"`
	Fingerprint    *string `yaml:"fingerprint,omitempty"`

	// http
	// Name           string            `yaml:"name"`
	// Server         string            `yaml:"server"`
	// Port           int               `yaml:"port"`
	// UserName       string            `yaml:"username,omitempty"`
	// Password       string            `yaml:"password,omitempty"`
	// TLS            bool              `yaml:"tls,omitempty"`
	SNI *string `yaml:"sni,omitempty"`
	// SkipCertVerify bool              `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint    string            `yaml:"fingerprint,omitempty"`
	Headers *map[string]string `yaml:"headers,omitempty"`

	// vmess
	// Name                string         `yaml:"name"`
	// Server              string         `yaml:"server"`
	// Port                int            `yaml:"port"`
	UUID    *string `yaml:"uuid,omitempty"`
	AlterID *int    `yaml:"alterId,omitempty"`
	// Cipher              string         `yaml:"cipher"`
	// UDP                 bool           `yaml:"udp,omitempty"`
	Network *string `yaml:"network,omitempty"`
	// TLS                 bool           `yaml:"tls,omitempty"`
	ALPN []string `yaml:"alpn,omitempty"`
	// SkipCertVerify      bool           `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint         string         `yaml:"fingerprint,omitempty"`
	ServerName          *string        `yaml:"servername,omitempty"`
	RealityOpts         RealityOptions `yaml:"reality-opts,omitempty"`
	HTTPOpts            HTTPOptions    `yaml:"http-opts,omitempty"`
	HTTP2Opts           HTTP2Options   `yaml:"h2-opts,omitempty"`
	GrpcOpts            GRPCOptions    `yaml:"grpc-opts,omitempty"`
	WSOpts              WSOptions      `yaml:"ws-opts,omitempty"`
	PacketAddr          *bool          `yaml:"packet-addr,omitempty"`
	XUDP                *bool          `yaml:"xudp,omitempty"`
	PacketEncoding      *string        `yaml:"packet-encoding,omitempty"`
	GlobalPadding       *bool          `yaml:"global-padding,omitempty"`
	AuthenticatedLength *bool          `yaml:"authenticated-length,omitempty"`
	// ClientFingerprint   string         `yaml:"client-fingerprint,omitempty"`

	// vless
	// Name              string            `yaml:"name"`
	// Server            string            `yaml:"server"`
	// Port              int               `yaml:"port"`
	// UUID              string            `yaml:"uuid"`
	Flow *string `yaml:"flow,omitempty"`
	// TLS               bool              `yaml:"tls,omitempty"`
	// ALPN              []string          `yaml:"alpn,omitempty"`
	// UDP               bool              `yaml:"udp,omitempty"`
	// PacketAddr        bool              `yaml:"packet-addr,omitempty"`
	// XUDP              bool              `yaml:"xudp,omitempty"`
	// PacketEncoding    string            `yaml:"packet-encoding,omitempty"`
	// Network           string            `yaml:"network,omitempty"`
	// RealityOpts       RealityOptions    `yaml:"reality-opts,omitempty"`
	// HTTPOpts          HTTPOptions       `yaml:"http-opts,omitempty"`
	// HTTP2Opts         HTTP2Options      `yaml:"h2-opts,omitempty"`
	// GrpcOpts          GrpcOptions       `yaml:"grpc-opts,omitempty"`
	// WSOpts            WSOptions         `yaml:"ws-opts,omitempty"`
	WSPath    *string            `yaml:"ws-path,omitempty"`
	WSHeaders *map[string]string `yaml:"ws-headers,omitempty"`
	// SkipCertVerify    bool              `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint       string            `yaml:"fingerprint,omitempty"`
	// ServerName        string            `yaml:"servername,omitempty"`
	// ClientFingerprint string            `yaml:"client-fingerprint,omitempty"`

	//snell
	// Name     string         `yaml:"name"`
	// Server   string         `yaml:"server"`
	// Port     int            `yaml:"port"`
	Psk *string `yaml:"psk,omitempty"`
	// UDP      bool           `yaml:"udp,omitempty"`
	Version  *int        `yaml:"version,omitempty"`
	ObfsOpts ObfsOptions `yaml:"obfs-opts,omitempty"`

	// trojan
	// Name              string         `yaml:"name"`
	// Server            string         `yaml:"server"`
	// Port              int            `yaml:"port"`
	// Password          string         `yaml:"password"`
	// ALPN              []string       `yaml:"alpn,omitempty"`
	// SNI               string         `yaml:"sni,omitempty"`
	// SkipCertVerify    bool           `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint       string         `yaml:"fingerprint,omitempty"`
	// UDP               bool           `yaml:"udp,omitempty"`
	// Network           string         `yaml:"network,omitempty"`
	// RealityOpts       RealityOptions `yaml:"reality-opts,omitempty"`
	// GrpcOpts          GrpcOptions    `yaml:"grpc-opts,omitempty"`
	// WSOpts            WSOptions      `yaml:"ws-opts,omitempty"`
	// ClientFingerprint string         `yaml:"client-fingerprint,omitempty"`

	// hysteria
	// Name                string   `yaml:"name"`
	// Server              string   `yaml:"server"`
	Port         *int    `yaml:"port,omitempty"`
	Ports        *string `yaml:"ports,omitempty"`
	Protocol     *string `yaml:"protocol,omitempty"`
	ObfsProtocol string  `yaml:"obfs-protocol,omitempty"` // compatible with Stash
	Up           *string `yaml:"up,omitempty"`
	UpSpeed      int     `yaml:"up-speed,omitempty"` // compatible with Stash
	Down         *string `yaml:"down,omitempty"`
	DownSpeed    int     `yaml:"down-speed,omitempty"` // compatible with Stash
	Auth         *string `yaml:"auth,omitempty"`
	AuthString   *string `yaml:"auth-str,omitempty"`
	Obfs         *string `yaml:"obfs,omitempty"`
	// SNI                 string   `yaml:"sni,omitempty"`
	// SkipCertVerify      bool     `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint         string   `yaml:"fingerprint,omitempty"`
	// ALPN                []string `yaml:"alpn,omitempty"`
	CustomCA            *string `yaml:"ca,omitempty"`
	CustomCAString      *string `yaml:"ca-str,omitempty"`
	ReceiveWindowConn   *int    `yaml:"recv-window-conn,omitempty"`
	ReceiveWindow       *int    `yaml:"recv-window,omitempty"`
	DisableMTUDiscovery *bool   `yaml:"disable-mtu-discovery,omitempty"`
	FastOpen            *bool   `yaml:"fast-open,omitempty"`
	HopInterval         *int    `yaml:"hop-interval,omitempty"`

	// wireguard
	// Server       string   `yaml:"server"`
	// Port         int      `yaml:"port"`
	Ip           *string  `yaml:"ip,omitempty"`
	Ipv6         *string  `yaml:"ipv6,omitempty"`
	PublicKey    *string  `yaml:"public-key,omitempty"`
	PreSharedKey *string  `yaml:"pre-shared-key,omitempty"`
	Reserved     []uint8  `yaml:"reserved,omitempty"`
	AllowedIPs   []string `yaml:"allowed-ips,omitempty"`
	// Name                string `yaml:"name"`
	PrivateKey *string `yaml:"private-key,omitempty"`
	Workers    *int    `yaml:"workers,omitempty"`
	MTU        *int    `yaml:"mtu,omitempty"`
	// UDP                 bool   `yaml:"udp,omitempty"`
	PersistentKeepalive *int                  `yaml:"persistent-keepalive,omitempty"`
	Peers               []WireGuardPeerOption `yaml:"peers,omitempty"`
	RemoteDnsResolve    *bool                 `yaml:"remote-dns-resolve,omitempty"`
	Dns                 []string              `yaml:"dns,omitempty"`

	// tuic
	// Name                  string   `yaml:"name"`
	// Server                string   `yaml:"server"`
	// Port                  int      `yaml:"port"`
	Token *string `yaml:"token,omitempty"`
	// UUID                  string   `yaml:"uuid,omitempty"`
	// Password              string   `yaml:"password,omitempty"`
	// Ip                    string   `yaml:"ip,omitempty"`
	HeartbeatInterval *int `yaml:"heartbeat-interval,omitempty"`
	// ALPN                  []string `yaml:"alpn,omitempty"`
	ReduceRtt             *bool   `yaml:"reduce-rtt,omitempty"`
	RequestTimeout        *int    `yaml:"request-timeout,omitempty"`
	UdpRelayMode          *string `yaml:"udp-relay-mode,omitempty"`
	CongestionController  *string `yaml:"congestion-controller,omitempty"`
	DisableSni            *bool   `yaml:"disable-sni,omitempty"`
	MaxUdpRelayPacketSize *int    `yaml:"max-udp-relay-packet-size,omitempty"`

	// FastOpen             bool   `yaml:"fast-open,omitempty"`
	MaxOpenStreams *int `yaml:"max-open-streams,omitempty"`
	CWND           *int `yaml:"cwnd,omitempty"`
	// SkipCertVerify       bool   `yaml:"skip-cert-verify,omitempty"`
	// Fingerprint          string `yaml:"fingerprint,omitempty"`
	// CustomCA             string `yaml:"ca,omitempty"`
	// CustomCAString       string `yaml:"ca-str,omitempty"`
	// ReceiveWindowConn    int    `yaml:"recv-window-conn,omitempty"`
	// ReceiveWindow        int    `yaml:"recv-window,omitempty"`
	// DisableMTUDiscovery  bool   `yaml:"disable-mtu-discovery,omitempty"`
	MaxDatagramFrameSize *int `yaml:"max-datagram-frame-size,omitempty"`
	// SNI                  string `yaml:"sni,omitempty"`

	UDPOverStream        *bool `yaml:"udp-over-stream,omitempty"`
	UDPOverStreamVersion *int  `yaml:"udp-over-stream-version,omitempty"`
}

type RealityOptions struct {
	PublicKey string `yaml:"public-key"`
	ShortID   string `yaml:"short-id"`
}

type WireGuardPeerOption struct {
	Server       string   `yaml:"server"`
	Port         int      `yaml:"port"`
	Ip           *string  `yaml:"ip,omitempty"`
	Ipv6         *string  `yaml:"ipv6,omitempty"`
	PublicKey    *string  `yaml:"public-key,omitempty"`
	PreSharedKey *string  `yaml:"pre-shared-key,omitempty"`
	Reserved     []uint8  `yaml:"reserved,omitempty"`
	AllowedIPs   []string `yaml:"allowed-ips,omitempty"`
}

type PluginOptions struct {
	MaxEarlyData    *int               `yaml:"max-early-data,omitempty"`
	Mode            *string            `yaml:"mode,omitempty"`
	Host            *string            `yaml:"host,omitempty"`
	Path            *string            `yaml:"path,omitempty"`
	EarlyDataHeader *string            `yaml:"early-data-header-name,omitempty"`
	ServerName      *string            `yaml:"servername,omitempty"`
	TLS             *bool              `yaml:"tls,omitempty"`
	SkipCertVerify  *bool              `yaml:"skip-cert-verify,omitempty"`
	Mux             *bool              `yaml:"mux,omitempty"`
	Headers         *map[string]string `yaml:"headers,omitempty"`
	WSOptions       WSOptions          `yaml:"ws-opts,omitempty"`
	H2Options       HTTP2Options       `yaml:"h2-opts,omitempty"`
	HTTPOptions     HTTPOptions        `yaml:"http-opts,omitempty"`
	GRPCOptions     GRPCOptions        `yaml:"grpc-opts,omitempty"`
	ObfsOptions     ObfsOptions        `yaml:"obfs-opts,omitempty"`
}

type HTTPOptions struct {
	Method  *string              `yaml:"method,omitempty"`
	Path    []string             `yaml:"path,omitempty"`
	Headers *map[string][]string `yaml:"headers,omitempty"`
}

type HTTP2Options struct {
	Host []string `yaml:"host,omitempty"`
	Path *string  `yaml:"path,omitempty"`
}

type WSOptions struct {
	Path    *string            `yaml:"path,omitempty"`
	Headers *map[string]string `yaml:"headers,omitempty"`
}

type GRPCOptions struct {
	GrpcServiceName *string `yaml:"grpc-service-name,omitempty"`
}

type ObfsOptions struct {
	Mode *string `yaml:"mode,omitempty"`
	Host *string `yaml:"host,omitempty"`
}

type Proxies struct {
	*ordered.Map[string, Proxy]
}

func (o Proxies) MarshalYAML() (interface{}, error) {
	content := make([]yaml.Node, o.Len())
	n := &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     consts.YAMLNodeShortTagSeq,
		Content: make([]*yaml.Node, len(content)),
	}

	i := 0
	var err error
	o.Range(func(_ string, value Proxy) bool {
		_err := content[i].Encode(value)
		if _err != nil {
			err = _err
			return false
		}
		content[i].Style = yaml.FlowStyle
		n.Content[i] = &content[i]
		i++
		return true
	})
	return n, err
}

func (o *Proxies) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		return fmt.Errorf("proxy-provider get a invalid type: %s", value.ShortTag())
	}
	contentLen := len(value.Content)
	kv := value.Content
	o.Map = ordered.NewMap[string, Proxy](contentLen)
	for i := 0; i < contentLen; i++ {
		if kv[i].ShortTag() != consts.YAMLNodeShortTagMap {
			return fmt.Errorf("hosts get a invalid content type: %s", kv[i].ShortTag())
		}
		proxy := Proxy{}
		kv[i].Decode(&proxy)
		o.Set(proxy.Name, proxy)
	}
	return nil
}
