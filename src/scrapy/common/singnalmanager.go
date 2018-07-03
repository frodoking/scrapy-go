package common

import "sync"

type SignalManager struct {
	connector map[string][]chan interface{}
	rWMutex *sync.RWMutex
}

func (sm *SignalManager) Connect(event string) chan interface{} {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	listener := make(chan interface{})
	if sm.connector == nil {
		sm.connector = make(map[string][]chan interface{})
	}
	if _, ok := sm.connector[event]; ok {
		sm.connector[event] = append(sm.connector[event], listener)
	} else {
		sm.connector[event] = []chan interface{}{listener}
	}
	return listener
}

func (sm *SignalManager) Disconnect(event string, ch chan interface{}) bool {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	if _, ok := sm.connector[event]; ok {
		for i, item := range sm.connector[event] {
			if item == ch {
				closed := safeClose(item)
				sm.connector[event] = append(sm.connector[event][:i], sm.connector[event][i+1:]...)
				return closed
			}
		}
	}
	return false
}

func (sm *SignalManager) DisconnectAll(event string) bool {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	if _, ok := sm.connector[event]; ok {
		for _, item := range sm.connector[event] {
			closed := safeClose(item)
			delete(sm.connector, event)
			return closed
		}
	}
	return false
}

func (sm *SignalManager) Send(event string, message string) (success bool) {
	sm.rWMutex.RLock()
	defer sm.rWMutex.RUnlock()
	defer func() {
		if recover() != nil {
			success = false
		}
	}()

	if _, ok := sm.connector[event]; ok {
		for _, handler := range sm.connector[event] {
			go func(handler chan interface{}) {
				sm.rWMutex.RLock()
				defer sm.rWMutex.RUnlock()

				if !isClosed(handler) {
					handler <- message
				}
			}(handler)
			return true
		}
	}
	return false
}

func isClosed(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func safeClose(ch chan interface{}) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}

var ScrapySignal = &SignalManager{make(map[string][]chan interface{}), new(sync.RWMutex)}
