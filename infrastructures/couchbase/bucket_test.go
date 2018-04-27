package couchbase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucket_Constructor(t *testing.T) {
	t.Run("with correct config should be able to connect to db", func(t *testing.T) {
		configCB := NewCouchbaseConfig("couchbase://127.0.0.1", "Administrator", "password")
		cb := NewCouchbase(configCB)

		configBucket := NewBucketConfig("test", "")
		bucket, err := NewBucket(*cb, configBucket)

		assert.Nil(t, err, "should be nil")
		assert.NotNil(t, bucket, "should be not nil")
	})

	t.Run("with wrong config should be not able to connect to db", func(t *testing.T) {
		configCB := NewCouchbaseConfig("couchbase://127.0.0.1", "Administrator", "password123")
		cb := NewCouchbase(configCB)

		configBucket := NewBucketConfig("test", "")
		_, err := NewBucket(*cb, configBucket)

		assert.NotNil(t, err, "should be not nil")
	})
}

func TestBucket_Execute(t *testing.T) {
	t.Run("get executor", func(t *testing.T) {
		configCB := NewCouchbaseConfig("couchbase://127.0.0.1", "Administrator", "password")
		cb := NewCouchbase(configCB)

		configBucket := NewBucketConfig("test", "")
		bucket, _ := NewBucket(*cb, configBucket)

		assert.NotNil(t, bucket.Execute(), "should be not nil")
	})
}
