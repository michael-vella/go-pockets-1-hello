# Description

First project from the **Learn Go with Pocket-Sized Projects** book by Aliénor Latour, Donia Chaiehloudj, and Pascal Bertrand.

This project extends the usual *Hello World* program and using Go's standard *flag* package allows the user to print the equivalent of *Hello World* in a language of their choosing.

# Examples

Running `go run main.go -lang=mt` prints **Hello dinja** to the console which is the equivalent of **Hello world** but in the Maltese language. The default language chosen is English so running the program without the `lang` flag (`go run main.go`) will print out the default **Hello world**. Other supported languages include French (`fr`) and Spanish (`sp`).

Finally, `go test` will run all the tests associated with the project.
