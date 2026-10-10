package models

type Config struct {
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Database string `yaml:"name"`
	} `yaml:"mysql"`
	Server struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Debug bool   `yaml:"debug"`
		SSL   struct {
			Enable   bool   `yaml:"enable"`
			Port     int    `yaml:"port"`
			CertFile string `yaml:"cert"`
			KeyFile  string `yaml:"key"`
		} `yaml:"ssl"`
	} `yaml:"server"`
	Jwt struct {
		Key    string `yaml:"key"`
		Issuer string `yaml:"issuer"`
	} `yaml:"jwt"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"pass"`
		Index    int    `yaml:"index"`
	} `yaml:"redis"`
	FrpsWeb struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Token string `yaml:"token"`
	} `yaml:"frps-web"`
	IKuai struct {
		Address string `yaml:"address"`
		Key     string `yaml:"key"`
	} `yaml:"ikuai"`
	Frps struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"frps"`
}
