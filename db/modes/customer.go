package modes

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"strconv"
	"strings"
	"time"
)

/*
*描述：导入客户的信息数据
*
*/

type Customer struct {
	Id 				int64		`json:"cust_id" xorm:"id"`					//id
	PId				int64		`json:"cust_pid" xorm:"pid"`				//项目id
	Num 			int64		`json:"cust_num" xorm:"num"`				//序列号
	CustName 		string		`json:"cust_custname" xorm:"custname"` 		//客户姓名
	CustPhone 		string		`json:"cust_custphone" xorm:"custphone"`	//客户电话
	ChannelId		string 		`json:"cust_channelid" xorm:"channelid"`	//客户来源渠道
	Date			string 		`json:"cust_date" xorm:"date"`				//数据时间
	At 				int64 		`json:"cust_at" xorm:"at"`					//插入时间
}
func (this *Customer) TableName() string {
	return "customer"
}

func (this *Customer) Save() (int64, error) {
	this.At = time.Now().Unix()
	return Db(1).Insert(this)
}

func (this *Customer) Get() (bool, error) {
	return Db(1).Get(this)
}

func (this *Customer) update(where string, field string) (int64, error) {
	return Db(1).Where(where).Cols(field).Update(this)
}

func (this *Customer) IdSet(field string) (int64, error) {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where, field)
}

func (this *Customer) where(fage, count, page int, where, field string) ([]Customer, error) {
	list := make([]Customer, 0)
	var err error
	if field == "" {
		field = "id"
	}
	page--
	if 0 == fage { // 从小到大排序
		err = Db(1).Where(where).
			Asc(field).
			Limit(count, page*count).
			Find(&list)
	} else { // 从大到小排序
		err = Db(1).Where(where).
			Desc(field).
			Limit(count, page*count).
			Find(&list)
	}
	return list, err
}

func (this *Customer)ParsingFile(imp_file string)error {

	// 首先读excel
	xlsx, err := excelize.OpenFile(imp_file)
	if err != nil {
		return nil
	}
	rows,err:=xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex()))
	//c_id :=0

	for i, row := range rows {

		if i == 0 {
			continue
		}
		k := strings.Contains(row[0], ".")

		if 	k != true {
			var lib Customer
			lib.PId = this.PId
			lib.Num ,_ = strconv.ParseInt(row[1], 10, 64)
			lib.CustName = row[2]
			lib.CustPhone = row[3]
			lib.ChannelId = row[4]
			lib.Date = row[5]
			lib.Id = 0
			//lib.Save()
			if _, err := lib.Save();err !=nil {
				fmt.Println(err.Error())

			}
			//c_id = int(lib.Id)

		}
	}
	return err
}