package app

type Config struct {
	App      `yaml:"app"`
	Database `yaml:"database"`
}

type Database struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type App struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
