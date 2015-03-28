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
// <expression> := <number> (<operator> <number>)?
// <variable> := <variable_name> | <number> | <string>
package strict
