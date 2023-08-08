package configuration

type Object struct {
	Redis struct {
		Address  string
		DB       int
		Password string
	}
	ServerHostPort string
}
