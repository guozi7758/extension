package cache

import (
	//"strconv"
	"reflect"
	"strconv"

	//"../../lib"
)

/*
 * 描述: 神马统计数据
 *
 *******************************************************************/
type SMInfo struct {
	Cost         int64   	`json:"in_cost"`       	//消费
	Click        int64   	`json:"click"`      	//点击
}
func(this *SMInfo)RKey()string{
	return "SMInfo_H"
}
func(this *SMInfo)RDb()int{
	return 0
}
func(this *SMInfo)ToMap()map[string]interface{}{
	m := make(map[string]interface{})
	elem := reflect.ValueOf(this).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	return m
}
func ( this *SMInfo ) Exists()( int64, error ){
	return Ur( this.RDb() ).Exists( this.RKey() ).Result()
}
func ( this *SMInfo )HMSet()error{
	m := this.ToMap()
	return Ur( this.RDb() ).HMSet( this.RKey(), m ).Err()
}

func ( this *SMInfo )Get()error{
	fage,err := this.Exists()
	if err == nil{
		if fage == 1 {
			mss, err := Ur( this.RDb() ).HGetAll( this.RKey() ).Result()
			if err == nil{
				this.Cost,_  = strconv.ParseInt( mss["Cost"] , 10, 64 )
				this.Click,_ = strconv.ParseInt( mss["Click"] , 10, 64 )
			}else{
				return err
			}

		}else{
			err = this.HMSet()
		}
	}

	return  err
}