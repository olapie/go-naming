package main

import (
	"fmt"
	"go.olapie.com/naming"
)

func main() {
	exampleCamelToSnake()
	examplePascalToSnake()
	examplePascalToSnakeWithAcronym()
	exampleSnakeToCamel()
	exampleSnakeToCamelWithAcronym()
	exampleSnakeToPascal()
	exampleSnakeToPascalWithAcronym()
	examplePascalToKebab()
}

func exampleCamelToSnake() {
	camel := "userId"
	fmt.Println(naming.ToSnake(camel)) // output: user_id
}

func examplePascalToSnake() {
	pascal := "UserId"
	fmt.Println(naming.ToSnake(pascal)) // output: user_id
}

func examplePascalToSnakeWithAcronym() {
	//naming.AddAcronym("ID")
	pascal := "HTTPRequestID"
	fmt.Println(naming.ToSnake(pascal, naming.WithAcronym())) // output: user_id
}

func exampleSnakeToCamel() {
	snake := "user_id"
	fmt.Println(naming.ToCamel(snake)) // output: userId
}

func exampleSnakeToCamelWithAcronym() {
	snake := "user_id"
	fmt.Println(naming.ToCamel(snake, naming.WithAcronym())) // output: userID
}

func exampleSnakeToPascal() {
	snake := "user_id"
	fmt.Println(naming.ToPascal(snake)) // output: UserId
}

func exampleSnakeToPascalWithAcronym() {
	snake := "user_id"
	fmt.Println(naming.ToPascal(snake, naming.WithAcronym())) // output: UserID
}

func examplePascalToKebab() {
	pascal := "UserID"
	fmt.Println(naming.ToKebab(pascal)) // output: user_id
}
