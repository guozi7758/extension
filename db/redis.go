package db

import "fmt"
import "time"
import "encoding/json"
import "github.com/go-redis/redis"

type RDB struct {
	Home	string		`json:"home"`
	Pass	string		`json:"pass"`
	Db	int		`json:"db"`
}

var GReHand []*redis.Client

func UseRedis( f int )*redis.Client{
	return GReHand[f]
}

func ( this *RDB )Init(){
	hand := redis.NewClient(&redis.Options{
			Addr:     this.Home,
			Password: this.Pass, // no password set
			DB:       this.Db,        // use default DB
			ReadTimeout:  time.Millisecond * time.Duration(100),
			WriteTimeout: time.Millisecond * time.Duration(100),
			IdleTimeout:  time.Second * time.Duration(60),
		})
	if _, err := hand.Ping().Result(); err != nil{
		panic( err.Error() )
	}
	GReHand = append( GReHand, hand )
}

func init(){
	var val []RDB
	fmt.Println("init redis")
	data, err := ReadConfig("./config/redis.json")
	if err != nil {
		panic("读取文件错误:" + err.Error() )
	}
	json.Unmarshal( data, &val )
	fmt.Println( "redis", val )
	for _, v := range val{
		v.Init()
	}
}
