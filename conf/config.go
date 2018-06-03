package conf

type Config struct {
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
	Ip          string
	Port        string
	Username    string
	Password    string
	Database    string
	Charset     string
	MaxIdle     int
	MaxOpen     int
	MaxLifetime int
}
