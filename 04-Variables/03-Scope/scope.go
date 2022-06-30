package main

var globalBlock = 125 // globalBlock is a global variable

func main() {

	println(globalBlock)
	{
		var localBlock = 124 // localBlock is a local variable
		println(localBlock)

		{
			var localBlock = 1250 // this is a local variable
			localBlock = localBlock * 9
			println(localBlock)
		}
		println(localBlock)
	}
	println(globalBlock)
	//println(localBlock)

	{
		// var block2 = 123
	}
}

func calc() {
	globalBlock = globalBlock + 10
}
