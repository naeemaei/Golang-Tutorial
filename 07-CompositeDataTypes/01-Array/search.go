package main

func main() {
	names := [8]string{"Ali", "Reza", "Milad", "Mahtab", "Farshad", "Fereshte", "Hadi", "Mahdi"}

	searchKey := "Milad"

	for i, name := range names {
		if name == searchKey {
			println("Name found. Index: ", i)
			break
		}
	}

	for j := 0; j < len(names); j++ {
		if names[j] == searchKey {
			println("Name found. Index: ", j)
			break
		}
	}
}