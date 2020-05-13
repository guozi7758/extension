package cache

import (
	"time"
)
import "strconv"
import "reflect"
import "../../lib"

/*
 * 描述: 首资日统计结构
 *
 *******************************************************************/
type FirstInfo struct {
		InCount         int64   `json:"in_count"`       //投递
        OutCount        int64   `json:"out_count"`      //领取
        At              int64   `json:"-"`              // 数据的时间点
}

func( this *FirstInfo )RKey()string{
        return "FirstInfo_H"
}

func ( this *FirstInfo )RDb()int{
        return 0
}

func ( this *FirstInfo )ToMap()map[string]interface{}{
        m := make(map[string]interface{})
        elem := reflect.ValueOf(this).Elem()
        relType := elem.Type()
        for i := 0; i < relType.NumField(); i++ {
                m[relType.Field(i).Name] = elem.Field(i).Interface()
        }
        return m
}

func ( this *FirstInfo ) Exists()( int64, error ){
        return Ur( this.RDb() ).Exists( this.RKey() ).Result()
}

func ( this *FirstInfo )HMSet()error{
	m := this.ToMap() 
	return Ur( this.RDb() ).HMSet( this.RKey(), m ).Err()
}

func ( this *FirstInfo )Get()error{
	fage, err := this.Exists()
	if err == nil {
		if fage == 1 {
			mss, err := Ur( this.RDb() ).HGetAll( this.RKey() ).Result()
			if err == nil {
				this.InCount,_  = strconv.ParseInt( mss["InCount"] , 10, 64 )
				this.OutCount,_ = strconv.ParseInt( mss["OutCount"] , 10, 64 )
				this.At,_       = strconv.ParseInt( mss["At"] , 10, 64 )
				if !lib.IsToday( this.At ){
						this.At = time.Now().Unix()
						this.InCount = 0
						this.OutCount =0
						return this.HMSet()
				}
			}else{
				return err
			}
		}else{
			this.At = time.Now().Unix()
			this.InCount = 0
			this.OutCount =0
			err = this.HMSet()
		}
	}
	return  err
}
func (this *FirstInfo)GetIn()error{
	fage,err := this.Exists()
	if err == nil {
		if fage ==1{
			mss,err := Ur(this.RDb()).HGetAll(this.RKey()).Result()
			if err == nil {
				this.InCount,_ =strconv.ParseInt(mss["InCount"],10,64)
			}
		}
	}
	return err
}
func ( this *FirstInfo )In()error{
	if err := this.Get(); nil != err {
		return err
	}
	return Ur( this.RDb() ).HIncrBy( this.RKey(), "InCount", 1 ).Err()
}

func ( this *FirstInfo )Out()error{
	if err := this.Get(); nil != err {
		return err
	}
	return Ur( this.RDb() ).HIncrBy( this.RKey(), "OutCount", 1 ).Err()
}







