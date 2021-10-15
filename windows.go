package main

import "fmt"

type windows struct {
}

func (w *windows) print() {
    fmt.Println("Print request for windows")
}

func (w *windows) setPrinter(p printer) {
    w.printer = p
}