package main

import (
	"context"
	"fmt"
	"time"
)

// 需求: 我是一个内容网站，通过getContent方法，我获取内容推荐给用户，
// 但是如果推荐服务超时了（不管因为什么），我就返回给用户最热的10条新闻，而不是直接返回504

// 这个方法的目的是，控制子调用的超时，因为整个getContent是3秒，而获取recommondContent是2秒，
// 这2秒是包含在这3秒钟的，要做到的就是这个效果
func shrinkTimeoutContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	t, ok := ctx.Deadline()
	if !ok {
		return context.WithTimeout(ctx, timeout)
	}

	nt := time.Now()
	if t.Before(nt) {
		return context.WithTimeout(ctx, 0)
	}
	return context.WithDeadline(ctx, nt.Add(timeout))
}

func GetRecommendContent(ctx context.Context) string {
	time.Sleep(time.Second * 4)
	return "recommend content"
}

func GetContent(ctx context.Context) string {
	rcCtx, cancel := shrinkTimeoutContext(ctx, time.Second * 2)

	var res string
	go func() {
		res = GetRecommendContent(rcCtx)
		cancel()
	}()

	select {
	case <- rcCtx.Done():
		if rcCtx.Err() == context.DeadlineExceeded {
			return "ten hot news"
		}
	}
	return res
}

func main() {
	rCtx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	fmt.Println(GetContent(rCtx))
}
