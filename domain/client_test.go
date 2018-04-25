package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_EqualTo(t *testing.T) {
	client := Client{1, "Dummy", "AAAA-0000", true}

	t.Run("Given same client objects should be equals", func(t *testing.T) {
		newClient := NewClient(1, "Dummy", "AAAA-0000", true)
		assert.Equal(t, true, newClient.EqualTo(client), "should be equals")
	})

	t.Run("Given different client id objects should be not equals", func(t *testing.T) {
		newClient := NewClient(2, "Dummy", "AAAA-0000", true)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})

	t.Run("Given different client name objects should be not equals", func(t *testing.T) {
		newClient := NewClient(1, "DummyX", "AAAA-0000", true)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})

	t.Run("Given different client code objects should be not equals", func(t *testing.T) {
		newClient := NewClient(1, "Dummy", "AAAA-0001", true)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})

	t.Run("Given different client isActive objects should be not equals", func(t *testing.T) {
		newClient := NewClient(1, "Dummy", "AAAA-0000", false)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})
}
