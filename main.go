package main

import (
	"fmt"
	helper "go_junior/helper"
)

func main() {
	sh := helper.StringHelper{}
	validate := helper.GeneralHelper{}
	str1 := `L�aro Concrete Miami,)}"' café Á|đ|Đ|Â|à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ|è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễì|í|ị|ỉ|ĩ|ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ|ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ|ỳ|ý|ỵ|ỷ|ỹ`
	str2 := `Laro Concrete Miami,)}"' cafe A|d|D|A|a|a|a|a|a|a|a|a|a|a|a|a|a|a|a|a|a|e|e|e|e|e|e|e|e|e|e|ei|i|i|i|i|o|o|o|o|o|o|o|o|o|o|o|o|o|o|o|o|o|u|u|u|u|u|u|u|u|u|u|u|y|y|y|y|y`

	match := sh.HasEqualString(str1, str2)

	fmt.Println(match, validate.IsNullOrEmpty(nil))
}
