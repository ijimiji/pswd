package config

import (
	"fmt"
	"os"
)

type config struct {
	HomeFolder string
}

var Default = config{
	HomeFolder: fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".config/pswd"),
}

func Get() config {
	return Default
}
