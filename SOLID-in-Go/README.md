# SOLID in Go

### Single Responsibility Principle

"A class should have one and only one reason to change, meaning that a class should have only one job"

- one and only one job for `type` and `function`

### Open/Closed Principle

"Objects or entities should be open for extension, but closed for modification"

- a `type` / `func` should not guess the input, it should work without modification

### Liskov Substitution Principle

"Let q(x) be a property provable about objects of x of type T. Then q(y) should be provable
for objects y of type S where S is a subtype of T"

- Any sub  type of a specific type can be considered as the type it was created from

### Interface Segregation Principle

"A client should never be forced to implement an interface that it doesn't use or clients shouldn't be forced
to depend on methods they do not use"

- Keep interfaces simple, preferably just one method

### Dependency Inversion Principle

"Entities must depend on abstractions not on concretions"

- Do not depend on specific type, accept abstractions or interfaces instead

### Plan

1. Intro to yourself & 5m Friday
2. Explain what is SOLID in general terms
2.1 Lay out all 5 principles that we'll cover in this video
3. SRP explanation
4. SRP example
5. OCP explanation
6. OCP example
7. LSP explanation
8. LSP example
9. ISP explanation
10. ISP example
11. DIP explanation
12. DIP example
13. The end (Outro) & Mention resources

### Resources

- [Principles of Object Oriented Design](https://scotch.io/bar-talk/s-o-l-i-d-the-first-five-principles-of-object-oriented-design)
