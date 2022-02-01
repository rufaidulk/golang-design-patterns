package notification

import (
	"design-patterns/creational/builder/notificationbuilder"
	"errors"
)

type SmsNotification struct {
	config    map[string]string
	recipient string
	body      string
}

func (s *SmsNotification) SetConfig() notificationbuilder.NotificationBuilder {
	//api config initialization
	s.config = make(map[string]string)
	s.config["access_token"] = "API_KEY"
	return s
}

func (s *SmsNotification) SetPayload(recipient string, message interface{}) notificationbuilder.NotificationBuilder {
	//build the payload
	s.recipient = recipient
	s.body = message.(string)

	return s
}

func (s *SmsNotification) Send() (response notificationbuilder.NotificationResponse, err error) {
	//calling the api
	apiResponse := smsMockApi()
	msg := apiResponse["message"].(string)
	if apiResponse["status"] == "SUCCESS" {
		response.Status = true
		response.Message = msg
	} else {
		response.Status = false
		response.Message = msg
		err = errors.New(msg)
	}

	response.Data = apiResponse

	return response, err
}

func smsMockApi() map[string]interface{} {
	res := make(map[string]interface{})

	res["status"] = "SUCCESS"
	res["message"] = "SMS Sent"
	res["trackId"] = "A1DG8UHJ"
	res["balance"] = 225.50

	return res
}
