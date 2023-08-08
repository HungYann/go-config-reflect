package config

import (
	"os"
	"testing"
)

func Test_parse(t *testing.T) {
	file, err := os.ReadFile("./file/config.json")
	if err != nil {
		panic(err)
	}

	parse(file)
}
