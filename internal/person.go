package internal

type Person struct {
	ID 			int 	`json:"id"`
	Name		string	`json:"name"`
	LastName	string	`json:"lastName"`
	Address		string 	`json:"address"`
	Dni			string	`json:"dni"`
	Cell		string	`json:"cell"`
}