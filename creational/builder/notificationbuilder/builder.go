package notificationbuilder

type NotificationBuilder interface {
	SetConfig() NotificationBuilder
	SetPayload(string, interface{}) NotificationBuilder
	Send() (NotificationResponse, error)
}

type NotificationResponse struct {
	Status  bool
	Message string
	Data    interface{}
}

type NotificationDirector struct {
	builder NotificationBuilder
}

func (n *NotificationDirector) SetBuilder(builder NotificationBuilder) {
	n.builder = builder
}

func (n *NotificationDirector) Build(recipientId string, message interface{}) (NotificationResponse, error) {
	return n.builder.SetConfig().SetPayload(recipientId, message).Send()
}
