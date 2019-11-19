package config

import (
	"os"
	"os/user"
	"strings"
)

type Config struct {
	Identity string
	Datastore Datastore
}

type Identity struct {
	PeerId string
}

type Datastore struct {
	Type string
	Path string
}

var defaultConfigFilePath = "~/.go-monesign/config"
var defaultConfigFile = `{
	"identity" : {},
	"datastore" : {
		"type": "leveldb",
		"path": "~/.go-monesign/datastore"
	}
}
`

func LoadConfig(filename string) (*Config, error) {
	if len(filename) == 0 {
		filename = defaultConfigFilePath
	}

	if strings.HasPrefix(filename, "~/") {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}

		dir := usr.HomeDir + "/"
		filename = strings.Replace(filename, "~/", dir, 1)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		WriteFile(filename, []byte(defaultConfigFile))
	}

	var cfg Config
	err := ReadConfigFile(filename, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
