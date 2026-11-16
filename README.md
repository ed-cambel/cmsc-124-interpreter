# Arkana

## Creators

- Edriellen Mey Cambel (ed-cambel)
- Rose Antonette Kaindoy (rskaindoy)

## Overview

This project involves designing and implementing an interpreter for a programming language developed as part of the CMSC 124 course. 
[what the language is for, who would use it, what writing it feels like.]

## Host language and build

- Host language: Go 1.27.0
- Module: `arkana`
- Version metadata: `go.mod`
- Build: `./build.sh`
- [Anything a fresh clone needs to know.]

## Running it


| Command | What it does |
|---|---|
| `./run <file>` | Executes a program. Only reveal statements produce output. |
| `./run --tokenize <file>` | Scans a source file and prints its token stream. |
| `./run --parse <file>` | Parses a source file as semicolon-terminated expressions and prints one tree per expression in prefix form. |
| `./run --eval <file>` | Evaluates each expression and prints its value. |
| `./run` | Starts the REPL. Executes statements and keeps state. |


Exit codes: 0 when program runs successfully, 65 when a lexical or syntax error is encountered, 70 when a runtime error occurs.

## File extension

`.arc`

## Lexical structure

### Keywords


| Keyword | Purpose          |
|---|---|
| `grain` | Integer data type |
| `dram` | Floating-point data type |
| `scroll` | String data type |
| `rune` | Character data type |
| `boon` | Boolean value `true` |
| `bane` | Boolean value `false`  |
| `nil` | No-value literal |
| `seal` | Constant declaration |
| `when` |  Runs specific block of code if the condition evaluates to `boon`. |
| `lest` |  Runs alternative block of code if the condition evaluates to `bane` or encounters `nil`. |
| `augur` | Evaluates target variable or expression for matching |
| `sign` | Defines literal value to match against target variable |
| `fate` | Fallback block when no cases match |
| `ritual` | Declares and initializes loop until condition becomes `bane` | 
| `skip` | Skips to the next iteration of loop |
| `dispel` | Terminates loop |
| `reveal` | Evaluates an expression and prints the value |


### Operators


| Operator | Category | Operands | Associativity | Precedence |
|---|---|---|---|---|
| `+` | arithmetic | binary | left | 6 |
| `-` | arithmetic | binary | left | 6 |
| `*` | arithmetic | binary | left | 7 |
| `/` | arithmetic | binary | left | 7 |
| `%` | arithmetic | binary | left | 7 |
| `-` | arithmetic | unary | right | 8 |
| `!` | logical | unary | right | 8 |
| `&&` | logical | binary | left | 3 |
| `\|\|` | logical | binary | left | 2 |
| `->` | assignment | binary | right | 1 |
| `>` | comparison | binary | left | 5 |
| `<` | comparison | binary | left | 5 |
| `==` | comparison | binary | left | 4 |
| `!=` | comparison | binary | left | 4 |
| `>=` | comparison | binary | left | 5 |
| `<=`| comparison | binary | left | 5 |


### Literals


| Kind | Syntax | Produces |
|---|---|---|
| Number | `67`, `3.14`, `.5` | Integer or floating-point value |
| String | `hello\n\tworld`, `quote: \"arkana\"` `yes\\no` | String value |
| Boolean | `boon`, `bane` | `true`  or `false`|
| Nil | `nil` | No value |


### Identifiers

- Start characters: letters (`A-Z`, `a-z`) or `_`
- Continue characters: letters, digits (`0-9`), or `_`
- Case-sensitive: yes
- Identifiers cannot begin with a digit.
- Keywords are lowercase and are recognized only when the complete lexeme exactly matches a reserved keyword.
- [Reserved patterns, length limits, or other restrictions.]

### Comments

- Line / Block comments: enclosed by opening tag (`/>`) and closing tag (`</`)
- Nesting: not supported

## Whitespace and termination

- Whitespace significant: no
- Statement terminator: semicolon (`;`)
- Block delimiters: braces `{}`
- Grouping delimiters: parentheses `()`

## Token output format

```
Token(type=STRING, lexeme="hello", literal=hello, line=1)
```
- type: the token's 'TokenType' as defined constants and printed in string form by `token.go`.
- lexeme: the source text the token was scanned from, including surrounding text (e.g. a string's lexeme includes quotation marks: `"hello"`, a number's lexeme is just its digits: `123`).
- literal: the token's represented value; strings and numbers are considered literals, operators and punctuation have no value and are considered \<nil>\.
- line: the source line the token started on. 

## Grammar

```
program 	  → declaration* EOF
declaration → varDecl | statement
varDecl     → type IDENTIFIER ( "->" expression )? ";" | "seal" IDENTIFIER "->" expression ";"
type        → "grain" | "dram" | "scroll" | "rune"
statement   → revStmt | exprStmt | block
revStmt     → "reveal" expression ";"
exprStmt    → expression ";"
block       → "{" declaration* "}"

expression → assignment
assignment → IDENTIFIER "->" assignment | or
or         → and ( "||" and )*
and        → equality ( "&&" equality )*
equality   → comparison ( ( "==" | "!=" ) comparison )*
comparison → term ( ( ">" | ">=" | "<" | "<=" ) term )*
term       → factor ( ( "+" | "-" ) factor )*
factor     → unary ( ( "*" | "/" | "%" ) unary )*
unary      → ( "!" | "-" ) unary | primary
primary    → NUMBER | STRING | "boon" | "bane" | "nil" | IDENTIFIER | "(" expression ")"

```

## Parse output format

```
(+ 1 2)
```

- Groupings print as: `(group ...)` e.g. `(* (group (+ 2 3)) 4)`
- Numbers print as: whole numbers plainly `4`, numbers with a decimal point keep it `3.12`

## Semantics

### Values and types

| Keyword | Go type | Notes |
|---|---|---|
| `grain` | `int` | - |
| `dram` | `float64` | - |
| `scroll` | `string` | - |
| `boon` / `bane` | `bool` | Literals only. No type keyword yet, so they can't be stored in variables. |
| `nil` | `nil` | Fits any declared type. |
| `rune` | n/a | No literal yet, so variables can only hold `nil`. |

- Fixed type: a variable’s type is set at declaration and cannot change.
- `nil`: assignable to a variable of any type.
- Promotion: `grain → dram` is allowed and converts the value
  - Example: `dram d -> 3`; holds `3.0`.
- Mismatch: `dram → grain` and every other type mismatch is a runtime error (exit 70).

### Value printing

- Numbers: trailing `.0` will be kept as is (e.g `2.0 + 3.0` prints `5.0`); 
- Nil: `nil`
- Strings: without quotes (e.g. `“hello world”` prints `hello world`)
- Non-terminating decimals: printed with up to 15 to 17 significant digits, as held by a `float64` (`1 / 3` prints `0.3333333333333333`)
- Overflow: a result outside the 64-bit int range is a runtime error. `9223372036854775807 + 1` fails with `Integer overflow.`


### Truthiness

- Only `bane`, `nil`, and the number `0` are false. Everything else is true, including `""`.
- `!` applies to any type: `!nil` is `boon`, `!bane` is `boon`, `!boon` is `bane`, `!0` is `boon`.
- `and` / `or` short-circuit and return the deciding operand.

### Operator semantics

- Evaluation order: left operand, then right
- Arithmetic: `- * / %` and unary `-` accept numbers only. Operands are checked before the operation (`checkNumberOperand` for unary, `checkNumberOperands` for binary)
  - `int op int` stays an int for `+ - * %`
  - `int / int` is promoted to a float, (e.g. `5 / 2` is `2.5`);interger division will be added
  - `int op float` converts the int to a float (e.g. `1 + 2.5` is `3.5`)
  - `%` takes the sign of the left operand: `5.5 % 2.5` is `0.5` (`5.5 - 2.5*2`), `-5.5 % 2.5` is `-0.5` (`-5.5 - 2.5*(-2)`), `0.5 % 2` is `0.5`.
- `+` on strings: string + string concatenates; mixed types result in an error (e.g. "a" + 1 is an error)
- Comparison: `< <= > >=` accept numbers only and return a boolean. Strings are not comparable.
- Equality across types: values of different types are not equal, (e.g. `1 == "1"` is `bane`). Values are compared by value, not by reference.
  - `1 == 1.0` is `boon`, and `nil == nil` is `boon`.
  - `nil == bane` is `bane`.
  - `0.1 + 0.2 == 0.3` is `bane`. Floats are compared by exact value, and `0.1 + 0.2` could be stored as `0.30000000000000004`, which is not numerically equal to `0.3`.
  - `-0` and `0` are equal
- Division by zero:  `/` and `%` by zero are runtime errors.


### Scope and bindings

- Paradigm: imperative; bindings are mutable except `seal`.
- Scope: lexical. `{ }` opens a nested scope. Lookup walks outward from the innermost scope and the innermost match wins.
- Declaration: creates a binding in the current scope only and never searches outward. It is legal at top level and inside blocks.
  -  `seal` requires an initializer (`seal x;` is a syntax error, exit 65). Its type is inferred and fixed. Reassigning it is a runtime error (exit 70).
- Assignment: updates the nearest binding and never creates one. It is an expression and yields the assigned value. An invalid target is a syntax error (exit 65).
- Redeclaration in the same scope: runtime error (exit 70)
- Uninitialized variable holds: yes (ex. `grain x;` has a value `nil`)
- Shadowing: an inner declaration hides the outer one without changing it. The outer binding is visible again when the block ends.
- Initializer: evaluated before the binding is defined, so an inner `grain x -> x;` reads the outer `x`.
- Undefined name: runtime error (exit 70) on both read and assign.
- Binding time:
  - The type is bound at declaration and can’t change. 
  - The value is bound at initialization and rebound on assignment (except `seal`). 
  - A name becomes visible after its declaration finishes executing.
- Storage lifetime:
  - Globals have static storage (program lifetime). 
  - Block locals are stack-dynamic: storage begins when the declaration executes and ends when the block exits. 
  - There is no explicit heap-dynamic storage. The Go map behind the environment doesn’t change what the language promises.

### Control flow and functions

- Logical operators return: [booleans, or the operand]
- Dangling else binds to: [which if]
- Closure capture of a loop variable: [per iteration, or shared]
- Function with no return statement produces: [value]
- Arity mismatch: [message and exit code]

## Native functions


| Name | Arguments | Returns | Notes |
|---|---|---|---|
| [name] | [count and types] | [type] | [caveats] |


## Errors and diagnostics

Message format:

```
[Line 2] Error: Unexpected character '#'
[Line 1] Error at ';': Expect expression.
[Line 1] Runtime error: Undefined variable 'y'
```


| Failure | Exit code |
|---|---|
| lexical error | 65 |
| syntax error | 65 |
| runtime error | 70 |


## Testing conventions


| Folder | Activity | Mode | Flag |
|---|---|---|---|
| tests/lab1 | Scanner | sidecar | `--tokenize` |
| tests/lab2 | Parser | sidecar | `--parse` |
| tests/lab3 | Evaluator | inline | `--eval` |
| tests/lab4 | Context | inline | none |
| tests/lab5 | Functions | inline | none |


```
[specific tests]...
```

Run locally with:

```bash
curl -sSL https://raw.githubusercontent.com/WhiteLicorice/cmsc-124-harness/v1.1/run_tests.py -o run_tests.py
./build.sh
python3 run_tests.py tests/lab0
python3 run_tests.py tests/lab1
python3 run_tests.py tests/lab2
python3 run_tests.py tests/lab3
python3 run_tests.py tests/lab4
```

## Sample code

```
/> A witch brews a potion. Her mana is borrowed inside the spell block. </
seal cauldron -> "Moonbrew";
grain mana -> 10;
dram potency -> 1.5;

reveal cauldron;
reveal mana;

{
  /> a new mana that only exists inside this block </
  grain mana -> 3;
  reveal mana;

  /> no new variable here, so this changes the outer potency </
  potency -> potency * 2;
}

/> the block is over: the outer mana is back, and potency stayed changed </
reveal mana;
reveal potency;
```

Output:

```
Moonbrew
10
3
10
3.0
```

## Design rationale

- Empty or comment-only files are treated as success cases. In the grammar section, `program → declaration* EOF` defines zero or more expressions.
- Semicolon (`;`) as terminator, because it makes the end unambiguous and gives the parser a clear point to resume after a syntax error, while newline rules would need line checks in every binary loop.
- Ints and floats kept as separate kinds, with 2 and 2.0 printing differently.
- Lexical scope since a name resolves by where it is written, not by who called, so you can read a program’s lookups straight from the source.
- Fixed types is observed. A variable declared `grain` can’t silently become a `scroll`, which catches wrong-type assignments at the point they happen. `grain` to `dram` is allowed because it loses nothing and matches the existing int-with-float arithmetic. The cost is an extra check on every declaration and assignment, plus the `nil` exception that makes uninitialized variables work.

[Why the language is the way it is. Cover the choices that surprised you, the
features you cut, and the decisions you reversed. Specific reasons, not
approval of your own work.]

## Known limitations

- `--parse` and `--eval` only accept expressions terminated by `;`. 
- `boon/bane` values cannot be stored in variables, since there is no bool type keyword yet.
- `rune` variables can only hold `nil`, since there is no `rune` literal yet.
- [What doesn't work, what is unimplemented, where behavior is worse than you
  would like.]

## Changelog

| Activity | What changed in the language |
|---|---|
| Lab 1 | Established token output format (`TYPE LEXEME LITERAL LINE`) |
| Lab 1 | Changed token output format (`Token(type=TYPE, lexeme=LEXEME, literal=LITERAL, line=LINE)`) |
| Lab 1 | Added scanner REPL through `./run` |
| Lab 2 | Added ; as the expression terminator |
| Lab 2 | Edited REPL to now parse and print trees |
| Lab 3 | Edited REPL to now evaluate |
| Lab 4 | Extended grammar to add reveal statements |
| Lab 4 | ./run <file> now executes programs |
| Lab 4 | Added `{ }` blocks with lexical scope and shadowing |
| Lab 4 | Added identifiers as expressions |
| Lab 4 | Added runtime errors for undefined names, redeclaration, type mismatch, and `seal` reassignment |
| Lab 4 | Added `grain` to `dram` promotion on assignment; any other type mismatch is a runtime error |
| Lab 4 | Edited REPL to now execute statements, keeps variables across lines, and prints bare expression values |