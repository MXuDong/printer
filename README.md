# Printer for golang

Print the color info.

## Usage
1. Import the printer
2. Init a printer instance
3. Invoke printer methods.

## Output common

### common method
```go
pi := DefaultStdoutPrinter()

pi.Print(formatStr, objs...) // output info
pi.ErrorE(err) // output error
```

### custom printer

#### set pre func

post, beauty func like this.

```go
// set wil PrefixStrFunc


// PreFuncWithLineNumber example: For deal with the line number
// If you need log rotate, you can reset prefunc by this func, to implement: to write different file from zero line
func PreFuncWithLineNumber() PrefixStrFunc {
    // if you use the line number mod, the pre and hou func can't have any \n
    count := 0
    return func(bePrintStr string) string {
        count++
        return strconv.Itoa(count)
    }
}
```

# Output format :
`_, _ = fmt.Fprintf(p.Out, string(preColor)+p.PrefixStrFunc(line)+line+p.PostfixStrFunc(line)+string(endColor)+"\n")`
If has more than one line info, it will be split.