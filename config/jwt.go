package config

type Jwt struct {
	ExpiresAt int64  `mapstructure:"expires_at" json:"expires_at" yaml:"expires_at"`
	Issuer    string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Subject   string `mapstructure:"subject" json:"subject" yaml:"subject"`
	Secretary string `mapstructure:"secretary" json:"secretary" yaml:"secretary"`
}
