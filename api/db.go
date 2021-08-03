package api

import (
	bolt "go.etcd.io/bbolt"
)

func GetDB() (*bolt.DB, error) {
	return bolt.Open("/mnt/site-comment/comment.db", 0600, nil)
}
