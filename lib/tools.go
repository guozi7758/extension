package lib

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"time"
)

/*
 * 描述: 生成MD5哈希值
 *
 ***************************************************************************/
func StrMd5Str(strPass string) string {
	w := md5.New()
	io.WriteString(w, strPass)
	return fmt.Sprintf("%x", w.Sum(nil))
}

func StrMd5STR(strPass string) string {
	w := md5.New()
	io.WriteString(w, strPass)
	return fmt.Sprintf("%X", w.Sum(nil))
}

/*
 * 描述: 手机号识别
 *
 ***************************************************************************/
func IsPhone(strPhone string) bool {
	regular := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(19[0-9])|(17[0-9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(strPhone)
}

/*
 * 描述: 身份证识别
 *
 ***************************************************************************/
func IsNumberId( number_id string) bool {
	regular := "^((1[1-5])|(2[1-3])|(3[1-7])|(4[1-6])|(5[0-4])|(6[1-5])|71|(8[12])|91)\\d{4}((19\\d{2}(0[13-9]|1[012])(0[1-9]|[12]\\d|30))|(19\\d{2}(0[13578]|1[02])31)|(19\\d{2}02(0[1-9]|1\\d|2[0-8]))|(19([13579][26]|[2468][048]|0[48])0229))\\d{3}(\\d|X|x)?$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString( number_id )
}

func RandNum( nMin, nMax int )int{

	for i := 0; i < 20 ; i++ {
		rand.Seed(time.Now().Unix())
		rnd := rand.Intn( nMax )
		if rnd > nMin {
			return rnd
		}
	}
	return int( nMin + ( ( nMax - nMin ) / 2 ) )

}
