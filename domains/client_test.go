package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_EqualTo(t *testing.T) {
	client := Client{"AAAA-0000", "Dummy", true}

	t.Run("Given same client objects should be equals", func(t *testing.T) {
		newClient := NewClient("Dummy", "AAAA-0000", true)
		assert.Equal(t, true, newClient.EqualTo(client), "should be equals")
	})

	t.Run("Given different client name objects should be not equals", func(t *testing.T) {
		newClient := NewClient("DummyX", "AAAA-0000", true)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})

	t.Run("Given different client code objects should be not equals", func(t *testing.T) {
		newClient := NewClient("Dummy", "AAAA-0001", true)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})

	t.Run("Given different client isActive objects should be not equals", func(t *testing.T) {
		newClient := NewClient("Dummy", "AAAA-0000", false)
		assert.Equal(t, false, newClient.EqualTo(client), "should be not equals")
	})
}
