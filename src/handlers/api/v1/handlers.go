package v1

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
