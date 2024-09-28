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
