This package provides functionalities to convert names between different styles(camelCase, PascalCase, snake_case, kebab-case). 
The implementations is based on state machine which is more efficient to process characters.  

Use Acronym for Acronym/Initialism  
* Abbreviation: Dec. (December)
* Acronym: ID (identifier)
* Initialism: ID (identity document) 


## Module path  

    go.olapie.com/naming


## Examples

``` 

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

```