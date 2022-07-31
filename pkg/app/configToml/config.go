package configToml

type TOML struct {
	port string `toml:"port"`
	log  string `toml:"log"`
}

func NewConfig() *TOML {
	return &TOML{
		port: ":3001",
		log:  "conf",
	}
}
