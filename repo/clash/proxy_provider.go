package clash

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/utils/consts"
)

type HealthCheck struct {
	Enable         bool    `yaml:"enable"`
	URL            string  `yaml:"url"`
	Interval       int     `yaml:"interval"`
	Lazy           *bool   `yaml:"lazy,omitempty"`
	ExpectedStatus *string `yaml:"expected-status,omitempty"`
}

type ProxyProvider struct {
	Type          string      `yaml:"type"`
	Path          *string     `yaml:"path,omitempty"`
	URL           *string     `yaml:"url,omitempty"`
	Interval      *int        `yaml:"interval,omitempty"`
	Filter        *string     `yaml:"filter,omitempty"`
	ExcludeFilter *string     `yaml:"exclude-filter,omitempty"`
	ExcludeType   *string     `yaml:"exclude-type,omitempty"`
	DialerProxy   *string     `yaml:"dialer-proxy,omitempty"`
	HealthCheck   HealthCheck `yaml:"health-check,omitempty"`
}

type ProxyProviders struct {
	*ordered.Map[string, ProxyProvider]
}

func (o ProxyProviders) MarshalYAML() (interface{}, error) {
	content := make([]yaml.Node, o.Len()<<1)
	n := &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     consts.YAMLNodeShortTagMap,
		Content: make([]*yaml.Node, len(content)),
	}

	i := 0
	var err error
	o.Range(func(key string, value ProxyProvider) bool {
		content[i].SetString(key)
		n.Content[i] = &content[i]
		i++
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

func (o *ProxyProviders) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("proxy-provider get a invalid type: %s", value.ShortTag())
	}
	contentLen := len(value.Content)
	if contentLen&1 != 0 {
		return fmt.Errorf("proxy-provider get a invalid content length: %d", contentLen)
	}
	kv := value.Content
	o.Map = ordered.NewMap[string, ProxyProvider](contentLen >> 1)
	for i := 0; i < contentLen; i += 2 {
		if kv[i].ShortTag() != consts.YAMLNodeShortTagStr || kv[i+1].ShortTag() != consts.YAMLNodeShortTagMap {
			return fmt.Errorf("hosts get a invalid content type: key: %s, value: %s", kv[i].ShortTag(), kv[i+1].ShortTag())
		}
		provider := ProxyProvider{}
		kv[i+1].Decode(&provider)
		o.Set(kv[i].Value, provider)
	}
	return nil
}
