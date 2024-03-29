# GoBridge Workshop

Notes and code from a GoBridge workshop: https://idiomat.gitbook.io/gobridge/

## Notes

* Before 1.13 go was opinionated about where go code lived. If $GOPATH is not set, it expexts the code to be in ~/go
* Go has default values, called a "zero value", for each type. For a string it is an empty string, for an int it is 0.
* Use `:=` to simultaneously declare and assign a new variable. This is called declaration in assignment.
  * If you attempt to declare a var again, you get an error that says "no new varibels on left side of :=".
  * You cannot use `:=` outside of `main`
* Strings defined with back ticks interprets the contents literally. Double quotes are not like this.
* Double quotes allow for interpolation, backticks do not.
* Single quotes can only contain one character. This is called a rune. Runes are an alias for [int32](https://devdocs.io/go/builtin/index#int32)
* Variables declared outside of the function body are global.
* `package` at the beginning of a file is used to say that different files are in the same scope.
* Uppercased identifiers like `Name` are exported (aka visible outside of the package when the package is imported). Lowercase are not.
* when using slice's `append`, you need to assign the result to a variable. Otherwise the original variable still points to the old data.

## `GOPATH`

I got confused about GOPATH while working on this. If you compare the direcotry structure here to how the [workshop](https://idiomat.gitbook.io/gobridge/your-first-go-program#action-steps-2) says to set up the code, you'll notice I did a few things different.

`$GOPATH` is meant to contain sources that have Go code in them. If you are trying to import a library, like I did in [gettingStarted/main.go](gettingStarted/main.go), you need to set the `GOPATH` to the current directory.

You can do this with ``export GOPATH=`pwd``. Since I set this code up with multiple directories, I made a bash function in my bash profile so I could quickly change my `GOPATH`.

```bash
function setgopath() {
 export GOPATH=`pwd` && echo "GOPATH is set to $GOPATH"
}
```

## Resources

* https://godoc.org
* https://play.golang.org/
* https://gobyexample.com/
