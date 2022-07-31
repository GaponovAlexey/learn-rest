package configToml

type TOML struct {
	port string `yaml:"port"`
	log  string `yaml:"log"`
}

func NewConfig() *TOML {
	return &TOML{
		port: ":3001",
		log:  "conf",
	}
}
