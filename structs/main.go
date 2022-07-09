package main

func main() {
	alex := NewPerson(
		"Alex",
		"Anderson",
		ContactInfo{
			email:   "alexanderson@email.com",
			zipCode: 12345,
		},
	)

	alex.print()

	alex.updateName("Alexa")

	alex.print()
}
