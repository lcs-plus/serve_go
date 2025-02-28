package config

type Elastic struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
