package clash

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/utils/consts"
)

type Hosts struct {
	*ordered.Map[string, string]
}

func (o Hosts) MarshalYAML() (interface{}, error) {
	content := make([]yaml.Node, o.Len()<<1)
	n := &yaml.Node{
		Kind:    yaml.MappingNode,
		Tag:     consts.YAMLNodeShortTagMap,
		Content: make([]*yaml.Node, len(content)),
	}

	i := 0
	o.Range(func(key, value string) bool {
		content[i].SetString(key)
		n.Content[i] = &content[i]
		i++
		content[i].SetString(value)
		n.Content[i] = &content[i]
		i++
		return true
	})
	return n, nil
}

func (o *Hosts) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.MappingNode {
		return fmt.Errorf("hosts get a invalid type: %s", value.ShortTag())
	}
	contentLen := len(value.Content)
	if contentLen&1 != 0 {
		return fmt.Errorf("hosts get a invalid content length: %d", contentLen)
	}
	kv := value.Content
	o.Map = ordered.NewMap[string, string](contentLen >> 1)
	for i := 0; i < contentLen; i += 2 {
		if kv[i].ShortTag() != consts.YAMLNodeShortTagStr || kv[i+1].ShortTag() != consts.YAMLNodeShortTagStr {
			return fmt.Errorf("hosts get a invalid content type: key: %s, value: %s", kv[i].ShortTag(), kv[i+1].ShortTag())
		}
		o.Set(kv[i].Value, kv[i+1].Value)
	}
	return nil
}
