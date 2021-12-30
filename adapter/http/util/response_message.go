package util

type message struct {
	Message string `json:"message"`
}

func ResponseMessage(err error) message {
	return message{
		Message: err.Error(),
	}
}
