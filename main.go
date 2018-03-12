package main

type Saiyan struct {
	Name string
	Power int
	Father *Saiyan
}
// func for  struct
// goku.Super()
func (s *Saiyan) Super(){
	s.Power += 1111111
}

func main() {
	gohan := &Saiyan {
		Name: "Gohan",
		Power: 10000,
		Father: &Saiyan {
			Name: "Goku",
			Power: 9001,
			Father: nil,
		},
	}

	println(gohan.Father.Power)

	goku := &Saiyan {
		Person: &Person{"Goku"},
		Power: 9001,
	}

	goku.Introduce()
}

func oldMain() {
	goku := new(Saiyan)
	goku.Name = "goku"
	goku.Power = 6060
	goku.Father = nil
	// same
	// goku := &Saiyan {
	//	Name: "goku",
	//	power: 6060,
	// }
	println(goku.Power)
}

func NewSaiyan(name string, power int) Saiyan {
	return Saiyan {
		Name: name,
		Power: power,
	}
}


type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}
