package appconfig

import (
	"flag"
	"fmt"
	"os"
)

type AppConfig struct {
	FilePath string
}

var Config = &AppConfig{}

func LoadInitConfig() {
	var filepath string
	ReadFlags(&filepath)
	Config.FilePath = filepath

}

func ReadFlags(filepath *string) {
	flag.StringVar(filepath, "filepath", "", "file absolute path config")
	flag.Parse()
	if *filepath == "" {
		fmt.Fprintln(os.Stderr, "Please input -filepath= <file path>")
		os.Exit(1)
	}
}
