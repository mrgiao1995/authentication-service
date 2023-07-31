package config

type AppConfigs struct {
	RedisConfig       *RedisConfig     `mapstructure:"redisConfig"`
	AuthServerConfig  *ApiServerConfig `mapstructure:"authServerConfig"`
	TokenServerConfig *ApiServerConfig `mapstructure:"tokenServerConfig"`
	DbConfig          *DbConfig        `mapstructure:"dbConfig"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Schema   string `mapstructure:"schema"`
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

type ApiServerConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	JwtSecret string `mapstructure:"jwtSecret"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
