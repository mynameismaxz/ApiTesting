package config

type Config struct{
	DB *DBConfig
}

type DBConfig struct{
	Dialect string
	Username string
	Password string
	Name string
	Charset string
}

func GetCon