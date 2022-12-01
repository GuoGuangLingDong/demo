package utils

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

const CAPTCHA_FORMAT = "captcha:%s"

type imageCodeRedis struct {
}

var ImageCodeRedis = new(imageCodeRedis)

//实现设置captcha的方法
func (r imageCodeRedis) Set(id, value string) (err error) {
	key := fmt.Sprintf(CAPTCHA_FORMAT, id)
	g.Redis().Do(nil, "SET", key, value, "ex", 600)
	return err
}

//实现获取captcha的方法
func (r imageCodeRedis) Get(id string, clear bool) (ret string, err error) {
	key := fmt.Sprintf(CAPTCHA_FORMAT, id)
	// val, err := RedisDb.Get(ctx, key).Result()
	gv, err := g.Redis().Do(nil, "GET", key)
	if err != nil {
		return
	}
	if gv != nil && !gv.IsEmpty() {
		err = gv.Scan(&ret)
		if err != nil {
			return
		}
	} else {
		return
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		// err := RedisDb.Del(ctx, key).Err()
		g.Redis().Do(nil, "DEL", key)
		if err != nil {
			return
		}
	}

	return
}

//实现验证captcha的方法
// func (r RedisStore) Verify(id, answer string, clear bool) bool {
// 	v := RedisStore{}.Get(id, clear)
// 	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
// 	return v == answer
// }
