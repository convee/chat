package conf

type Config struct {
	Tcp   TcpConfig
	Redis map[string]RedisConfig
	Mysql map[string]MysqlConfig
}

type TcpConfig struct {
	Address string
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
