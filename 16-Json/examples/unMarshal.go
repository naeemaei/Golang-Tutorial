package examples

import (
	"encoding/json"
	"fmt"
)

func UnmarshalExample1() {
	person1Json := []byte(`{"name":"Ali","family":"Rezaee","age":17}`)

	var person1 = Person{}
	err := json.Unmarshal(person1Json, &person1)

	if err != nil {
		panic(err)
	}

	println(person1.Name)
	println(person1.Family)
	println(person1.Age)
}

func UnmarshalExample2() {
	personsJson := []byte(`[{"name":"Ali","family":"Rezaee","age":17},{"name":"Milad","family":"Mohammadi","age":20},{"name":"Peyman","family":"Mohammadi"}]`)

	var persons = []Person{}
	err := json.Unmarshal(personsJson, &persons)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", persons)
}
