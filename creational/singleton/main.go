package main

import (
	"design-patterns/creational/singleton/api"
	"design-patterns/creational/singleton/sessionmanager"
	"fmt"
)

func main() {
	fmt.Println("Singleton design pattern")
	mySession := sessionmanager.GetSessionStore()
	mySession.Put("ram@test.com")
	ramsId := mySession.Get("ram@test.com")

	newSession := sessionmanager.GetSessionStore()
	newSession.Put("ram@test.com")
	ramsNewId := newSession.Get("ram@test.com")

	fmt.Println(ramsId)
	fmt.Println(ramsNewId)
	testSingleton()
	api.TestApi()
}

func testSingleton() {
	mySession := sessionmanager.GetSessionStore()
	mySession.Put("ram@test.com")
	ramsId := mySession.Get("ram@test.com")

	newSession := sessionmanager.GetSessionStore()
	newSession.Put("ram@test.com")
	ramsNewId := newSession.Get("ram@test.com")

	fmt.Println(ramsId)
	fmt.Println(ramsNewId)
}
