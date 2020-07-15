package handler

import (
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

//MyError .
type MyError struct {
	Code          pq.ErrorCode `json:"code"`
	Error         string       `json:"error"`
	Message       string       `json:"message"`
	InternalQuery error        `json:"internalQuery"`
}

//Message .
type Message struct {
	Message string `json:"message"`
}

//CheckErr .
func CheckErr(err error) (MyError, bool) {
	if pgerr, ok := err.(*pq.Error); ok {
		e := MyError{
			// InternalQuery: pgerr.InternalQuery,
			Code:    pgerr.Code,
			Error:   pgerr.Message,
			Message: "Algo aconteceu em nossos servidores, tente novamente por favor! Se o problema persistir entre em contato conosco.",
		}

		switch pgerr.Code {
		case "23503":
			e.Message = "Não foi possível excluir o registro, ele está referenciado em outro local. Verifique os outros cadastros."
			break
		}

		fmt.Println(e)
		return e, true
	}

	if err != nil {
		fmt.Println(err)
		e := MyError{
			InternalQuery: err,
			Code:          "171097",
			Error:         err.Error(),
			Message:       "Algo de estranho ocorreu agora!",
		}

		fmt.Println(e)
		return e, true
	}
	e := MyError{
		Message: "",
	}

	return e, false
}

//ReturnMessage .
func ReturnMessage(message string) []byte {
	payload, _ := json.Marshal(Message{
		Message: message,
	})
	return payload
}

//ReturnError .
func (e MyError) ReturnError() []byte {
	payload, _ := json.Marshal(e)
	return payload
}
