package  main

import(
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool


func CreatePool(maxActive,maxIdle int,idleTimeout time.Duration,address string) {


      pool = &redis.Pool{
		   MaxIdle:maxIdle,
		   MaxActive:maxActive,
		   IdleTimeout:idleTimeout,
		   Dial:func()(redis.Conn ,error){
                 return redis.Dial("tcp",address)
		   },
	  }




}
