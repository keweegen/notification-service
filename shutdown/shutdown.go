package shutdown

import (
	"context"
	"github.com/keweegen/notification/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var signals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

type (
	HandlerType func() error

	handler struct {
		name    string
		handler HandlerType
	}
)

type Shutdown struct {
	logger    logger.Logger
	ctx       context.Context
	ctxCancel context.CancelFunc
	handlers  []handler
	mx        *sync.Mutex
}

func New(logger logger.Logger) *Shutdown {
	ctx, cancel := signal.NotifyContext(context.Background(), signals...)

	s := &Shutdown{
		logger:    logger.With("entity", "shutdown"),
		ctx:       ctx,
		ctxCancel: cancel,
		handlers:  make([]handler, 0),
		mx:        new(sync.Mutex),
	}

	return s
}

func (s *Shutdown) AddHandler(name string, h HandlerType) *Shutdown {
	s.mx.Lock()
	s.handlers = append(s.handlers, handler{name: name, handler: h})
	s.mx.Unlock()

	s.logger.Debug("add handler", "name", name)

	return s
}

func (s *Shutdown) Context() context.Context {
	return s.ctx
}

func (s *Shutdown) ListenContextDone() {
	for {
		<-s.ctx.Done()
		s.ctxCancel()
		s.runHandlers()
		time.Sleep(5 * time.Second)
		return
	}
}

func (s *Shutdown) runHandlers() {
	for i := len(s.handlers) - 1; i >= 0; i-- {
		element := s.handlers[i]

		s.logger.Debug("run handler", "name", element.name)

		if err := element.handler(); err != nil {
			s.logger.Debug("failed run handler",
				"name", element.name,
				"error", err)
		}
	}
}
