package helper

type Helper interface {
	NormalizeStr(str string) string
	HasEqualString(str1 string, str2 string, ignoreCase ...bool) bool
	IsNullOrEmpty(str interface{}) bool
}
