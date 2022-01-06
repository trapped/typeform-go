package typeform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marotpam/typeform-go"
)

func TestRetrieveUserSelf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := typeform.NewUserService(newFakeServerClient(t))

		u, err := svc.RetrieveSelf()
		assert.Nil(t, err)

		assert.Equal(t, userIDRetrievedUserSelf, u.UserID)
	})
}
