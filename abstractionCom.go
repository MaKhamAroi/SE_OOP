package main

import "fmt"

type (
	Computer interface {
		Input()
	}

	ComputerPC struct {
		Brand       string
		CPU         string
		RAM         string
		OS          string
		PowerSupply string
	}

	ComputerLaptop struct {
		Brand      string
		CPU        string
		RAM        string
		OS         string
		Battery    string
		IsCanPress bool
	}
)

func (m ComputerPC) Input() {
	fmt.Println("Starting computer pressing turn on the computer PC:",m.Brand, m.CPU, m.RAM, m.OS, m.PowerSupply)
}

func (m ComputerLaptop) Input() {
	fmt.Println("Starting computer pressing turn on the computer laptop:",m.Brand, m.CPU, m.RAM, m.OS, m.Battery, m.IsCanPress)
}

func main() {
	computerPC := ComputerPC{
		Brand:       "Dell",
		CPU:         "Intel Core i7",
		RAM:         "16GB",
		OS:          "Windows, MacOS",
		PowerSupply: "65W",
	}

	computerLaptop := ComputerLaptop{
		Brand:      "HP",
		CPU:        "AMD Ryzen 5",
		RAM:        "8GB",
		OS:         "Windows, Linux",
		Battery:    "56Wh",
		IsCanPress: true,
	}

	computerPC.Input()
	computerLaptop.Input()
}
