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
}
