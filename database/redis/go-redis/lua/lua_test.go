package lua__test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//noinspection GoInvalidCompositeLiteral
func TestString(t *testing.T) {
	ctx := context.Background()
	Client.FlushAll(ctx)

	Client.Set(ctx, "foo", "bar", 0)

	var luaScript = redis.NewScript(`return redis.call("GET" , KEYS[1])`)

	n, err := luaScript.Run(ctx, Client, []string{"foo"}).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(n, err)
}

func TestZSet(t *testing.T) {
	ctx := context.Background()
	Client.FlushAll(ctx)

	foo := []*redis.Z{
		{
			1732, "George Washington",
		},
		{
			1809, "Abraham Lincoln",
		},
		{
			1858, "Theodore Roosevelt",
		},
	}

	Client.ZAdd(ctx, "presidents", foo...)

	var luaScript = redis.NewScript(`
        local elements = redis.call("ZRANGE" , KEYS[1] , 0 , 0) 
        redis.call("ZREM" , KEYS[1] , elements[1])
        return elements[1]
    `)

	n, err := luaScript.Run(ctx, Client, []string{"presidents"}, 1).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(n, err)
}

func TestEvalSha(t *testing.T) {
	ctx := context.Background()
	Client.FlushAll(ctx)

	Client.Set(ctx, "foo", "bar", -1)

	var luaScript = `return redis.call("INFO")`
	//返回的脚本会产生一个 sha1 哈希值,下次用的时候可以直接使用这个值
	result, err := Client.ScriptLoad(ctx, luaScript).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // bb2a35aa586628e1875ba50005a113e06b463cc4

	foo := Client.EvalSha(ctx, result, []string{})
	fmt.Println(foo.Val())
}

func TestSubStock(t *testing.T) {
	ctx := context.Background()
	Client.FlushAll(ctx)

	//设置初始值
	Client.Set(ctx, "stock", "10", 0)
	//编写脚本 - 检查数值，是否够用，够用再减，否则返回减掉后的结果
	var luaScript = redis.NewScript(`
		local value = redis.call("Get", KEYS[1])
		print("当前值为 " .. value);
		if( value - KEYS[2] >= 0 ) then
			local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
   			print("剩余值为" .. leftStock );
			return leftStock
		else
			print("数量不够，无法扣减");
			return value - KEYS[2]
		end
		return -1
	`)

	//执行脚本
	n, err := luaScript.Run(ctx, Client, []string{"stock", "6"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}

func TestEvalSha1(t *testing.T) {
	var script = `
		local value = redis.call("Get", KEYS[1])
		print("当前值为 " .. value);
		if( value - KEYS[2] >= 0 ) then
			local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
   			print("剩余值为" .. leftStock );
			return leftStock
		else
			print("数量不够，无法扣减");
			return value - KEYS[2]
		end
		return -1
	`

	ctx := context.Background()
	luaHash, _ := Client.ScriptLoad(ctx, script).Result() //返回的脚本会产生一个sha1哈希值,下次用的时候可以直接使用这个值

	n, err := Client.EvalSha(ctx, luaHash, []string{"stock", "2"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", n, err)
}
