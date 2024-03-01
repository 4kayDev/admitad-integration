package di

import (
	"context"
	"fmt"
	"net"

	"github.com/4kayDev/admitad-integration/internal/pkg/clients/admitad"
	"github.com/4kayDev/admitad-integration/internal/pkg/rpc"
	"github.com/4kayDev/admitad-integration/internal/pkg/service"
	"github.com/4kayDev/admitad-integration/internal/pkg/storage/sql"
	"github.com/4kayDev/admitad-integration/internal/utils/config"
	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Container struct {
	ctx context.Context
	cfg *config.Config

	netListener *net.Listener
	grpcServer  *grpc.Server

	db                 *gorm.DB
	transactionManager trm.Manager
	storage            *sql.Storage

	admitadClient *admitad.Client

	service *service.Service

	rpcServer *rpc.Server
}

func NewContainer(ctx context.Context, cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) GetNetListener() *net.Listener {
	return get(&c.netListener, func() *net.Listener {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.cfg.Port))
		if err != nil {
			panic(err)
		}

		return &listener
	})
}

func (c *Container) GetGRPCServer() *grpc.Server {
	return get(&c.grpcServer, func() *grpc.Server {
		grpcServer := grpc.NewServer()
		reflection.Register(grpcServer)
		return grpcServer
	})
}

func (c *Container) GetPostgresDB() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewSQLite(c.cfg)
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.db))
	})
}

func (c *Container) GetService() *service.Service {
	return get(&c.service, func() *service.Service {
		return service.NewService(c.GetPostgresDB(), c.GetAdmitadClient())
	})
}

func (c *Container) GetRPCServer() *rpc.Server {
	return get(&c.rpcServer, func() *rpc.Server {
		return rpc.NewServer(c.GetService())
	})
}

func (c *Container) GetAdmitadClient() *admitad.Client {
	return get(&c.admitadClient, func() *admitad.Client {
		return admitad.New(&c.cfg.Admitad)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
