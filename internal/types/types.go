package types

type HttpServer struct {
	Env string `env:"ENV" env-default:"DEVELOPMENT"`
	Port string `env:"PORT" env-required:"true"`
}

type Variables struct {
	DatabaseUrl string `env:"DATABASE_URL" env-required:"true"`
}
