package clash

type Script struct {
	Code   string `yaml:"code,omitempty"`
	Engine string `yaml:"engine,omitempty"`
	// TODO change type to ordered.Map
	Shortcuts map[string]string `yaml:"shortcuts,omitempty"`
}
