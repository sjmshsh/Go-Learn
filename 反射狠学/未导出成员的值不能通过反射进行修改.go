	addrValue := valueUser.Elem().FieldByName("addr")
	if addrValue.CanSet() {
		addrValue.SetString("HH")
	} else {
		fmt.Println()
	}
