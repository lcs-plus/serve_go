package config

type Configuration struct {
	HttpApp  Httpapp  `mapstructure:"httpapp" json:"httpapp" yaml:"httpapp"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt"`
	Redis    Redis    `mapstructure:"redis" json:"redis"`
	Elastic  Elastic  `mapstructure:"elastic" json:"elastic"`
}
