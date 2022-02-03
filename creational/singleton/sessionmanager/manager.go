package sessionmanager

import (
	"sync"

	"github.com/google/uuid"
)

type SessionStore struct {
	store map[string]string
}

var sessionOnce sync.Once
var session *SessionStore

func GetSessionStore() *SessionStore {
	sessionOnce.Do(func() {
		session = new(SessionStore)
		session.store = make(map[string]string)
	})

	return session
}

func (s *SessionStore) Put(id string) {
	_, exist := s.store[id]
	if !exist {
		uuid := uuid.New().String()
		s.store[id] = uuid
	}
}

func (s *SessionStore) Get(id string) string {
	return s.store[id]
}
