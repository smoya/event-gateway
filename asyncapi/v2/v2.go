package v2

import "github.com/asyncapi/event-gateway/asyncapi"

type Document struct {
	Extendable
	ServersField map[string]Server `mapstructure:"servers"`
}

func (d Document) Channels() []asyncapi.Channel {
	panic("implement me")
}

func (d Document) HasChannels() bool {
	panic("implement me")
}

func (d Document) ClientPublishableChannels() []asyncapi.Channel {
	panic("implement me")
}

func (d Document) ClientPublishableMessages() []asyncapi.Message {
	panic("implement me")
}

func (d Document) ClientSubscribableChannels() []asyncapi.Channel {
	panic("implement me")
}

func (d Document) ClientSubscribableMessages() []asyncapi.Message {
	panic("implement me")
}

func (d Document) Messages() []asyncapi.Message {
	panic("implement me")
}

func (d Document) Servers() []asyncapi.Server {
	var servers []asyncapi.Server
	for _, s := range d.ServersField {
		servers = append(servers, s)
	}

	return servers
}

func (d Document) HasServers() bool {
	return len(d.ServersField) > 0
}

type Channel struct {
}

func (c Channel) HasExtension(name string) bool {
	panic("implement me")
}

func (c Channel) Extension(name string) interface{} {
	panic("implement me")
}

func (c Channel) IDField() string {
	panic("implement me")
}

func (c Channel) ID() string {
	panic("implement me")
}

func (c Channel) Path() string {
	panic("implement me")
}

func (c Channel) Description() string {
	panic("implement me")
}

func (c Channel) HasDescription() bool {
	panic("implement me")
}

func (c Channel) Parameters() []asyncapi.ChannelParameter {
	panic("implement me")
}

func (c Channel) HasParameters() bool {
	panic("implement me")
}

func (c Channel) Operations() []asyncapi.Operation {
	panic("implement me")
}

func (c Channel) Messages() []asyncapi.Message {
	panic("implement me")
}

type ChannelParameter struct {
}

func (c ChannelParameter) HasExtension(name string) bool {
	panic("implement me")
}

func (c ChannelParameter) Extension(name string) interface{} {
	panic("implement me")
}

func (c ChannelParameter) IDField() string {
	panic("implement me")
}

func (c ChannelParameter) ID() string {
	panic("implement me")
}

func (c ChannelParameter) Name() string {
	panic("implement me")
}

func (c ChannelParameter) Description() string {
	panic("implement me")
}

func (c ChannelParameter) Schema() asyncapi.Schema {
	panic("implement me")
}

type Operation struct {
}

func (o Operation) HasExtension(name string) bool {
	panic("implement me")
}

func (o Operation) Extension(name string) interface{} {
	panic("implement me")
}

func (o Operation) IDField() string {
	panic("implement me")
}

func (o Operation) ID() string {
	panic("implement me")
}

func (o Operation) Description() string {
	panic("implement me")
}

func (o Operation) HasDescription() bool {
	panic("implement me")
}

func (o Operation) IsApplicationPublishing() bool {
	panic("implement me")
}

func (o Operation) IsApplicationSubscribing() bool {
	panic("implement me")
}

func (o Operation) IsClientPublishing() bool {
	panic("implement me")
}

func (o Operation) IsClientSubscribing() bool {
	panic("implement me")
}

func (o Operation) Messages() []asyncapi.Message {
	panic("implement me")
}

type Message struct {
	Extendable
	asyncapi.Identifiable
}

func (m Message) HasExtension(name string) bool {
	panic("implement me")
}

func (m Message) Extension(name string) interface{} {
	panic("implement me")
}

func (m Message) IDField() string {
	panic("implement me")
}

func (m Message) ID() string {
	panic("implement me")
}

func (m Message) UID() string {
	panic("implement me")
}

func (m Message) Name() string {
	panic("implement me")
}

func (m Message) Title() string {
	panic("implement me")
}

func (m Message) HasTitle() bool {
	panic("implement me")
}

func (m Message) Description() string {
	panic("implement me")
}

func (m Message) HasDescription() bool {
	panic("implement me")
}

func (m Message) Summary() string {
	panic("implement me")
}

func (m Message) ContentType() string {
	panic("implement me")
}

func (m Message) Payload() asyncapi.Schema {
	panic("implement me")
}

type Schema struct {
}

func (s Schema) HasExtension(name string) bool {
	panic("implement me")
}

func (s Schema) Extension(name string) interface{} {
	panic("implement me")
}

func (s Schema) AdditionalItems() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) AdditionalProperties() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) AllOf() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) AnyOf() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) CircularProps() []string {
	panic("implement me")
}

func (s Schema) Const() interface{} {
	panic("implement me")
}

func (s Schema) Contains() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) ContentEncoding() string {
	panic("implement me")
}

func (s Schema) ContentMediaType() string {
	panic("implement me")
}

func (s Schema) Default() interface{} {
	panic("implement me")
}

func (s Schema) Definitions() map[string]asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Dependencies() map[string]asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Deprecated() bool {
	panic("implement me")
}

func (s Schema) Discriminator() string {
	panic("implement me")
}

func (s Schema) Else() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Enum() []interface{} {
	panic("implement me")
}

func (s Schema) Examples() []interface{} {
	panic("implement me")
}

func (s Schema) ExclusiveMaximum() float64 {
	panic("implement me")
}

func (s Schema) ExclusiveMinimum() float64 {
	panic("implement me")
}

func (s Schema) Format() string {
	panic("implement me")
}

func (s Schema) HasCircularProps() bool {
	panic("implement me")
}

func (s Schema) If() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) IsCircular() bool {
	panic("implement me")
}

func (s Schema) Items() []asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Maximum() float64 {
	panic("implement me")
}

func (s Schema) MaxItems() float64 {
	panic("implement me")
}

func (s Schema) MaxLength() float64 {
	panic("implement me")
}

func (s Schema) MaxProperties() float64 {
	panic("implement me")
}

func (s Schema) Minimum() float64 {
	panic("implement me")
}

func (s Schema) MinItems() float64 {
	panic("implement me")
}

func (s Schema) MinLength() float64 {
	panic("implement me")
}

func (s Schema) MinProperties() float64 {
	panic("implement me")
}

func (s Schema) MultipleOf() float64 {
	panic("implement me")
}

func (s Schema) Not() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) OneOf() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Pattern() string {
	panic("implement me")
}

func (s Schema) PatternProperties() map[string]asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Properties() map[string]asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Property(name string) asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Name() string {
	panic("implement me")
}

func (s Schema) PropertyNames() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) ReadOnly() bool {
	panic("implement me")
}

func (s Schema) Required() string {
	panic("implement me")
}

func (s Schema) Then() asyncapi.Schema {
	panic("implement me")
}

func (s Schema) Title() string {
	panic("implement me")
}

func (s Schema) Type() []string {
	panic("implement me")
}

func (s Schema) Uid() string {
	panic("implement me")
}

func (s Schema) UniqueItems() bool {
	panic("implement me")
}

func (s Schema) WriteOnly() bool {
	panic("implement me")
}

type Server struct {
	Extendable
	NameField        string                    `mapstructure:"name"`
	DescriptionField string                    `mapstructure:"description"`
	ProtocolField    string                    `mapstructure:"protocol"`
	URLField         string                    `mapstructure:"url"`
	VariablesField   map[string]ServerVariable `mapstructure:"variables"`
}

func (s Server) Variables() []asyncapi.ServerVariable {
	var vars []asyncapi.ServerVariable
	for _, v := range s.VariablesField {
		vars = append(vars, v)
	}

	return vars
}

func (s Server) IDField() string {
	return "name"
}

func (s Server) ID() string {
	return s.NameField
}

func (s Server) Name() string {
	return s.NameField
}

func (s Server) HasName() bool {
	return s.NameField != ""
}

func (s Server) Description() string {
	return s.DescriptionField
}

func (s Server) HasDescription() bool {
	return s.DescriptionField != ""
}

func (s Server) URL() string {
	return s.URLField
}

func (s Server) HasURL() bool {
	return s.URLField != ""
}

func (s Server) Protocol() string {
	return s.ProtocolField
}

func (s Server) HasProtocol() bool {
	return s.ProtocolField != ""
}

type ServerVariable struct {
	Extendable
	NameField string   `mapstructure:"name"`
	Default   string   `mapstructure:"default"`
	Enum      []string `mapstructure:"enum"`
}

func (s ServerVariable) IDField() string {
	return "name"
}

func (s ServerVariable) ID() string {
	return s.NameField
}

func (s ServerVariable) Name() string {
	return s.NameField
}

func (s ServerVariable) HasName() bool {
	return s.NameField != ""
}

func (s ServerVariable) DefaultValue() string {
	return s.Default
}

func (s ServerVariable) AllowedValues() []string {
	return s.Enum
}

type Extendable struct {
	Raw map[string]interface{} `mapstructure:",remain"`
}

func (e Extendable) HasExtension(name string) bool {
	_, ok := e.Raw[name]
	return ok
}

func (e Extendable) Extension(name string) interface{} {
	return e.Raw[name]
}
