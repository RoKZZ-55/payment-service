package config

type Config struct {
	Server
	Postgres
}

type Server struct {
	BindAddr string `env:"BIND_ADDR" envDefault:":8000"`
}

type Postgres struct {
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres123"`
	Host     string `env:"HOST" envDefault:"host.docker.internal:5432"`
	DBName   string `env:"DB_NAME" envDefault:"postgres"`
	SSL      string `env:"SSL" envDefault:"disable"`
}

func New() *Config {
	return new(Config)
}
