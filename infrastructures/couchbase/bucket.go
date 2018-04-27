package couchbase

import (
	"github.com/couchbase/gocb"
)

type Bucket struct {
	bucket *gocb.Bucket
}

func NewBucket(couchbase Couchbase, config BucketConfig) (*Bucket, error) {
	bucket, err := couchbase.cluster.OpenBucket(config.bktName, config.bktPassword)
	return &Bucket{bucket}, err
}

func (b *Bucket) Execute() *gocb.Bucket {
	return b.bucket
}
