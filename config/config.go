package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Hostname string
	Dbname   string
	Charset  string
	Port     string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "root",
			Hostname: "localhost",
			Dbname:   "test",
			Charset:  "utf8",
			Port:     "3306",
		},
	}
}
