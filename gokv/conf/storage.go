package conf

import (
	"log"

	"github.com/jnnkrdb/gokv/pkg/storage"
	"github.com/jnnkrdb/gokv/pkg/storage/cachestorage"
)

// the actual reference for the storage backend
var STORAGE storage.Storage

// initialize the storage and load it into the cache
func LoadStorage() {
	log.Printf("[INF] loading storage backend: %s\n", NC.Storage.Type)
	switch NC.Storage.Type {
	case "cache":
		STORAGE = cachestorage.CreateCacheStorage()
	default:
		log.Fatalf("the configuration for the storage is not valid. you have to decide which implementation to use.")
	}
}
