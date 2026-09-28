package models

type mysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Database string `yaml:"name"`
}

type serverSSLConfig struct {
	Enable   bool   `yaml:"enable"`
	Port     int    `yaml:"port"`
	CertFile string `yaml:"cert"`
	KeyFile  string `yaml:"key"`
}

type serverConfig struct {
	Host         string          `yaml:"host"`
	Port         int             `yaml:"port"`
	Debug        bool            `yaml:"debug"`
	DefaultGroup int             `yaml:"default-group"`
	SSL          serverSSLConfig `yaml:"ssl"`
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
type ikuaiConfig struct {
	Address string `yaml:"address"`
	Key     string `yaml:"key"`
}
type Config struct {
	MySQL   mysqlConfig   `yaml:"mysql"`
	Server  serverConfig  `yaml:"server"`
	Jwt     jwtConfig     `yaml:"jwt"`
	Redis   redisConfig   `yaml:"redis"`
	FrpsWeb frpsWebConfig `yaml:"frps-web"`
	IKuai   ikuaiConfig   `yaml:"ikuai"`
}
