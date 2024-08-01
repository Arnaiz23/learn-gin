package users

type User struct {
	Id       int
	Name     string
	LastName string
	Age      int
}

var Users = []User{
	{Id: 1, Name: "John", LastName: "Doe", Age: 30},
	{Id: 2, Name: "Jane", LastName: "Doe", Age: 25},
	{Id: 3, Name: "Alice", LastName: "Smith", Age: 28},
	{Id: 4, Name: "Bob", LastName: "Brown", Age: 35},
}
