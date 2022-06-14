package config

type Config struct {
	Server
	Postgres
	Auth
}

type Server struct {
	BindAddr string `env:"BIND_ADDR" envDefault:":443"`
}

type Postgres struct {
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres123"`
	Host     string `env:"HOST" envDefault:"host.docker.internal:5432"`
	DBName   string `env:"DB_NAME" envDefault:"postgres"`
	SSL      string `env:"SSL" envDefault:"disable"`
}

type Auth struct {
	Username string `env:"USERNAME" envDefault:"service"`
	Password string `env:"PASSWORD" envDefault:"service"`
}

func New() *Config {
	return new(Config)
}
