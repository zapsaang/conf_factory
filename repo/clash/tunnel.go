package clash

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/samber/lo"
	"github.com/zapsaang/conf_factory/utils/consts"
	"gopkg.in/yaml.v3"
)

type tunnel struct {
	Network []string `yaml:"network"`
	Address string   `yaml:"address"`
	Target  string   `yaml:"target"`
	Proxy   string   `yaml:"proxy"`
}

type Tunnel tunnel

func (t *Tunnel) BLength() int {
	l := 0
	for i, n := range t.Network {
		l += len(n)
		if i != len(t.Network)-1 {
			l++
		}
	}

	return l + len(t.Address) + len(t.Target) + len(t.Proxy) + 3
}

func (t Tunnel) MarshalYAML() (interface{}, error) {
	if !strings.Contains(t.Target, ":") {
		n := &yaml.Node{}
		_t := tunnel(t)
		return n, n.Encode(_t)
	}
	var buf bytes.Buffer
	buf.Grow(t.BLength())
	for i, n := range t.Network {
		buf.WriteString(n)
		if i != len(t.Network)-1 {
			buf.WriteByte('/')
		}
	}
	buf.WriteByte(',')
	buf.WriteString(t.Address)
	buf.WriteByte(',')
	buf.WriteString(t.Target)
	buf.WriteByte(',')
	buf.WriteString(t.Proxy)
	return buf.String(), nil
}

func (t *Tunnel) UnmarshalYAML(value *yaml.Node) error {
	switch value.ShortTag() {
	case consts.YAMLNodeShortTagStr:
		// parse udp/tcp,address,target,proxy
		parts := lo.Map(strings.Split(value.Value, ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		})
		if len(parts) != 4 {
			return fmt.Errorf("invalid tunnel config %s", value.Value)
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
		t.Network = network
		t.Address = address
		t.Target = target
		t.Proxy = parts[3]
		return nil
	case consts.YAMLNodeShortTagMap:
		_t := tunnel{}
		if err := value.Decode(&_t); err != nil {
			return err
		}
		t.Network = _t.Network
		t.Address = _t.Address
		t.Target = _t.Target
		t.Proxy = _t.Proxy
		return nil
	default:
		return fmt.Errorf("invalid tunnel config: type: %s, content: %s", value.ShortTag(), value.Value)
	}
}
