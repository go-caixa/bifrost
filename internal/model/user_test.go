package model_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/go-caixa/bifrost/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestUserAssginedID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assertion := assert.New(t)

		newUser := model.User{
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		newUser.AssginedID()
		assertion.NotEqual("", newUser.ID)
	})
}
