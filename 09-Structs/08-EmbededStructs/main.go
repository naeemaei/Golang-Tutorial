package main

import "encoding/json"

type Product struct {
	Name   string
	Color  string
	Length int
	Width  int
	Weight int
	Price  int
	Brand  string
	MadeIn string
}

type ElectronicProduct struct {
	Product
	Ram                    int
	Cpu                    string
	ScreenSize             float32
	OperatingSystem        string
	OperatingSystemVersion string
}

type Mobile struct {
	ElectronicProduct
	SimcardCapacity int
	SimcardType     string
	NetworkType     string
	CameraCount     int
}

type Loptop struct {
	ElectronicProduct
	UsbPortCount int
	KeyboardType string
	HasCdRom     bool
}

func main() {

	mobile := Mobile{}
	mobile.Name = "Samsung M4150"
	mobile.Brand = "Samsung"
	mobile.Color = "Blue"
	mobile.Ram = 4
	mobile.OperatingSystem = "Android"
	mobile.OperatingSystemVersion = "11"

	mobile.CameraCount = 2
	mobile.NetworkType = "5G"
	mobile.SimcardType = "Nano"
	mobile.SimcardCapacity = 3

	loptop := Loptop{}

	loptop.Name = "MSI GX 4570"
	loptop.Brand = "MSI"
	loptop.Color = "Black"
	loptop.Ram = 16
	loptop.OperatingSystem = "Ubuntu"
	loptop.OperatingSystemVersion = "22.04"

	loptop.HasCdRom = false
	loptop.UsbPortCount = 3
	loptop.KeyboardType = "Light"

	mobileJson, _ := json.Marshal(mobile)
	loptopJson, _ := json.Marshal(loptop)

	println(string(mobileJson))
	println(string(loptopJson))

}
