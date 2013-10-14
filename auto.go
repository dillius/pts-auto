package main

import (
       "fmt"
       "os"
       "strconv"
)

func updateUser(amt int) {
     user := os.Getenv("PTS_USER")
     fmt.Println("Updating " + user)
     fmt.Println("Amount " + strconv.Itoa(amt))
}

func bihourUpdate() {
     rate := os.Getenv("PTS_BIHOUR_RATE")
     if rate == "" {
        rate = "0"
     }
     amt, err := strconv.Atoi(rate)
     if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
     }
     fmt.Println(amt)
     updateUser(amt)
}

func main() {
     fmt.Printf("Hello, world!\n")
     updateUser(1)
     bihourUpdate()
}
