package main

import (
	"design-patterns/creational/singleton/api"
	"design-patterns/creational/singleton/sessionmanager"
	"fmt"
)

const USER_EMAIL = "ram@test.com"

func main() {
	fmt.Println("Singleton Design Pattern")
	firstSession := sessionmanager.GetSessionStore()
	firstSession.Put(USER_EMAIL)
	firstId := firstSession.Get(USER_EMAIL)

	secondSession := sessionmanager.GetSessionStore()
	secondSession.Put(USER_EMAIL)
	secondId := secondSession.Get(USER_EMAIL)

	thirdId := testSingleton()
	fourthId := api.TestApi(USER_EMAIL)

	fmt.Println("First ID:", firstId)
	fmt.Println("Second ID:", secondId)
	fmt.Println("Third ID:", thirdId)
	fmt.Println("Fourth ID:", fourthId)

	if (firstId == secondId) && (secondId == thirdId) && (thirdId == fourthId) {
		fmt.Println("Singleton Test Success")
	} else {
		fmt.Println("Singleton Test Failed")
	}
}

func testSingleton() string {
	mySession := sessionmanager.GetSessionStore()
	mySession.Put(USER_EMAIL)
	id := mySession.Get(USER_EMAIL)

	return id
}
