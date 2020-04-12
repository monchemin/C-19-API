package service

import (
	mocks "c19/security/mocks/repository"
	"c19/security/model"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecurityService_CreateUser(t *testing.T) {
	t.Run("unable", func(*testing.T){
		r := &mocks.SecurityRepository{}
		ss := NewSecurityService(r)
		_, e := ss.CreateUser(context.Background(), model.UserCreateRequest{})
		assert.NotNil(t, e)
	})
}
