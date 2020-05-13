package modes

import (
	"fmt"
	"time"
)

/*
 * 描述:推广导入excel数据表
 * 表名: file_data
 *
 ********************************************************************/
type FileData struct {
	Id       int64  `json:"impo_id" xorm:"id"`        // 文件id
	UId      int64  `json:"impo_uid" xorm:"uid"`      // 员工id
	ProId    int64  `json:"impo_pid" xorm:"miso_id"`  // 咨询项目id
	Type     int    `json:"impo_state" xorm:"state"`  // 文件是否解析 0未解析 1解析
	KeyFile  string `json:"impo_url" xorm:"impo_url"` // 文件url
	CreateAt int64  `json:"impo_at" xorm:"at"`        // 创建时间
}

func (this *FileData) TableName() string {
	return "file_data"
}

func (this *FileData) Save() (int64, error) {
	this.CreateAt = time.Now().Unix()
	return Db(0).Insert(this)
}

func (this *FileData) Get() (bool, error) {
	return Db(0).Get(this)
}

func (this *FileData) update(where string, field string) (int64, error) {
	return Db(0).Where(where).Cols(field).Update(this)
}

func (this *FileData) IdSet(field string) (int64, error) {
	where := fmt.Sprintf("id = %d", this.Id)
	return this.update(where, field)
}

func (this *FileData) where(fage, count, page int, where, field string) ([]ProId, error) {
	list := make([]ProId, 0)
	var err error
	if field == "" {
		field = "id"
	}
	page--
	if 0 == fage { // 从小到大排序
		err = Db(0).Where(where).
			Asc(field).
			Limit(count, page*count).
			Find(&list)
	} else { // 从大到小排序
		err = Db(0).Where(where).
			Desc(field).
			Limit(count, page*count).
			Find(&list)
	}
	return list, err
}
