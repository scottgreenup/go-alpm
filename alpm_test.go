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
        fmt.Printf("Name      : %s\n", p.Name())
        fmt.Printf("Desc      : %s\n", p.Desc())
        fmt.Printf("Base      : %s\n", p.Base())
        fmt.Printf("Version   : %s\n", p.Version())
        fmt.Printf("\n")
    }
}
