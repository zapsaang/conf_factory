package clash

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var configSample = `
port: 7890
socks-port: 7891
redir-port: 7892
tproxy-port: 7893
mixed-port: 7890
authentication:
 - "user1:pass1"
 - "user2:pass2"
allow-lan: false
bind-address: '*'
mode: rule
log-level: info
ipv6: false
external-controller: 127.0.0.1:9090
external-ui: folder
secret: ""
interface-name: en0
routing-mark: 6666
hosts:
  '*.clash.dev': 127.0.0.1
  '.dev': 127.0.0.1
  'alpha.clash.dev': '::1'
profile:
  store-selected: true
  store-fake-ip: false
dns:
  enable: false
  listen: 0.0.0.0:53
  ipv6: false 
  default-nameserver:
    - 114.114.114.114
    - 8.8.8.8
  enhanced-mode: fake-ip
  fake-ip-range: 198.18.0.1/16 
  use-hosts: true 
  search-domains: [local] 
  fake-ip-filter:
	- '*.lan'
    - localhost.ptlogin2.qq.com
  nameserver:
    - 114.114.114.114 
    - 8.8.8.8 
    - tls://dns.rubyfish.cn:853 
    - https://1.1.1.1/dns-query 
    - dhcp://en0 
    - '8.8.8.8#en0'
  fallback:
    - tcp://1.1.1.1
    - 'tcp://1.1.1.1#en0'
  fallback-filter:
    geoip: true
    geoip-code: CN
    ipcidr:
      - 240.0.0.0/4
    domain:
      - '+.google.com'
      - '+.facebook.com'
      - '+.youtube.com'
  nameserver-policy:
    'www.baidu.com': '114.114.114.114'
    '+.internal.crop.com': '10.0.0.1'
proxies:
  - name: "ss1"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    udp: true
  - name: "ss2"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    plugin: obfs
    plugin-opts:
      mode: tls 
      host: bing.com
  - name: "ss3"
    type: ss
    server: server
    port: 443
    cipher: chacha20-ietf-poly1305
    password: "password"
    plugin: v2ray-plugin
    plugin-opts:
      mode: websocket 
      tls: true 
      skip-cert-verify: true
      host: bing.com
      path: "/"
      mux: true
      headers:
        custom: value
  - name: "vmess"
    type: vmess
    server: server
    port: 443
    uuid: uuid
    alterId: 32
    cipher: auto
    udp: true
    tls: true
    skip-cert-verify: true
    servername: example.com 
    network: ws
    ws-opts:
      path: /path
      headers:
        Host: v2ray.com
      max-early-data: 2048
      early-data-header-name: Sec-WebSocket-Protocol
  - name: "vmess-h2"
    type: vmess
    server: server
    port: 443
    uuid: uuid
    alterId: 32
    cipher: auto
    network: h2
    tls: true
    h2-opts:
      host:
        - http.example.com
        - http-alt.example.com
      path: /
  - name: "vmess-http"
    type: vmess
    server: server
    port: 443
    uuid: uuid
    alterId: 32
    cipher: auto
    udp: true
    network: http
    http-opts:
      method: "GET"
      path:
        - '/'
        - '/video'
      headers:
        Connection:
          - keep-alive
  - name: vmess-grpc
    server: server
    port: 443
    type: vmess
    uuid: uuid
    alterId: 32
    cipher: auto
    network: grpc
    tls: true
    servername: example.com
    skip-cert-verify: true
    grpc-opts:
      grpc-service-name: "example"
  - name: "socks"
    type: socks5
    server: server
    port: 443
    username: username
    password: password
    tls: true
    skip-cert-verify: true
    udp: true
  - name: "http"
    type: http
    server: server
    port: 443
    username: username
    password: password
    tls: true 
    skip-cert-verify: true
    sni: custom.com
  - name: "snell"
    type: snell
    server: server
    port: 44046
    psk: yourpsk
    version: 2
    obfs-opts:
      mode: http 
      host: bing.com
  - name: "trojan"
    type: trojan
    server: server
    port: 443
    password: yourpsk
    udp: true
    sni: example.com 
    alpn:
      - h2
      - http/1.1
    skip-cert-verify: true
  - name: trojan-grpc
    server: server
    port: 443
    type: trojan
    password: "example"
    network: grpc
    sni: example.com
    skip-cert-verify: true
    udp: true
    grpc-opts:
      grpc-service-name: "example"
  - name: trojan-ws
    server: server
    port: 443
    type: trojan
    password: "example"
    network: ws
    sni: example.com
    skip-cert-verify: true
    udp: true
    ws-opts:
      path: /path
      headers:
        Host: example.com
  - name: "ssr"
    type: ssr
    server: server
    port: 443
    cipher: chacha20-ietf
    password: "password"
    obfs: tls1.2_ticket_auth
    protocol: auth_sha1_v4
    obfs-param: domain.tld
    protocol-param: "#"
    udp: true
proxy-groups:
  - name: "relay"
    type: relay
    proxies:
      - http
      - vmess
      - ss1
      - ss2
  - name: "auto"
    type: url-test
    proxies:
      - ss1
      - ss2
      - vmess1
    tolerance: 150
    lazy: true
    url: 'http://www.gstatic.com/generate_204'
    interval: 300
  - name: "fallback-auto"
    type: fallback
    proxies:
      - ss1
      - ss2
      - vmess1
    url: 'http://www.gstatic.com/generate_204'
    interval: 300
  - name: "load-balance"
    type: load-balance
    proxies:
      - ss1
      - ss2
      - vmess1
    url: 'http://www.gstatic.com/generate_204'
    interval: 300
    strategy: consistent-hashing 
  - name: Proxy
    type: select
    disable-udp: true
    filter: 'someregex'
    proxies:
      - ss1
      - ss2
      - vmess1
      - auto
  - name: en1
    type: select
    interface-name: en1
    routing-mark: 6667
    proxies:
      - DIRECT
  - name: UseProvider
    type: select
    use:
      - provider1
    proxies:
      - Proxy
      - DIRECT
proxy-providers:
  provider1:
    type: http
    url: "url"
    interval: 3600
    path: ./provider1.yaml
    health-check:
      enable: true
      interval: 600
      lazy: true
      url: http://www.gstatic.com/generate_204
  test:
    type: file
    path: /test.yaml
    health-check:
      enable: true
      interval: 36000
      url: http://www.gstatic.com/generate_204
tunnels:
  - tcp/udp,127.0.0.1:6553,114.114.114.114:53,proxy
  - tcp,127.0.0.1:6666,rds.mysql.com:3306,vpn
  - network: [tcp, udp]
    address: 127.0.0.1:7777
    target: target.com
    proxy: proxy
rules:
  - DOMAIN-SUFFIX,google.com,auto
  - DOMAIN-KEYWORD,google,auto
  - DOMAIN,google.com,auto
  - DOMAIN-SUFFIX,ad.com,REJECT
  - SRC-IP-CIDR,192.168.1.201/32,DIRECT
  - IP-CIDR,127.0.0.0/8,DIRECT
  - GEOIP,CN,DIRECT
  - DST-PORT,80,DIRECT
  - SRC-PORT,7777,DIRECT
  - RULE-SET,apple,REJECT 
  - MATCH,auto`

func TestUnmarshalConfig(t *testing.T) {

	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *RawConfig
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				buf: func() []byte {
					return []byte(configSample)
				}(),
			},
			want:    &RawConfig{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalRawConfig(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalRawConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotYAML, err := MarshalRawConfig(got)
			if err != nil {
				t.Errorf("MarshalRawConfig() error = %v", err)
				return
			}
			t.Logf("\n%s", string(gotYAML))
			ggot, err := UnmarshalRawConfig(gotYAML)
			if err != nil {
				t.Errorf("UnmarshalRawConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, ggot) {
				t.Errorf(cmp.Diff(got, ggot))
			}
		})
	}
}
