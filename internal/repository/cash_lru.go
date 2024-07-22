package repository

import (
	lru "github.com/hashicorp/golang-lru/v2"
	//"github.com/sirupsen/logrus"
)

func InitCacheLRU() *lru.Cache[int, float64] {
	cache, _ := lru.NewWithEvict(2,
		func(key int, value float64) {
			//logrus.Infof("Evicted: key=%v value=%v\n", key, value)
		},
	)
	return cache
}
