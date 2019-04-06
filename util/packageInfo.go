package util

//package util can be renamed to something else other than the dir name,
//for example if renamed to "test" the usuage must then also be test.GetPackageName
//BUT the import must be the dir name

//GetPackageName exported function GetPackageName should have comment or be unexported go-lint
func GetPackageName() string {
	return "util"
}
