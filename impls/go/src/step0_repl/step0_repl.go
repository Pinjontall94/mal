package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)


func inputPrompt(label string) string {
    // NOTE: Figure out why this works
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label)
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}


func READ(x string) string {
    return x
}


func EVAL(x string) string {
    return x
}


func PRINT(x string) string {
    return x
}


func rep(x string) string {
    res_read := READ(x)
    res_eval := EVAL(res_read)
    res_print := PRINT(res_eval)
    return res_print
}


func main() {
    for {
        user_input := inputPrompt("user> ")
        res_rep := rep(user_input)
        fmt.Printf("%v\n", res_rep)
    }
}
