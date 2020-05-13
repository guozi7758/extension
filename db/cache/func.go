package cache

import (
	"../../db"
	"github.com/go-redis/redis"
)


func Ur( f int )*redis.Client{
	return db.UseRedis( f )
}

func init(){

	//go func(){
	//	var ch chan int
	//	ticker := time.NewTicker(time.Second * 1200)
	//	go func() {
	//		for range ticker.C {
	//			fmt.Println( time.Now().Format("2006-01-02 15:04:05"))
	//			var sminfo SMInfo
	//			tack_id := lib.SmZero.GetTaskId()
	//			//time.Sleep(time.Second *5)
	//			file_data := lib.SmZero.GetFileId( tack_id )
	//			smclick,smcost := lib.SmZero.GetFileCilck( file_data )
	//			fmt.Println(smclick,smcost)
	//			sminfo.Click = smclick
	//			sminfo.Cost = smcost
	//			if  err := sminfo.HMSet(); nil != err{
	//				fmt.Println("采集失败:", err)
	//			}
	//		}
	//		ch <- 1
	//	}()
	//	<-ch
	//}()

}
