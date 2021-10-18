package config

type TimeoutConfig struct {
	ConnectTimeout uint32 `yaml:"ConnectTimeout"`
	WriteTimeout   uint32 `yaml:"WriteTimeout"`
	ReadTimeout    uint32 `yaml:"ReadTimeout"`

	WriteBackTimeout uint32 `yaml:"WriteBackTimeout"`
}
