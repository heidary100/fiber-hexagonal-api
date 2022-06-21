package domain_test

import (
	"testing"

	"github.com/heidary100/fiber-hexagonal-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

// ··· USER TESTS ··· //

func TestNewUser(t *testing.T) {
	user := domain.NewUser("1001", "hello world", "helloworld")

	assert.Equal(t, "1001", user.ID)
	assert.Equal(t, "hello world", user.Name)
	assert.Equal(t, "helloworld", user.Username)
}
