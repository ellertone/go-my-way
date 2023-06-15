package main

import "fmt"

// var tokenS2 = 1000 // how to declare outside of a fuction

const LoginToken string = "Login" //public due tp capitalisation

func main() {
	var username string = "Ellertone"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal uint = 32
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFlt float32 = 32.5489532684
	fmt.Println(smallFlt)
	fmt.Printf("Variable is of type: %T \n", smallFlt)

	// default values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type
	var website = "techwalk.hasnode.dev"
	fmt.Println(website)

	//no var style
	noOfUsers := 30000
	fmt.Println(noOfUsers)

	//using global Loging token
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
