package app

import (
	"fmt"
	"github.com/devlibx/gox-base"
	"github.com/devlibx/gox-base/metrics"
	"github.com/devlibx/gox-base/serialization"
	config2 "github.com/harishb2k/go-template-project/internal/config"
	"github.com/harishb2k/go-template-project/internal/handler"
	"github.com/harishb2k/go-template-project/pkg/clients"
	"github.com/harishb2k/go-template-project/pkg/core/bootstrap"
	noop "github.com/harishb2k/go-template-project/pkg/database/composit"
	"github.com/harishb2k/go-template-project/pkg/database/dynamodb"
	"github.com/harishb2k/go-template-project/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewServerCommand() *cobra.Command {

	appConfig := config2.ApplicationConfig{}
	err := serialization.ReadYaml("./config/app.yaml", &appConfig)
	if err != nil {
		panic(err)
	}

	var s server.Server
	injector := fx.New(

		// Main entry point for server
		fx.Invoke(NewApplicationEntryPoint),

		// Bootstrap dependencies - e.g. Gox HTTP, messaging, caching, ...
		bootstrap.Module,

		// Client module - these are the HTTP clients which this service can call
		clients.Module,

		noop.CompositeDatabaseModule,
		fx.Supply(&dynamodb.DynamoConfig{Region: "ap-south-1", Timeout: 1}),

		// Basic dependency - underlying server, CrossFunc, configs for application
		fx.Provide(server.NewServer),
		fx.Provide(NewCrossFunctionProvider),

		fx.Supply(appConfig.App),
		fx.Supply(&appConfig.ServerConfig),    // Gox-Http config which is needed by bootstrap module
		fx.Supply(&appConfig.MessagingConfig), // For messaging (if you don't use messaging pass a object with messaging enabled = false)
		fx.Supply(&appConfig.MetricConfig),    // For messaging (if you don't use messaging pass a object with metric config enabled = false)

		// Register all HTTP API handlers
		handler.RandomHandlerModule,
		handler.UserHandlerModule,

		// Instance of underlying server
		fx.Populate(&s),
	)

	return &cobra.Command{
		Use:   "gox",
		Short: "Small help code",
		Long:  `Long help code`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Your service running code here...")
			injector.Run()
		},
	}
}

func NewCrossFunctionProvider(metric metrics.Scope) gox.CrossFunction {
	var loggerConfig zap.Config
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, _ := loggerConfig.Build()
	return gox.NewCrossFunction(logger, metric)
}
