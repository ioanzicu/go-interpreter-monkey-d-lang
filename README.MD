# The MONKEY D. Programming Language Interpreter

| Source Code | -> | Tokens | -> | Abstract Syntax Tree |


## Major Parts

- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluator


## Features
- C-like syntax
- variable bindings
    ```
        // string
        let name = "Ioan";

        // int
        let age = 23;    
        let total = (33 * (10 * 10)) + 3;
        
        // array
        let pets = ["Barni", "Mandarin", "Dora", "Rosa"];
        pets[1] // "Mandarin"

        // hash 
        let ioan = {"name": "Ioan", "age": 23};
        ioan["age"] // 23
    ```
- integers and booleans

- arithmetic expressions
- built-in functions
    ```
        // bind functions to names
        let double = fn(a) { return a * 2 };
        add(2) // 4

        // recursive call
        let fibinacci = fn(x) {
            if (x == 0) {
                0
            } else {
                if (x == 1) {
                    1
                } else {
                    fibonacci(x - 1) + fibonacci(x - 2);
                }
            }
        };
    ```
- first-class and higher-order functions - functions that take other functions as arguments
    ```
        let twice = fn(f, x) {
            return f(f(x));
        };

        let addTwo = fn(x) {
            return x + 2
        };

        twice(addTwo, 2); // 6
    ```
- closures
- a string data structure
- an array data structure
- a hash data structure



## Parsing

- From root node to childern

- "Top-down operator precedence" - recursive descent parser - "Pratt parser"

 - [BNF for EcmaScript.jj](https://tomcopeland.blogs.com/EcmaScript.html#prod7)


## Expressions In MONKEY D.

- Everything besides `let` and `return` statements is an expression

    - Prefix Operators
        - `-5`
        - `!true`
        - `!false`

    - Infix operators or "binary operators"
        - `5 + 5`
        - `5 - 5`
        - `5 / 5`
        - `5 * 5`
    
    - Comparison Operators
        - `foo == bar`
        - `foo != bar`
        - `foo < bar`
        - `foo > bar`

    - Parentheses to group expressions and influence the order of evaluation
        - `5 * (5 + 5)`
        - `((5 + 5) * 5) * 5`
    
    - Call expressions
        - `add(2, 3)`
        - `add(add(2, 3), add(5, 10))`
        - `max(5, add(5, (5 * 5)))`

    - Identifiers expressions
        - `foo * bar / foobar`
        - `add(foo, bar)`

    - Function literals are expressions
        - `let add = fn(x, y) { return x + y };`
    
    - Function literal in place of an identifier
        - `fn(x, y) { return x + y }(5, 5);`
        - `(fn(x) { return x }(5) + 10) * 10;`
    
    - if expression
        - `let result = if (10 > 5) { true } else { false };`
        
            `result // => true`
