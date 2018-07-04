package common

import "sync"

type SignalManager struct {
	connector map[Signal][]chan interface{}
	rWMutex   *sync.RWMutex
}

func (sm *SignalManager) Connect(signal Signal) chan interface{} {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	listener := make(chan interface{})
	if sm.connector == nil {
		sm.connector = make(map[Signal][]chan interface{})
	}
	if _, ok := sm.connector[signal]; ok {
		sm.connector[signal] = append(sm.connector[signal], listener)
	} else {
		sm.connector[signal] = []chan interface{}{listener}
	}
	return listener
}

func (sm *SignalManager) Disconnect(signal Signal, ch chan interface{}) bool {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	if _, ok := sm.connector[signal]; ok {
		for i, item := range sm.connector[signal] {
			if item == ch {
				closed := safeClose(item)
				sm.connector[signal] = append(sm.connector[signal][:i], sm.connector[signal][i+1:]...)
				return closed
			}
		}
	}
	return false
}

func (sm *SignalManager) DisconnectAll(signal Signal) bool {
	sm.rWMutex.Lock()
	defer sm.rWMutex.Unlock()

	if _, ok := sm.connector[signal]; ok {
		for _, item := range sm.connector[signal] {
			closed := safeClose(item)
			delete(sm.connector, signal)
			return closed
		}
	}
	return false
}

func (sm *SignalManager) Send(signal Signal, message interface{}) (success bool) {
	sm.rWMutex.RLock()
	defer sm.rWMutex.RUnlock()
	defer func() {
		if recover() != nil {
			success = false
		}
	}()

	if _, ok := sm.connector[signal]; ok {
		for _, handler := range sm.connector[signal] {
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

var ScrapySignal = &SignalManager{make(map[Signal][]chan interface{}), new(sync.RWMutex)}
