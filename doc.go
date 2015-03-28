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
	<program> := <scope>+
	<scope> := <scope_name>? | "{" <line> "}"
	<line> := <statement> | <control>

	<evaluate_scope> := <args>? <scope> "()"
	<args> := <list> ">>"
	<list> := "[" <list_value> "]"
	<list_value> := <value> ("," <list_value>)
	<int> := / \d* /
	<float> := / \d*.\d* /
	<number> := <int> | <float>
	<string> := "\"" / .* / "\""
	<variable_name> := / [a-z][a-zA-Z0-9]*  /
	<value> := <number> | <string> | <evaluate_scope> | <variable_name>

	<add> := "+"
	<subtract> := "-"
	<divide> := "/"
	<multiply> := "*"
	<operator> := <add> | <subtract> | <divide> | <multiply>

	<or> := "or"
	<and> := "and"
	<more> := <or> | <and>

	<statement> := <assignment>? <expression>
	<assignment> := <variable_name> "="
	<expression> := <value> (<operator> <value>)? (<more> <expression>)?

	<while> := "while" <expression> <scope>
	<if> := "if" <expression> <scope>
	<control> := <while> | <if>

	<scope_name> := /[A-Z][a-zA-Z0-9]+/
*/
/*
{
	string_variable = "Hello"
	integer_variable = 2

	// output: Hello 2
	println(string_variable, integer_variable)


	int i = number()

	// output: 2
	println(i)

	int c = [2, 3] >> add()

	// output: 5
	println(c)

	int z = {
		// computations and stuff
		return 10
	}()

	// output: 10
	println(z)

	// output: hello
	hello()

	int d = 0
	int e = 3

	// output: hi hi hi
	while d < e {
		println("hi")
		d = d + 1
	}

	int x = 2

	// output: eh eh
	while x < 2 and true {
		println("eh")
		x = x + 1
	}
}

number {
	return 2
}

[int a, int b] >> add{
	return a + b
}

hello {
	println("hello")
}
*/
