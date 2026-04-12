# Interview-useful language dimensions (checklist)

Use as a **menu**, not a mandatory dump. Pick items that match the user’s question and language.

## Core mechanics

- Lexical structure: keywords, identifiers, literals, operators, precedence
- Statements vs expressions; blocks and indentation/braces
- Comments and docstrings: purely non-semantic vs tools that read them
- Imports/modules/packages; initialization order

## Types and values

- Static vs dynamic typing; gradual typing if applicable
- Value vs reference semantics; identity vs equality
- Coercion / casting rules; truthiness and falsy sets
- `None` / null / optional types; sentinels

## Functions and call model

- Arguments: positional, keyword, defaults, `*args`, `**kwargs` (or language equivalent)
- Closures; capture by name vs by value; late binding gotchas
- Lambdas / anonymous functions limits
- Decorators, wrappers, higher-order functions

## Data structures

- Built-in collections: complexity guarantees where the language specifies them
- Mutability; shallow vs deep copy; views
- Comprehensions / generators / iterators / coroutines (and memory implications)

## OOP and protocols

- Classes, instances, inheritance, MRO (if relevant)
- Special methods / operator overloading / protocols or interfaces
- `self` / `this` binding

## Concurrency and async

- Threads vs processes vs async model for this language
- GIL or equivalent; when parallelism is real vs cooperative
- Race conditions; locks; deadlocks (high level)

## Errors and resources

- Exceptions vs error values; exception hierarchy
- `finally` / context managers / defer patterns
- Resource cleanup guarantees

## Metaprogramming and reflection

- `eval`, introspection, code objects, macros (if any)

## Tooling chain (light touch)

- Interpreter/compiler, bytecode, optimizer flags (only when it clarifies behavior the user asked about)

## How to use in a session

1. If the user’s question maps to one bullet cluster, go deep there.
2. If **Gaps** in the concept file lists a cluster, offer it as the next step.
3. After covering a cluster well enough for interviews, note it in **Detailed notes** and trim **Gaps** accordingly.
