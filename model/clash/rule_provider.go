package clash

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/utils/consts"
)

type RuleProvider struct {
	Type     string  `provider:"type"`
	Behavior string  `provider:"behavior"`
	Path     *string `provider:"path,omitempty"`
	URL      *string `provider:"url,omitempty"`
	Format   *string `provider:"format,omitempty"`
	Interval *int    `provider:"interval,omitempty"`
}

type RuleProviders struct {
	*ordered.Map[string, RuleProvider]
}

func (o RuleProviders) MarshalYAML() (interface{}, error) {
	content := make([]yaml.Node, o.Len()<<1)
	n := &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     consts.YAMLNodeShortTagMap,
		Content: make([]*yaml.Node, len(content)),
	}

	i := 0
	var err error
	o.Range(func(key string, value RuleProvider) bool {
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

func (o *RuleProviders) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("proxy-provider get a invalid type: %s", value.ShortTag())
	}
	contentLen := len(value.Content)
	if contentLen&1 != 0 {
		return fmt.Errorf("proxy-provider get a invalid content length: %d", contentLen)
	}
	kv := value.Content
	o.Map = ordered.NewMap[string, RuleProvider](contentLen >> 1)
	for i := 0; i < contentLen; i += 2 {
		if kv[i].ShortTag() != consts.YAMLNodeShortTagStr || kv[i+1].ShortTag() != consts.YAMLNodeShortTagMap {
			return fmt.Errorf("hosts get a invalid content type: key: %s, value: %s", kv[i].ShortTag(), kv[i+1].ShortTag())
		}
		provider := RuleProvider{}
		kv[i+1].Decode(&provider)
		o.Set(kv[i].Value, provider)
	}
	return nil
}
