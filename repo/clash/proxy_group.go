package clash

import (
	"fmt"

	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/utils/consts"
	"gopkg.in/yaml.v3"
)

type ProxyGroup struct {
	Name           string   `yaml:"name"`
	Type           string   `yaml:"type"`
	Proxies        []string `yaml:"proxies,omitempty"`
	Use            []string `yaml:"use,omitempty"`
	URL            *string  `yaml:"url,omitempty"`
	Tolerance      *int     `yaml:"tolerance,omitempty"`
	Strategy       *string  `yaml:"strategy,omitempty"`
	Interval       *int     `yaml:"interval,omitempty"`
	Lazy           *bool    `yaml:"lazy,omitempty"`
	DisableUDP     *bool    `yaml:"disable-udp,omitempty"`
	Filter         *string  `yaml:"filter,omitempty"`
	ExcludeFilter  *string  `yaml:"exclude-filter,omitempty"`
	ExcludeType    *string  `yaml:"exclude-type,omitempty"`
	ExpectedStatus *string  `yaml:"expected-status,omitempty"`

	// basic
	TFO         *bool   `yaml:"tfo,omitempty"`
	MPTCP       *bool   `yaml:"mptcp,omitempty"`
	Interface   *string `yaml:"interface-name,omitempty"`
	RoutingMark *int    `yaml:"routing-mark,omitempty"`
	IPVersion   *string `yaml:"ip-version,omitempty"`
	DialerProxy *string `yaml:"dialer-proxy,omitempty"`
}

type ProxyGroups struct {
	*ordered.Map[string, ProxyGroup]
}

func (o ProxyGroups) MarshalYAML() (interface{}, error) {
	content := make([]yaml.Node, o.Len())
	n := &yaml.Node{
		Kind:    yaml.SequenceNode,
		Tag:     consts.YAMLNodeShortTagSeq,
		Content: make([]*yaml.Node, len(content)),
	}

	i := 0
	var err error
	o.Range(func(_ string, value ProxyGroup) bool {
		_err := content[i].Encode(value)
		if _err != nil {
			err = _err
			return false
		}
		n.Content[i] = &content[i]
		i++
		return true
	})
	return n, err
}

func (o *ProxyGroups) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		return fmt.Errorf("proxy-provider get a invalid type: %s", value.ShortTag())
	}
	contentLen := len(value.Content)
	kv := value.Content
	o.Map = ordered.NewMap[string, ProxyGroup](contentLen)
	for i := 0; i < contentLen; i++ {
		if kv[i].ShortTag() != consts.YAMLNodeShortTagMap {
			return fmt.Errorf("hosts get a invalid content type: %s", kv[i].ShortTag())
		}
		group := ProxyGroup{}
		kv[i].Decode(&group)
		o.Set(group.Name, group)
	}
	return nil
}
