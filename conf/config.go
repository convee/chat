package conf

type Config struct {
	Redis map[string]RedisConfig
	Mysql map[string]MysqlConfig
}
type RedisConfig struct {
	Address string
}

type MysqlConfig struct {
	Username string
	Password string
	Ip       string
	Port     string
	DbName   string
}
