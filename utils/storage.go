package utils

import (
	"github.com/DevitoDbug/golangJWTAuthTemplate/models"
)

var Storage map[string]models.User

func init() {
	Storage = make(map[string]models.User)
}
