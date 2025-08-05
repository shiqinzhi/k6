package js

import (
	"context"

	"github.com/grafana/sobek"
	"github.com/shiqinzhi/k6/internal/event"
	"github.com/shiqinzhi/k6/internal/js/eventloop"
	"github.com/shiqinzhi/k6/js/common"
	"github.com/shiqinzhi/k6/lib"
)

type events struct {
	global, local *event.System
}

type moduleVUImpl struct {
	ctx       context.Context
	initEnv   *common.InitEnvironment
	state     *lib.State
	runtime   *sobek.Runtime
	eventLoop *eventloop.EventLoop
	events    events
}

func (m *moduleVUImpl) Context() context.Context {
	return m.ctx
}

func (m *moduleVUImpl) Events() common.Events {
	return common.Events{Global: m.events.global, Local: m.events.local}
}

func (m *moduleVUImpl) InitEnv() *common.InitEnvironment {
	return m.initEnv
}

func (m *moduleVUImpl) State() *lib.State {
	return m.state
}

func (m *moduleVUImpl) Runtime() *sobek.Runtime {
	return m.runtime
}

func (m *moduleVUImpl) RegisterCallback() func(func() error) {
	return m.eventLoop.RegisterCallback()
}
