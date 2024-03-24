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
	examplePascalToKebabWithAcronym()
	exampleCustomAcronym()
}

func exampleCamelToSnake() {
	camel := "httpRequestId"
	fmt.Println(naming.ToSnake(camel)) // output: http_request_id
}

func examplePascalToSnake() {
	pascal := "HttpRequestId"
	fmt.Println(naming.ToSnake(pascal)) // output: http_request_id
}

func examplePascalToSnakeWithAcronym() {
	pascal := "HTTPRequestID"
	fmt.Println(naming.ToSnake(pascal, naming.WithAcronym())) // output: http_request_id
}

func exampleSnakeToCamel() {
	snake := "http_request_id"
	fmt.Println(naming.ToCamel(snake)) // output: httpRequestId
}

func exampleSnakeToCamelWithAcronym() {
	snake := "http_request_id"
	fmt.Println(naming.ToCamel(snake, naming.WithAcronym())) // output: httpRequestID
}

func exampleSnakeToPascal() {
	snake := "http_request_id"
	fmt.Println(naming.ToPascal(snake)) // output: HttpRequestId
}

func exampleSnakeToPascalWithAcronym() {
	snake := "http_request_id"
	fmt.Println(naming.ToPascal(snake, naming.WithAcronym())) // output: HTTPRequestID
}

func examplePascalToKebab() {
	pascal := "HttpRequestId"
	fmt.Println(naming.ToKebab(pascal)) // output: http-request-id
}

func examplePascalToKebabWithAcronym() {
	pascal := "HTTPRequestID"
	fmt.Println(naming.ToKebab(pascal)) // output: http-request-id
}

func exampleCustomAcronym() {
	pascal := "CUSTOMAcronym"
	naming.AddAcronym("CUSTOM")
	fmt.Println(naming.ToSnake(pascal, naming.WithAcronym())) // output: custom_acronym
}
