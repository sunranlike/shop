package models

import (
	"context"
	"fmt"
	"time"
)

var ctx = context.Background()

const (
	CAPTCHA = "captcha:"
)

type RedisStore struct { //
}

// Set 这里具体实现接口：实现了三个方法：set，get，Verify 这三个方法组合起来实现了Store接口，
//就可以作为 base64Captcha.Store 的合法赋值。
//why so？接口即协议，你想要使用一些功能，你就要按照我的接口（协议）去实现这个功能，这样子才可以合法，
//我内部会拿你给我实现的方法去使用，也就是说你实现了的方法会被我被调用
//另外具体的redis实例会在rediscore中配置

func (receiver RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := RedisDb.Set(ctx, key, value, time.Minute*2).Err()
	return err
}
func (receiver RedisStore) Get(id string, clear bool) string {
	fmt.Println("调用redis Get,id:", id)
	key := CAPTCHA + id
	fmt.Println("check point 1")
	val, err := RedisDb.Get(ctx, key).Result()
	fmt.Println("check point 2")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := RedisDb.Del(ctx, key).Err()
		fmt.Println("check point 3")
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	//我们也要实现整个验证方法,逻辑比较简单:
	//就是先调用get方法获取验证码id,并且清除这个id
	//判断并且返回
	fmt.Println("verify:", "id:", id, "answer", answer)
	v := RedisStore{}.Get(id, clear)
	//判断然后直接返回,v就是对应id的真实验证码值,然后拿来和answer对比,对比是否
	return v == answer
}
