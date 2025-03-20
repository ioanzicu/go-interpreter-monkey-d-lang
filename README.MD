# The Monkey D. Programming Language Interpreter

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