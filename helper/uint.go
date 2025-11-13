package helper

func UintPtr(u uint) *uint {
	return &u
}

func UintValue(u *uint) uint {
	if u == nil {
		return 0
	}
	return *u
}
