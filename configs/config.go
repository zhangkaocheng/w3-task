package configs

var JWTConf *JWTConfig

type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Server ServerConfig `yaml:"server"`
	Log    LogConfig    `yaml:"log"`
	Gin    GinConfig    `yaml:"gin"`
	JWT    JWTConfig    `yaml:"jwt"`
}

type MySQLConfig struct {
	DSN             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime_hour"`
	LogMode         bool   `yaml:"log_mode"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type LogConfig struct {
	Level   string `yaml:"level"`
	FilePat string `yaml:"file_path"`
}

type GinConfig struct {
	Debug bool `yaml:"debug"`
}

type JWTConfig struct {
	Secret     string `yaml:"secret"`
	ExpireHour int    `yaml:"expire_hour"`
}
