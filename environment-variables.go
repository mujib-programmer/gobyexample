// Environment variables are a universal mechanism for conveying configuration information to Unix programs.
// Let’s look at how to set, get, and list environment variables.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv. This will return an empty string if the key isn’t present in the environment.
	os.Setenv("Foo", "1")
	fmt.Println("Foo:", os.Getenv("Foo"))
	fmt.Println("Bar:", os.Getenv("Bar"))

	// Use os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	// You can strings.Split them to get the key and value. Here we print all the keys.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

// My Output:
//
// Running the program shows that we pick up the value for FOO that we set in the program, but that BAR is empty.
// The list of keys in the environment will depend on your particular machine.
/*
$ go run environment-variables.go
Foo: 1
Bar:

LD_LIBRARY_PATH
LS_COLORS
LC_MEASUREMENT
...
*/
//
// If we set BAR in the environment first, the running program picks that value up.
/*
$ Bar=2 go run environment-variables.go
Foo: 1
Bar: 2

Bar
LD_LIBRARY_PATH
LS_COLORS
LC_MEASUREMENT
...
*/
