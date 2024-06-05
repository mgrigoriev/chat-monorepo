package main

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/repository/chatmessages_storage"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server"
	middleware_errors "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server/middleware/errors"
	middleware_logging "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server/middleware/logging"
	middleware_metrics "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server/middleware/metrics"
	middleware_recovery "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server/middleware/recovery"
	middleware_tracing "github.com/mgrigoriev/chat-monorepo/chatmessages/internal/server/middleware/tracing"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/internal/usecases"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/logger"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/postgres"
	jaeger_tracing "github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/tracing"
	"github.com/mgrigoriev/chat-monorepo/chatmessages/pkg/transaction_manager"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const grpcPort = "9090"
const httpPort = "8080"
const internalServerPort = "8888"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	logger.SetLevel(zapcore.DebugLevel)

	logger.Info(ctx, "start app init")
	serviceName := os.Getenv("JAEGER_SERVICE_NAME")
	if err := jaeger_tracing.Init(serviceName); err != nil {
		logger.Fatal(ctx, err)
	}

	// repository
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "0.0.0.0"
	}

	DSN := fmt.Sprintf("user=mikhail host=%s port=5432 dbname=chatmessages pool_max_conns=10", dbHost)

	pool, err := postgres.NewConnectionPool(ctx, DSN,
		postgres.WithMaxConnIdleTime(5*time.Minute),
		postgres.WithMaxConnLifeTime(time.Hour),
		postgres.WithMaxConnectionsCount(10),
		postgres.WithMinConnectionsCount(5),
	)
	if err != nil {
		logger.Fatal(ctx, err)
	}

	txManager := transaction_manager.New(pool)
	storage := chatmessages_storage.New(txManager)

	// services

	// ...

	// usecases

	uc := usecases.NewUsecase(usecases.Deps{
		ChatMessagesStorage: storage,
		TransactionManager:  txManager,
	})

	serverCfg := server.Config{
		GrpcPort:           grpcPort,
		HttpPort:           httpPort,
		InternalServerPort: internalServerPort,
		ChainUnaryInterceptors: []grpc.UnaryServerInterceptor{
			// https://github.com/grpc-ecosystem/go-grpc-middleware?tab=readme-ov-file#middleware
			grpc_opentracing.OpenTracingServerInterceptor(opentracing.GlobalTracer(), grpc_opentracing.LogPayloads()), // Order matters e.g. tracing interceptor have to create span first for the later exemplars to work.
			middleware_logging.LogErrorUnaryInterceptor(),
			middleware_tracing.DebugOpenTracingUnaryServerInterceptor(true, true), // расширение для grpc_opentracing.OpenTracingServerInterceptor
			middleware_metrics.MetricsUnaryInterceptor(),
			middleware_recovery.RecoverUnaryInterceptor(),
		},
		UnaryInterceptors: []grpc.UnaryServerInterceptor{
			middleware_errors.ErrorsUnaryInterceptor(),
		},
	}
	serverDeps := server.Deps{Usecase: uc}
	srv, err := server.NewServer(serverCfg, serverDeps)
	if err != nil {
		logger.Fatalf(ctx, "failed to create server: %v", err)
	}

	srv.Start(ctx)
}
