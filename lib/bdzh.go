package lib

import (
	"fmt"
	"regexp"
)
type BdZh struct {
	UserName string
	PassWord string
	ToKen	 string
}
/*
 *描述：百度账户and 密码 and token
 *Header头部信息
 ****************************************/

func NewName(name string) *BdZh {
	return &BdZh{
		UserName : name,
	}
}


func (p *BdZh) SetPassWord(PassWord string) {
	pattern :="\\D+"
	result,_ :=	regexp.MatchString(pattern,PassWord)
	if result == true {
		p.PassWord = PassWord
	} else {
		fmt.Println("密码非纯数字")

	}
}
func (p *BdZh) GetPassWord() string {
	return p.PassWord
}

func (p *BdZh) SetToken(ToKen string) {
	
	if len(ToKen) == 32  {
		p.ToKen = ToKen
	} else {
		fmt.Println("无效token")
	}
}

func (p *BdZh) GetToken() string {
	return p.ToKen
}
