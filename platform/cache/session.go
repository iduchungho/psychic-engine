package cache

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"sync"
)

var store *session.Store
var lockSession = &sync.Mutex{}

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
