package main

type BigData struct {
	// 1 gb of memory
	// ..
	// ..
	// ..
}

func doSomethingWithData(data *BigData) {
	// manipulate the data inside this function
}

func main() {
	data := &BigData{} // 1 gb

	// something like this would be bad -> not perfomant...
	// which should reinforce a decision for us to use pointers
	for i := 0; i < 10000; i++ {
		// copy 1 gb of data inside if that function
		// -> by using pointers instead, we will only be passing 8 bytes of pointer
		doSomethingWithData(data)
	}
}
