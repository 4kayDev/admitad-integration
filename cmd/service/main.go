package main

import (
	"context"
	"fmt"

	"github.com/4kayDev/admitad-integration/internal/di"
	pb "github.com/4kayDev/admitad-integration/internal/generated/proto/admitad_integration"
	"github.com/4kayDev/admitad-integration/internal/utils/config"
	"github.com/4kayDev/admitad-integration/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	config := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.NewContainer(context.Background(), config)
	pb.RegisterAdmitadIntegrationServer(container.GetGRPCServer(), container.GetRPCServer())
	
	err := container.GetGRPCServer().Serve(*container.GetNetListener())
	if err != nil {
		fmt.Println(err, "Error while serving grpcServer", "SERVICE", "main")
		panic(err)
	}
}
