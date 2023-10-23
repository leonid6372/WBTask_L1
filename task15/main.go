// Task 15

// ...

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	justString := someFunc()
}
