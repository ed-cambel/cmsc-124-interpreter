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
| `./run <file>` | [Executes a program. Available from Lab 4.] |
| `./run --tokenize <file>` | Scans a source file and prints its token stream. |
| `./run --parse <file>` | [Prints the parsed tree.] |
| `./run --eval <file>` | [Evaluates each expression and prints its value.] |
| `./run` | Starts the scanner REPL. |


Exit codes: 0 when program runs successfully, 65 when semantic errors encountered, 70 [when].

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
| `lest` |  Runs altervative block of code if the condition evaluates to `bane` or encounters `nil`. |
| `augur` | Evaluates target variable or expression for matching |
| `sign` | Defines literal value to match against target variable |
| `fate` | Fallback block when no cases match |
| `ritual` | Declares and initializes loop until condition becomes `bane` | 
| `skip` | Skips to the next iteration of loop |
| `dispel` | Terminates loop |


### Operators


| Operator | Category | Operands | Associativity | Precedence |
|---|---|---|---|---|
| `+` | arithmetic | binary | TBD | TBD |
| `-` | arithmetic | binary | TBD | TBD |
| `*` | arithmetic | binary | TBD | TBD |
| `/` | arithmetic | binary | TBD | TBD |
| `%` | arithmetic | binary | TBD | TBD |
| `!` | logical | unary | TBD | TBD |
| `&&` | logical | binary | TBD | TBD |
| `\|\|` | logical | binary | TBD | TBD |
| `->` | assignment | binary | TBD | TBD |
| `>` | comparison | binary | TBD | TBD |
| `<` | comparison | binary | TBD | TBD |
| `==` | comparison | binary | TBD | TBD |
| `>=` | comparison | binary | TBD | TBD |
| `<=`| comparison | binary | TBD | TBD |
| [op] | [arithmetic, comparison, logical, assignment, other] | [unary or binary] | [left, right, none] | [1 = loosest] |


### Literals


| Kind | Syntax | Produces |
|---|---|---|
| Number | `67`, `3.14`, `.5` | Integer or floating-point value |
| String | `hello\n\tworld`, `quote: \"arkana\"` `yes\\no` | String value |
| Boolean | `boon`, `bane`] | `true`  or `false`|
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
- [Harness note: comment_prefix in tests/lab*/manifest.json is set to the
  token above.]

## Whitespace and termination

- Whitespace significant: [yes or no, and where]
- Statement terminator: [e.g. semicolon, newline, none]
- Block delimiters: [e.g. braces, indentation]
- Grouping delimiters: [e.g. parentheses]

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
[Your complete context-free grammar, current as of the latest activity.
Unambiguous, with precedence and associativity encoded in rule structure.]
```

## Parse output format

```
[one line of real --parse output, e.g. (+ 1.0 (* 2.0 3.0))]
```

- Groupings print as: [form]
- Numbers print as: [form]

## Semantics

### Values and types

[What runtime values exist, and how they are represented in the host
language.]

### Value printing

- Numbers: [e.g. 5 rather than 5.0]
- Nil: [spelling]
- Strings: [with or without quotes]

### Truthiness

[The complete rule. Which values are false in a condition; everything else is
true.]

### Operator semantics

- Arithmetic: [accepted operand types]
- `+` on strings: [concatenation, error, or coercion]
- Mixed types: [what happens]
- Comparison: [accepted operand types]
- Equality across types: [false, or an error]
- Division by zero: [value produced, or runtime error]

### Scope and bindings

- Redeclaration in the same scope: [allowed or an error]
- Uninitialized variable holds: [value]
- Shadowing: [behavior]
- Undefined name: [static error with exit 65, or runtime error with exit 70]

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
[one real runtime error]
```


| Failure | Exit code |
|---|---|
| [lexical error] | 65 |
| [syntax error] | 65 |
| [runtime error] | 70 |


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
python3 run_tests.py tests/lab1
```

## Sample code

```
[a short program]
```

Output:

```
[its output]
```

## Design rationale

[Why the language is the way it is. Cover the choices that surprised you, the
features you cut, and the decisions you reversed. Specific reasons, not
approval of your own work.]

## Known limitations

- The REPL currently supports lexical scanning only; parsing and evaluation are not yet implemented.
- [What doesn't work, what is unimplemented, where behavior is worse than you
  would like.]

## Changelog


| Activity | What changed in the language |
|---|---|
| Lab 1 | Established token output format (`TYPE LEXEME LITERAL LINE`) |
| Lab 1 | Changed token output format (`Token(type=TYPE, lexeme=LEXEME, literal=LITERAL, line=LINE)`) |
| Lab 1 | Added scanner REPL through `./run` |
