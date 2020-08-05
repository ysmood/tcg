package tcg

// Throw error
func Throw(err error) {
	panic(err)
}

// Catch error
func Catch(handler func(error)) {
	val := recover()
	if err, ok := val.(error); ok {
		handler(err)
	}
}

// Guard converts the throwed error to error value
func Guard(fn func()) (err error) {
	defer Catch(func(e error) { err = e })

	fn()

	return
}
