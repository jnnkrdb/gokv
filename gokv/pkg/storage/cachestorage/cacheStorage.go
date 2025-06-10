package cachestorage

import (
	"fmt"
	"log"
)

type CacheStorage map[string]map[string]string

func CreateCacheStorage() CacheStorage {
	return make(CacheStorage)
}

// -------------------------------------------------------------- required functions

func (cs CacheStorage) ListBuckets() ([]string, error) {
	var bucketList = []string{}
	for bucketKey := range cs {
		bucketList = append(bucketList, bucketKey)
	}
	return bucketList, nil
}

func (cs CacheStorage) DeleteBucket(bucket string) error {
	delete(cs, bucket)
	return nil
}

func (cs CacheStorage) ListKeys(bucket string) ([]string, error) {
	var keyList = []string{}
	for keyKey := range cs[bucket] {
		keyList = append(keyList, keyKey)
	}
	return keyList, nil
}

func (cs CacheStorage) GetKey(bucket, key string) (string, error) {
	b, ok := cs[bucket]
	if !ok {
		return "", fmt.Errorf("bucket %s does not exist or couldn't be found", bucket)
	}
	v, ok := b[key]
	if !ok {
		return "", fmt.Errorf("key [%s/%s] does not exist or couldn't be found", bucket, key)
	}
	return v, nil
}

func (cs CacheStorage) Write(bucket, key, value string) error {
	if _, ok := cs[bucket]; !ok {
		log.Printf("[WRN] couldn't indicate bucketMap -> creating new one\n")
		cs[bucket] = make(map[string]string)
	}
	if tmp, ok := cs[bucket]; !ok {
		log.Printf("[WRN] couldn't indicate bucketMap: %v", tmp)
		return fmt.Errorf("no bucketMap found")
	}
	cs[bucket][key] = value
	return nil
}

func (cs CacheStorage) DeleteKey(bucket, key string) error {
	delete(cs[bucket], key)
	return nil
}
