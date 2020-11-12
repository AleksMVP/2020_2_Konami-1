package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"konami_backend/internal/pkg/csrf"
	"testing"
	"time"
)

func TestCSRF(t *testing.T) {
	size := 24

	rb := make([]byte, size)
	_, err := rand.Read(rb)
	assert.NoError(t, err)
	expireTime := 3600
	sID := "dase13r23f"
	testError := errors.New("testError")

	rs := base64.URLEncoding.EncodeToString(rb)

	t.Run("TestOnUsedToken", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf.NewMockRepository(ctrl)

		cryptHashToken, err := NewCsrfUseCase(rs, int64(expireTime), csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.Create(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Validate(token).
			Return(nil)

		csrfRepo.EXPECT().
			Add(token, int64(expireTime)).
			Return(nil)

		ok, err := cryptHashToken.Check(sID, token)
		assert.NoError(t, err)
		assert.Equal(t, ok, true)
	})

	t.Run("TestOnUsedToken", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf.NewMockRepository(ctrl)

		cryptHashToken, err := NewCsrfUseCase(rs, int64(expireTime), csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.Create(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Validate(token).
			Return(testError)

		ok, err := cryptHashToken.Check(sID, token)
		assert.NoError(t, err)
		assert.Equal(t, ok, false)
	})

	t.Run("TestOnUsedToken", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		csrfRepo := csrf.NewMockRepository(ctrl)

		cryptHashToken, err := NewCsrfUseCase(rs, int64(expireTime), csrfRepo)
		assert.NoError(t, err)

		token, err := cryptHashToken.Create(sID, time.Now().Unix())
		assert.NoError(t, err)

		csrfRepo.EXPECT().
			Validate(token).
			Return(nil)

		csrfRepo.EXPECT().
			Add(token, int64(expireTime)).
			Return(testError)

		ok, err := cryptHashToken.Check(sID, token)
		assert.Error(t, err)
		assert.Equal(t, ok, false)
	})
}
