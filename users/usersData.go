package users

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}

// var Users = []User{
// 	{ID: 1, Name: "John", LastName: "Doe", Age: 30},
// 	{ID: 2, Name: "Jane", LastName: "Doe", Age: 25},
// 	{ID: 3, Name: "Alice", LastName: "Smith", Age: 28},
// 	{ID: 4, Name: "Bob", LastName: "Brown", Age: 35},
// }
