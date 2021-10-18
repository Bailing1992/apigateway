package config

type ProxyServiceConfig struct {
	PSM             string        `yaml:"PSM"`
	ServiceProtocol string        `yaml:"ServiceProtocol"`
	ServiceTimeout  TimeoutConfig `yaml:"ServiceTimeout"`
	RouteList       []RouteConfig `yaml:"RouteList"`
}

type RouteConfig struct {
	Path       string   `yaml:"Path"`
	Method     string   `yaml:"Method"`
	CommonArgs []string `yaml:"CommonArgs"`
	Timeout    int      `yaml:"Timeout"`
}
