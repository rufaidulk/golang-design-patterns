package main

import (
	"design-patterns/creational/builder/notification"
	"design-patterns/creational/builder/notificationbuilder"
	"fmt"
)

func main() {
	fmt.Printf("----------------- Builder Pattern --------------\n\n")

	fmt.Println("----------------- Email Builder --------------")

	const email = "user@test.com"
	const emailBody = "Welcome email"
	director := notificationbuilder.NotificationDirector{}
	emailBuilder := &notification.EmailNotification{}
	director.SetBuilder(emailBuilder)

	if res, err := director.Build(email, emailBody); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("----------------- SMS Builder --------------")

	const phone = "00005555"
	const smsBody = "Welcome SMS"
	smsBuilder := &notification.SmsNotification{}
	director.SetBuilder(smsBuilder)

	if res, err := director.Build(phone, smsBody); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
