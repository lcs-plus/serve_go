package config

type Httpapp struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
}
