package core

import "errors"

var (
	ErrFailedToSyncCacheWithAPI = errors.New("failedToSyncCacheWithAPI")
	ErrFailedToSyncCache        = errors.New("failedToSyncCache")
)
