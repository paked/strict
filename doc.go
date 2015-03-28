// Package strict is a prototype for a programming language...
// Grammar!
// <program> := <scope>
// <scope> := <name>? "{" <line>+ "}"
// <line> := <statement> | <expression> | <control>
// <control> := ("while" | "if") <expresion> <scope>
// <declaration> := <type_name> <statement>
// <statement> := <variable_name> "=" <expression>
// <type_name> := /[A-Z][a-zA-Z]*/
// <variable_name> := /[a-z][A-Za-z0-9]/
// <string> := "\"" /\w*/ "\""
// <integer> := \d+
// <float> := \d+ "." \d+
// <number> := <integer> | <float>
// <expression> := <variable> (<operator> <variable>)?
// <variable> := <variable_name> | <number> | <string>
package strict

/*
{
	string_variable = "Hello"
	integer_variable = 2

	// output: Hello 2
	println(string_variable, integer_variable)


	i = number()

	// output: 2
	println(i)

	c = [2, 3] >> add()

	// output: 5
	println(c)

	z = {
		// computations and stuff
		return 10
	}()

	// output: 10
	println(z)

	// output: hello
	hello()
}

number {
	return 2
}

add [int a, int b] {
	return a + b
}

hello {
	println("hello")
}
*/
