package main

import (
	"github.com/vinoMamba/pharos-go/cmd"
	"github.com/vinoMamba/pharos-go/config"
)

func main() {
	config.LoadAppConfig()
	cmd.Execute()
}
