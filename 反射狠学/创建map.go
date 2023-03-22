var userMap map[int]*common.User
mapType := reflect.TypeOf(userMap)
// mapValue:=reflect.MakeMap(mapType)
mapValue := reflect.MakeMapWithSize(mapType, 10)

user := &common.User{
	Id:     7,
	Name:   "杰克逊",
	Weight: 65.5,
	Height: 1.68,
}
key := reflect.ValueOf(user.Id)
mapValue.SetMapIndex(key, reflect.ValueOf(user))                    //SetMapIndex 往map里添加一个key-value对
mapValue.MapIndex(key).Elem().FieldByName("Name").SetString("令狐一刀") //MapIndex 根据Key取出对应的map
userMap = mapValue.Interface().(map[int]*common.User)
fmt.Printf("user name %s %s\n", userMap[7].Name, user.Name)
