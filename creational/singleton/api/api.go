package api

import (
	"design-patterns/creational/singleton/sessionmanager"
	"fmt"
)

func TestApi() {
	apiSession := sessionmanager.GetSessionStore()

	apiSession.Put("ram@test.com")
	fmt.Println(apiSession.Get("ram@test.com"))
}
