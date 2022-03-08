# Simple cache

Very simple in memory cache with ttl

Installation:
```sh
go get -u github.com/mtr888/simplecache
```

Usage:
```go
cache := simplecache.NewCache()

cache.Set("userId", 42, 5*time.Second)
userId, err := cache.Get("userId")
cache.Delete("userId")
```