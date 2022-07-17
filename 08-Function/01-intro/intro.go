package main

func main() {
	Print1()
	str := Print2()
	println(str)
	str3 := Print3("33")
	println(str3)
}

// without input and output
func Print1() {
	println("Hello World1")
}

// without input and with output
func Print2() string {
	return "Hello World2"
}

// with input and output
func Print3(str string) string {
	return "Hello World " + str
}
