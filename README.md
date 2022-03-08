# simplecache

Very simple in memory cache

# import

go get github.com/mtr888/simplecache

# Usage

cache := simplecache.NewCache()

cache.Set("userId", 42)
userId := cache.Get("userId")
cache.Delete("userId")