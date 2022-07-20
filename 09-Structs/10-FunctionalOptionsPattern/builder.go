package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	Family    string
	Gender    string
	Height    int
	Weight    int
	HairColor string
}

type PersonBuilder struct {
	Person
}

func main() {

	// 1
	person := Person{Name: "Alireza", Age: 23}

	person.Family = "Rezaei"

	// 2 - builder pattern
	personBuilder := &PersonBuilder{}
	person2 := personBuilder.SetName("Maryam").SetFamily("Rezaei").SetAge(23).SetGender("Female").SetHairColor("Black").SetHeight(180).SetWeight(76).Build()

	fmt.Printf("Person1: %+v \n", person)
	fmt.Printf("Person2: %+v \n", person2)
}

func (builder *PersonBuilder) SetName(name string) *PersonBuilder {
	builder.Name = name
	return builder
}

func (builder *PersonBuilder) SetFamily(family string) *PersonBuilder {
	builder.Family = family
	return builder
}

func (builder *PersonBuilder) SetGender(gender string) *PersonBuilder {
	builder.Gender = gender
	return builder
}

func (builder *PersonBuilder) SetAge(age int) *PersonBuilder {
	builder.Age = age
	return builder
}

func (builder *PersonBuilder) SetHeight(height int) *PersonBuilder {
	builder.Height = height
	return builder
}

func (builder *PersonBuilder) SetWeight(weight int) *PersonBuilder {
	builder.Weight = weight
	return builder
}

func (builder *PersonBuilder) SetHairColor(hairColor string) *PersonBuilder {
	builder.HairColor = hairColor
	return builder
}

func (builder *PersonBuilder) Build() Person {
	return builder.Person
}

// Query builder
// if(name != nil)
// 	db.Persons
// .Where("Name = 'Ali'")
// .Where("Family = 'Rezaei'")
// .Select("Name, Family")
// Result => select name,family from persons where name = 'Ali' and family = 'Rezaei'