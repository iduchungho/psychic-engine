package cache

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"sync"
)

var store *session.Store
var lockSession = &sync.Mutex{}
var storeSlice map[string]*session.Store

// TODO: set multi session id

func GetSessionStore() *session.Store {
	if store == nil {
		lockSession.Lock()
		defer lockSession.Unlock()
		if store == nil {
			store = session.New()
			return store
		} else {
			return store
		}
	}
	return store
}

func GetSessionStoreSlice(name string) *session.Store {
	if storeSlice == nil {
		storeSlice = make(map[string]*session.Store)
	}
	if storeSlice[name] == nil {
		lockSession.Lock()
		defer lockSession.Unlock()
		if storeSlice[name] == nil {
			storeSlice[name] = session.New()
			return storeSlice[name]
		} else {
			return storeSlice[name]
		}
	}
	return storeSlice[name]
}
