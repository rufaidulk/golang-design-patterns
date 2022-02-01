package notification

import (
	"design-patterns/creational/builder/notificationbuilder"
	"errors"
)

type EmailNotification struct {
	config    map[string]string
	recipient string
	body      string
}

func (e *EmailNotification) SetConfig() notificationbuilder.NotificationBuilder {
	//api config initialization
	e.config = make(map[string]string)
	e.config["access_token"] = "API_KEY"
	return e
}

func (e *EmailNotification) SetPayload(recipient string, message interface{}) notificationbuilder.NotificationBuilder {
	//build the payload
	e.recipient = recipient
	e.body = message.(string)

	return e
}

func (e *EmailNotification) Send() (response notificationbuilder.NotificationResponse, err error) {
	//calling the api
	apiResponse := emailMockApi()
	if apiResponse["status"] == "SUCCESS" {
		response.Status = true
		response.Message = apiResponse["message"]
	} else {
		response.Status = false
		response.Message = apiResponse["message"]
		err = errors.New(apiResponse["message"])
	}

	response.Data = apiResponse

	return response, err
}

func emailMockApi() map[string]string {
	res := make(map[string]string)

	res["status"] = "SUCCESS"
	res["message"] = "Email Sent"
	res["trackId"] = "A1DG8UHJ"

	return res
}
