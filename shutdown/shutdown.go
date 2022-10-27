package shutdown

import (
    "github.com/keweegen/notification/logger"
    "os"
    "os/signal"
    "sync"
)

type (
    Handler func() error

    handler struct {
        name    string
        handler Handler
    }
)

type Shutdown struct {
    ch       chan struct{}
    logger   logger.Logger
    handlers []handler
    mx       *sync.Mutex
}

var S *Shutdown

func New(logger logger.Logger) *Shutdown {
    S = new(Shutdown)
    S.ch = make(chan struct{})
    S.handlers = make([]handler, 0)
    S.logger = logger.With("entity", "shutdown")
    S.mx = new(sync.Mutex)

    return S
}

func (s *Shutdown) AddHandler(name string, h Handler) *Shutdown {
    s.mx.Lock()
    s.handlers = append(s.handlers, handler{name: name, handler: h})
    s.mx.Unlock()
    s.logger.Debug("Add handler %s", name)
    return s
}

func (s *Shutdown) Listen() {
    go s.listen()
    s.logger.Debug("Shutdown is listen...")
}

func (s *Shutdown) ReadCh() {
    <-s.ch
    s.logger.Info("Service has been stopped!")
}

func (s *Shutdown) listen() {
    sigint := make(chan os.Signal, 1)
    signal.Notify(sigint, os.Interrupt)
    <-sigint
    s.runHandlers()
    close(s.ch)
}

func (s *Shutdown) runHandlers() {
    s.mx.Lock()
    for i := len(s.handlers) - 1; i >= 0; i-- {
        element := s.handlers[i]

        s.logger.Debug("Run handler %s", element.name)

        if err := element.handler(); err != nil {
            s.logger.Error("Failed shutdown run handler %s: %s", element.name, err)
        }
    }
    s.mx.Unlock()
}
