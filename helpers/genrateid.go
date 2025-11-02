package helpers

import (
	"github.com/google/uuid"
)


func GenerateContextId(sType string) string{
	
	ctxId := sType + uuid.NewString()

	return ctxId

}