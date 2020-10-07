package cache

import "com.github/fabiosebastiano/go-rest-api/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
