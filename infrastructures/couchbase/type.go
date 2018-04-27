package couchbase

type CouchbaseConfig struct {
	clAddress  string
	clUsername string
	clPassword string
}

type BucketConfig struct {
	bktName     string
	bktPassword string
}

func NewCouchbaseConfig(address string, username string, password string) CouchbaseConfig {
	return CouchbaseConfig{address, username, password}
}

func NewBucketConfig(bktName string, bktPassword string) BucketConfig {
	return BucketConfig{bktName, bktPassword}
}
