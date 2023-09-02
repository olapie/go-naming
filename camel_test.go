package naming

import "testing"

func TestToCamel(t *testing.T) {
	type TestCase struct {
		Input  string
		Output string
	}
	tests := []TestCase{
		{"-----", ""},
		{"---._ \t", ""},
		{"--b-._ \ta", "bA"},
		{"b--b-._ \ta", "bBA"},
		{"b--bC-._ \ta", "bBCA"},
		{"bef--bC-._ \ta", "befBCA"},
		{"xml___http", "xmlHTTP"},
		{"xml.http", "xmlHTTP"},
		{"_xml.http", "xmlHTTP"},
		{"__xml.http__", "xmlHTTP"},
		{"URLString", "urlString"},
	}
	for _, test := range tests {
		output := ToCamel(test.Input)
		if output != test.Output {
			t.Fatalf("Test %s, got %s, want %s", test.Input, output, test.Output)
		}
	}
}

func TestToCamelWithoutAcronym(t *testing.T) {
	type TestCase struct {
		Input  string
		Output string
	}
	tests := []TestCase{
		{"-----", ""},
		{"---._ \t", ""},
		{"--b-._ \ta", "bA"},
		{"b--b-._ \ta", "bBA"},
		{"b--bC-._ \ta", "bBCA"},
		{"bef--bC-._ \ta", "befBCA"},
		{"xml___http", "xmlHTTP"},
		{"xml.http", "xmlHTTP"},
		{"_xml.http", "xmlHTTP"},
		{"__xml.http__", "xmlHTTP"},
		{"URLString", "urlstring"},
	}
	for _, test := range tests {
		output := ToCamel(test.Input, WithoutAcronym())
		if output != test.Output {
			t.Fatalf("Test %s, got %s, want %s", test.Input, output, test.Output)
		}
	}
}
