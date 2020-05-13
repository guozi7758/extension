package modes

import (
	"fmt"
	"time"
)

/*
 * 描述:运营日统计表
 * 表名: daily_report
 *
 ********************************************************************/
type DailyReport struct{
	Id		int64		`json:"dai_id" xorm:"id"`		//日统计报表ID
	Calls		int64		`json:"dai_calls" xorm:"calls"`		//拨打数量
	ORders		int64		`json:"dai_orders" xorm:"orders"`	//订单总数
	OAmount		int64		`json:"dai_o_amount" xorm:"o_amount"`	//订单金额
	Deal		int64		`json:"dai_deal" xorm:"deal"`		//成交单数
	Damount		int64		`json:"dai_d_amount" xorm:"d_amount"`	//成交金额
	Firstc		int64		`json:"dai_first_c" xorm:"first_c"`	//今日首资数量
	UseFirs		int64		`json:"dai_use_firs" xorm:"use_firs"`	//今日使用首资数量
	Unix		int64		`json:"dai_unix" xorm:"unix"`		//时间点，全是凌晨
	At		int64		`json:"dai_at" xorm:"at"`		//落库时间
}

func (this *DailyReport) TableName() string {
	return  "daily_report"
}

func ( this *DailyReport )Save()( int64, error ){
	this.At = time.Now().Unix()
	return Db(0).Insert(this)
}

func ( this *DailyReport)Get()( bool, error){
	return Db(0).Get(this)
}

func (this  *DailyReport)update( where string, field string )(int64, error){
	return Db(0).Where(where).Cols(field).Update(this)
}

/*
 * fage 2: 首资
 * 	1: 拨打量
 *	3: 订单
 *
 ****************************************************************************/
func ( this *DailyReport )In( fage int, unix int64, val ...int64 )error{
	this.Unix = unix - ( (unix + 28800) % 86400 )
	var field string
	var in_mode DailyReport
	fa, err := this.Get()
       if nil == err {
	       switch ( fage ){
			case 1:
				in_mode.Calls = val[0]
				field = "calls"
			case 2:
				in_mode.Firstc = val[0]
				in_mode.UseFirs = val[1]
				field = "first_c,use_firs"
			case 3:
				in_mode.ORders = val[0]
				in_mode.OAmount = val[1]
				in_mode.Deal   = val[2]
				in_mode.Damount = val[3]
				field = "orders,o_amount,deal,d_amount"
	       }
	       if fa {
		       in_mode.Id = this.Id
			_, err = in_mode.IdSet( field )
		}else{
			in_mode.Unix = this.Unix
			_, err = in_mode.Save()
		}
	}
	return err
}

func ( this *DailyReport)IdSet(field string)(int64, error)  {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where,field)
}





