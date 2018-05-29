package conf

type Config struct {
	Redis RedisConfig
	Mysql MysqlConfig
}
type RedisConfig struct {
	Name    string
	Address string
}

type MysqlConfig struct {
	Name     string
	Username string
	Password string
	Ip       string
	Port     string
	DbName   string
}
