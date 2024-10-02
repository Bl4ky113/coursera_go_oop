# Functions, Methods, and Interfaces in Go
# By University of California, Irvine

Start: 09/11/2024
End: 10/01/2024

Sessions:
- 09/11/2024 I hope that I can finish this one as fast as the other.
- 09/18/2024
- 09/28/2024 It wasn't, but now will be
- 09/29/2024
- 09/30/2024

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

## OOP in Go

Go officialy doesn't have an Object Oriented Programming, 
but we can recreate the objects with properties and methods with Parcial Encapsulation.

This is allowed thanks to functions beeing able to receive a variable by a Receiver, 
which is defined after the func keyword and before the name of the function. 
This will allow us to use the Receiver just like a parameter inside the function, but we 
have a complete different syntax for calling the function:

func (bar type) foo () {
    ...
}

barElement.foo()

This syntax is technically the same as any other OOP language. But the Receiver, since it behaves like a parameter,
there might be problemns with passing it as a pointer or a value, either way we can pass it by pointer by just changing the 
type in the definition.

One main problemn is that, when passed as value, we can't rewrite the content of the receiver outside the function scope.
The derreference and reference of the pointers while making the call or inside the function is not needed,
Golang will do it for us automatically, meaning that we can jus tuse them as a normal variable.

We can encapsulate data as private by just using the mixedCaps notation in Golang, won't be available outside the file, 
but inside is free to use. Also MixedCaps means that the method, attribute or anything is Public for everything.

## Polymorphism

We want to implement multiple classes or structs with a same method but with different execution, like Speak, in a normal OOP language we could 
create a class and make some subclasses by inheritance and then overwritting the method if needed.

In Golang, since we don't have inheritance, we got to use interfaces. These are defined just like structs, the main thing is that we don't
define attributes but methods. 

type foo interface {
    bar(param1 type, param2 type, ...) rtnType
    ...
}

And if we define a struct in the same file and define each method inside the interface the compiler will automatically 
know that the struct is using the interface.

The curious thing is that we can create a variable with the type of the interface. 
This will allow us to define another variable which is an struct implementing the interface 
and asign that one to the interface var. 

Now, if we only used the interface var, it would do nothing at all, and trying to call any method, attribute or anything at all
it would do a runtime error and panic.

The main usage of this interfaces is in the case that we have a lot of somewhat same structures and we need to implement 
a method on every single one. Using only structs and methods, we would have to rewrite every single method and in most 
cases almost the same way for every single struct. 

But with interfaces, we can define an interface for all thoose structs and then, like structs, we can define a 
method for the objects or structs that implement that interface, by passing to a function a Receiver with the 
interface type. Mainly in this interfaces we might want to check first if the passed struct is NIL in case 
that we passed only an empty interface reference and not a struct that implements it. 

Also in this interfaces methods we might want to differenciate types in the methods. We can do it by 
accessing the type of the variable by using:

foo.(type)

This works with ANY type, so we can use it in a switch statement in order to do stuff acordingly:

switch varType := var.(type) {
    case foo:
        ...
    ...
    default:
        ...
}

We can also the interface methods to make a method for all types and structs by just defining the Receiver as an 
empty interface, or one without methods, this might be usefull mainly for logging and debuging stuff.

## Error handling 

Up until now you might have noticed that we event havent touched the error handling in Golang, 
this is because there's next to none, but a little bit of the stuff that we have learnt up until now
we can use it to implement a error handling system.

Mainly, in almost every function there's the normal return value, if there's any, and 
the error return. Which tell us if there has been an error while executing the function.
This also mean that if we didn't have any error, the return value of the error will be NIL.
Generally these errors are made with the interface error:

type error interface {
    Error() string
}

Then the basic error handling would be checking if the error is NIL, and if it isn't is, 
we have to handle it acordingly, mainly only printing it and see what the hell happend.

## Final

It was a good course, nothing more to say
