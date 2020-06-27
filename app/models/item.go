package models

type Items struct {
	Id	  int
	Name  string
	Price int
}
//
//// 静的に仮データを登録する
//var items []Items
//
//func NewItem() {
//	items = append(items,
//		Items{ Id: 1, Name: "aaa"},
//		Items{ Id: 2, Name: "bbb"},
//	)
//}
//
//func GetItem()  {
//	items = append(items,
//		Items{ Id: 1, Name: "aaa"},
//		Items{ Id: 2, Name: "bbb"},
//	)
//}