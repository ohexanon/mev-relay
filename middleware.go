package main

import (
	"time"

	"github.com/cornelk/hashmap"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func memoryRateLimiting(m *hashmap.Map[string, []time.Time]) context.Handler {
	return func(ctx iris.Context) {
		ip := ctx.Request().RemoteAddr
		reqs, ok := m.Get(ip)
		if !ok {
			go func(m1 *hashmap.Map[string, []time.Time], ip string) {
				m1.Set(ip, []time.Time{time.Now()})
			}(m, ip)
			ctx.Next()
			return
		}
		reqsLastMinute := 0
		oneMinuteAgo := time.Now().Add(-1 * time.Second)
		for i, req := range reqs {
			if req.After(oneMinuteAgo) {
				reqsLastMinute += 1
			} else {
				if len(reqs) == i+1 {
					reqs = reqs[:i]
				} else {
					reqs = append(reqs[:i], reqs[i+1:]...)
				}
			}
		}

		go func(m1 *hashmap.Map[string, []time.Time], ip string, r []time.Time) {
			r = append(r, time.Now())
			m.Set(ip, r)
		}(m, ip, reqs)

		if reqsLastMinute > 5 {
			ctx.StatusCode(iris.StatusTooManyRequests)
			ctx.WriteString("chill out")
			return
		}
		ctx.Next()
	}
}
