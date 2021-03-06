package bootstrap

import (
	"github.com/devlibx/gox-base"
	goxMessaging "github.com/devlibx/gox-messaging"
	"github.com/devlibx/gox-messaging/factory"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
)

type MessagingFactory goxMessaging.Factory

type messagingServiceImpl struct {
	gox.CrossFunction
	logger     *zap.Logger
	initDoOnce sync.Once
	goxMessaging.Factory
}

func NewMessagingFactory(lifecycle fx.Lifecycle, cf gox.CrossFunction, configuration *goxMessaging.Configuration) (MessagingFactory, error) {
	service := messagingServiceImpl{
		CrossFunction: cf,
		logger:        cf.Logger(),
		Factory:       factory.NewMessagingFactory(cf),
	}
	return &service, nil
}
