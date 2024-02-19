package main

import (
	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad"
	"github.com/4kayDev/admitad-integration/internal/utils/config"
	"github.com/4kayDev/admitad-integration/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	client := admitad.New(&config.Admitad)
}
