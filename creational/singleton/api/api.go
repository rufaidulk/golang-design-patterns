package api

import (
	"design-patterns/creational/singleton/sessionmanager"
)

func TestApi(email string) string {
	apiSession := sessionmanager.GetSessionStore()
	apiSession.Put(email)

	return apiSession.Get(email)
}
