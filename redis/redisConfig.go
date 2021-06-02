package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sys_info/util"
)

//初始化 记录服务器ip到redis
func RedisInit() {

	c, err := redis.Dial("tcp", "39.98.239.44:7480")
	redis.DialDatabase(3)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok")
	}
	defer c.Close()

	// 密码鉴权
	_, err = c.Do("AUTH", "suone2one")
	if err != nil {
		fmt.Println("auth failed:", err)
	} else {
		fmt.Println("auth ok:")
	}
	ip := util.GetIp()
	c.Do("sadd","servers",ip)
	println("redis - servers记录ip: ",ip,"完成")
}

//日志存入redis
func SaveLogToRedis(info string,timeunix int64,ip string) {
	// 连接redis
	c, err := redis.Dial("tcp", "39.98.239.44:7480")
	redis.DialDatabase(3)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("Connect to redis ok")
	}
	defer c.Close()

	// 密码鉴权
	_, err = c.Do("AUTH", "suone2one")
	if err != nil {
		fmt.Println("auth failed:", err)
	} else {
		fmt.Println("auth ok:")
	}

	// 写入数据
	_, err = c.Do("ZADD", ip,timeunix,info)
	if err != nil {
		fmt.Println("redis set failed:", err)
	} else {
		fmt.Println("redis set ok")
	}
}