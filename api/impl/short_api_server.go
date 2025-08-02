package impl

import (
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)
type GoShortApiServer struct {
	*ShortApiImpl
}

func NewGoShortApiServer(db *gorm.DB , redisClient *redis.Client) *GoShortApiServer {
	return &GoShortApiServer{
		ShortApiImpl: &ShortApiImpl{
			db: db,
			redis: redisClient,
		},
	}
}
