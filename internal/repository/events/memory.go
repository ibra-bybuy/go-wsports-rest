package events

import (
	"context"
	"sync"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

type Memory struct {
	sync.RWMutex
	data *[]model.Event
}

func NewMemory() *Memory {
	return &Memory{data: &[]model.Event{}}
}

func (m *Memory) Add(ctx context.Context, events *[]model.Event) bool {
	m.RLock()
	defer m.RUnlock()
	*m.data = append(*m.data, *events...)
	return true
}
