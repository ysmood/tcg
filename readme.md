# Overview

The proposal borrowed some ideas from [The check/handle proposal](https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling.md).
One of the issues of "check/handle proposal" is that it doesn't support function chaining well: https://gist.github.com/DeedleFake/5e8e9e39203dff4839793981f79123aa.

To make a long story short, let's see some code examples first.

Without the proposal:

```go
type data struct {}

func (d data) bar() (string, error) {
    return "", errors.New("err")
}

func foo() (data, error) {
    return data{}, errors.New("err")
}

func do () (string, error) {
    d, err := foo()
    if err != nil {
        return "", err
    }

    s, err := d.bar()
    if err != nil {
        return "", err
    }

    return s, nil
}
```

With the proposal:

```go
type data struct {}

func (d data) bar() string {
    // The throw is like the return, you have to state the origin return values with an extra error.
    // The last value of the list must fulfil error interface, or compile error.
    throw "", errors.New("err")
}

func foo() (d data) {
    // Same as the return, in this case we can omit the return values.
    throw errors.New("err")

    // Throw is like return, this line is unreachable.
    return
}

func do () (string, error) {
    // The catch is similar to the handle in the check/handle proposal.
    catch err { // the type of err is error
        return "", err // you must return corresponding values just like a normal function
    }

    // The error will propagate like the panic until it's caught by a `catch` clause.
    s := foo().bar()
    return s, nil
}
```

Use keyword `guard` to convert the throw version to normal error value style, so that we can precisely handle each error separately:

```go
func do () (string, error) {
    d, err := guard foo()
    if err != nil {
        return "", err
    }

    s, err := guard d.bar()
    if err != nil {
        return "", err
    }

    return s, nil
}
```

Currently, I created a demo lib to simulate the taste: https://github.com/ysmood/tcg/blob/master/example_test.go
We shouldn't just use the `panic` to do it. For the implementation of `throw`, it's similar to [setjmp and longjmp](https://en.wikipedia.org/wiki/Setjmp.h). I feel the performance won't be affected.

To avoid forgetting `catch` we can statically do something similar to https://github.com/kisielk/errcheck by analyzing `throw` keyword.

Sure the keywords are not decided, we can use other words if they are better.
