// You can edit this code!
// Click here and start typing.
package main

func main() {
	num := 20
	println("Address of num:", &num)
	changeNumberByValue(num)
	println("value of num:", num)

	changeNumberByRef(&num)
	println("value of num:", num)

}
func changeNumberByRef(num *int) {
	println("changeNumberByRef Address of num:", num)
	*num = *num + 2
}

func changeNumberByValue(num int) {
	println("changeNumberByValue Address of num:", &num)
	num = num + 2
}
