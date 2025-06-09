package storage

import (
	"github.com/jnnkrdb/gokv/conf"
)

var STORAGE Storage

func init() {
	if conf.NC.Storage.Type == "cache" {
		STORAGE = make(CacheStorage)
	}
}

type Storage interface {
	GetValue(string, string) string
	WriteValue(string, string, string)
}

type CacheStorage map[string]map[string]string

func (cs CacheStorage) GetValue(bucket, key string) string
func (cs CacheStorage) WriteValue(bucket, key, value string)
