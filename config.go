package common

import (
	"context"
	"fmt"
	"log"

	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	ConsulStr = "http://124.222.29.179:8500"
	UserFileKey = "mysql-product"
	RedisFileKey = "redis"
)

func GetConsulConfig(url string, fileKey string) (*viper.Viper, error) {
	conf := viper.New()
	conf.AddRemoteProvider("consul", url, fileKey)
	conf.SetConfigType("json")
	err := conf.ReadRemoteConfig()
	if err != nil {
		log.Println("viper conf err: ", err)
	}
	return conf, nil
}

func GetMysqlFromConsul(vip *viper.Viper) (db *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	str := vip.GetString("user") + ":" + vip.GetString("pwd") + "@tcp(" + vip.GetString("host") + ":" + vip.GetString("port") + ")/" + vip.GetString("database") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(str)
	db, err = gorm.Open(mysql.Open(str), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Println("db err :", err)
	}
	log.Println(db)
	return db, nil
}

func GetRedisFromConsul(vip *viper.Viper) (*redis.Client, error) {
	red := redis.NewClient(
		&redis.Options{
			Addr: vip.GetString("addr"),
			Password: vip.GetString("password"),
			DB: vip.GetInt("DB"),
			PoolSize: vip.GetInt("poolSize"),
			MinIdleConns: vip.GetInt("minIdleConn"),
	})
	fmt.Println("redis: ", red)
	return red, nil
}

func SetUserToken(red *redis.Client, key string, val []byte, timeTTL time.Duration) {
	red.Set(context.Background(), key, val, timeTTL)
}

func GetUserToken(red *redis.Client, key string) string {
	res, err := red.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("GetUserToken: ", err)
	}
	return res
}