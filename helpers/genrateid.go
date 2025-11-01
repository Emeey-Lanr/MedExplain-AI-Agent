package helpers

import (
	"github.com/google/uuid"
)


func GenerateContextId() string{
	
	ctxId := "ctx-" + uuid.NewString()

	return ctxId

}