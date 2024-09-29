# Functions, Methods, and Interfaces in Go
# By University of California, Irvine

Start: 09/11/2024
End:

Sessions:
- 09/11/2024 I hope that I can finish this one as fast as the other.
- 09/18/2024
- 09/29/2024 It wasn't, but now will be

## Functions basics 

The general stuff:
- Reusability
- Abstraction

The structure for a Go function is

func foo (foo1 type, foo2, foo3 type, ...) (type, type, ...) {
    ...
    return bar, bar, ...
}

In Go we have the normal pass by value parameter and the pass by reference parameter.
There's a bunch of discussion on which should we generally use since, in C, there's a lot more
optimizations for using passing by reference than passing by value. But there's some sort of optimizations
when passing a parameter by value made by the compiler.

### Functions as First class elements
In functional programming, functions are used just like types. Meaning that, we use then as argument/return values,
created dynamically, stored in structures and such.

The thing is that, mainly in return and args, we have to declare the whole stuff that the function will take and return.
But this behaviour will help us a lot.

fn := func bar (arg type, arg type, ...) (returntype) {
    ...
}

Keep in mind that if we define a function as a variable, the function scope will be able to 
get all the values defined at the same scope of the variable with the function, not 
only the function scope itself.

### Anonymous Functions
As we can pass the functions as args/return values, we can just write the definition of the function instead of 
defining one for the file or project scope.

foo(func bar (arg type) (type) { ... }, ...)

### Variadic Argument Number
When we have a function and we might have 1 to inf arguments of a type 
we might only use a slice as a argument, but we can also use the variable or variadic arguments.
By using elipsis at the begining of the type we declare that we are going to we from 1 to multiple arguments.
Meaning that the argument will technically a slice but we won't have to pass one when we call the function.

func foo (bar ...int) {
    print(bar)
    ...
}

foo(0, 1, 3, 4, 5)

the output will be the slice: [0, 1, 3, 4, 5]

If we need to pass a slice to a variadic argument, we will use the elipsis at the end of the slice

foo([]int{0,1,2,3}...)

### Deferred function call
Is a way to call functions after the current block of code ends it's execution. 
We just use the defer keywword and call the function after it.

But you better plan ahead that if you passed or used somestuff in the call that might change after the 
defer call, they won't affect the defer call. 
Meaining that the defer call copies or creates the context of the scope when the defer function was created
and then making the function call with that same copy.
