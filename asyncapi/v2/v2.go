package v2

type DocumentV2 struct {
	ExtendableV2
	ServersField map[string]ServerV2 `mapstructure:"servers"`
}

func (d DocumentV2) Servers() []ServerV2 {
	var servers []ServerV2
	for _, s := range d.ServersField {
		servers = append(servers, s)
	}

	return servers
}

type ServerV2 struct {
	ExtendableV2
	NameField      string                      `mapstructure:"name"`
	ProtocolField  string                      `mapstructure:"protocol"`
	URLField       string                      `mapstructure:"url"`
	VariablesField map[string]ServerVariableV2 `mapstructure:"variables"`
}

func (s ServerV2) Name() string {
	return s.NameField
}

func (s ServerV2) HasName() bool {
	return s.NameField != ""
}

func (s ServerV2) URL() string {
	return s.URLField
}

func (s ServerV2) HasURL() bool {
	return s.URLField != ""
}

func (s ServerV2) Protocol() string {
	return s.ProtocolField
}

func (s ServerV2) Variables() []ServerVariableV2 {
	var vars []ServerVariableV2
	for _, v := range s.VariablesField {
		vars = append(vars, v)
	}

	return vars
}

type ServerVariableV2 struct {
	ExtendableV2
	Default string   `mapstructure:"default"`
	Enum    []string `mapstructure:"enum"`
}

func (s ServerVariableV2) Name() string {
	panic("implement me")
}

func (s ServerVariableV2) HasName() string {
	panic("implement me")
}

func (s ServerVariableV2) DefaultValue() string {
	panic("implement me")
}

func (s ServerVariableV2) AllowedValues() []string {
	panic("implement me")
}

type ExtendableV2 struct {
	Raw map[string]interface{} `mapstructure:",remain"`
}

func (e ExtendableV2) HasExtension(name string) bool {
	_, ok := e.Raw[name]
	return ok
}

func (e ExtendableV2) Extension(name string) interface{} {
	return e.Raw[name]
}
