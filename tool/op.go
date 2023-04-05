package tool

func Swap[T any](left *T, right *T) {
	temp := *left
	*left = *right
	*right = temp
}

func Exchange[T any](on *T, with T) (result T) {
	result = *on
	*on = with
	return
}
