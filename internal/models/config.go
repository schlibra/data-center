package models

type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Database string `yaml:"name"`
}

type serverConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Debug        bool   `yaml:"debug"`
	DefaultGroup int    `yaml:"default-group"`
}

type jwtConfig struct {
	Key    string `yaml:"key"`
	Issuer string `yaml:"issuer"`
}

type redisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"pass"`
	Index    int    `yaml:"index"`
}

type frpsWebConfig struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Token string `yaml:"token"`
}

type Config struct {
	MySQL   mysqlConfig   `yaml:"mysql"`
	Server  serverConfig  `yaml:"server"`
	Jwt     jwtConfig     `yaml:"jwt"`
	Redis   redisConfig   `yaml:"redis"`
	FrpsWeb frpsWebConfig `yaml:"frps-web"`
}
