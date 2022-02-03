package main

import (
	"design-patterns/creational/singleton/api"
	"design-patterns/creational/singleton/sessionmanager"
	"testing"
)

func TestSessionUidIsSameForSameDataStoredInDiffWay(t *testing.T) {
	firstSession := sessionmanager.GetSessionStore()
	firstSession.Put(USER_EMAIL)
	firstId := firstSession.Get(USER_EMAIL)

	secondSession := sessionmanager.GetSessionStore()
	secondSession.Put(USER_EMAIL)
	secondId := secondSession.Get(USER_EMAIL)

	thirdId := mockSingleton()
	fourthId := api.TestApi(USER_EMAIL)

	if (firstId != secondId) || (secondId != thirdId) || (thirdId != fourthId) {
		t.Error("Singleton failed")
	}

}
