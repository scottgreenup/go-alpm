package alpm

import (
    "fmt"
    "testing"
)

func TestStandard(t *testing.T) {
    handle, err := NewHandle("/", "/var/lib/pacman")
    if err != nil {
        panic(err)
    }
    defer handle.Close()

    database, err := NewDatabase(handle)
    if err != nil {
        panic(err)
    }

    fmt.Println(database)

    packages, err := database.Packages()
    if err != nil {
        panic(err)
    }

    for _, p := range packages {
        fmt.Println(p.Name())
    }
}
