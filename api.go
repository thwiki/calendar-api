package main

import (
	"strings"

	client "github.com/bozd4g/go-http-client"
	"github.com/bradfitz/gomemcache/memcache"
)

var httpClient = client.New(Api)
var memcacheStore = memcache.New(Memcached)

type EventsParameter struct {
	Action        string `url:"action"`
	Format        string `url:"format"`
	Formatversion int    `url:"formatversion"`
	Query         string `url:"query"`
}

func CacheKey(sections ...string) string {
	return "calendar:" + strings.Join(sections, ":")
}

func GetEvents(start string, end string) (result []byte, err error) {
	cacheKey := CacheKey("events", start, end)

	if cached, err2 := memcacheStore.Get(cacheKey); err2 == nil {
		result = cached.Value
		return
	}

	defer (func() {
		memcacheStore.Set(&memcache.Item{Key: cacheKey, Value: result, Expiration: TTL})
	})()

	request, err := httpClient.GetWith("/api.php", EventsParameter{
		Action:        "ask",
		Format:        "json",
		Formatversion: 2,
		Query:         "[[事件开始::>" + start + "]][[事件开始::<" + end + "]]|?事件编号=code|?事件页面=name|?事件开始=startDate|?事件结束=endDate|?事件描述=desc|?事件图标=icon|sort=事件开始|order=asc",
	})

	if err != nil {
		return
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return
	}

	result = response.Get().Body

	if err != nil {
		return
	}

	return
}
