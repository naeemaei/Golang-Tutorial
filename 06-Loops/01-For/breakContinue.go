package main

func main() {

	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		println(i)
	}

	for i := 1; i <= 100; i++ {
		if i == 50 {
			break
		}
		println(i)
	}
}
