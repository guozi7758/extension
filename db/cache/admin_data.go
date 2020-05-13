package cache
/*
 * 描述: 工作台数据
 *
 *******************************************************************/
type AdminData struct {
	TFirst FirstInfo `json:"first"`		//首咨数据
}
func (this *AdminData)Get()error{
	return this.TFirst.Get()
}

