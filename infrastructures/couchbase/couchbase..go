package couchbase

import "github.com/couchbase/gocb"

type Couchbase struct {
	cluster *gocb.Cluster
}

func NewCouchbase(config CouchbaseConfig) *Couchbase {
	cluster, _ := gocb.Connect(config.clAddress)
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: config.clUsername,
		Password: config.clPassword,
	})
	return &Couchbase{cluster}
}
