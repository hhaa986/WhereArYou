package config

import "github.com/BurntSushi/toml"

type TomlStruct struct {
	Title string
	DB    dbInfo `toml:"DB"`
}

type dbInfo struct {
	Info string `toml:"info"`
}

func SetToml() TomlStruct {
	var conf TomlStruct
	_, err := toml.DecodeFile("config.toml", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
