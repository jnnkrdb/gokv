package cachestorage

import "fmt"

type CacheStorage map[string]map[string]string

func CreateCacheStorage() CacheStorage {
	return make(CacheStorage)
}

// -------------------------------------------------------------- required functions

func (cs CacheStorage) ListBuckets() ([]string, error) {
	return []string{}, fmt.Errorf("CacheStorage.ListBuckets not implemented")
}

func (cs CacheStorage) ListKeys(bucket string) ([]string, error) {
	return []string{}, fmt.Errorf("CacheStorage.ListKeys not implemented")
}

func (cs CacheStorage) GetKey(bucket, key string) (string, error) {
	return "", fmt.Errorf("CacheStorage.GetValue not implemented")
}

func (cs CacheStorage) Write(bucket, key, value string) error {
	return fmt.Errorf("CacheStorage.WriteValue not implemented")
}

func (cs CacheStorage) Delete(bucket, key string) error {
	return fmt.Errorf("CacheStorage.DeleteKey not implemented")
}
