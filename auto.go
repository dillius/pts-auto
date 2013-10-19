package main

import (
       "fmt"
       "os"
       "net/http"
       "net/url"
       "io/ioutil"
)

const api_url = "http://www.peaktolerablestupidity.com/api/entry"

func getUser() {
     resp, _ := http.Get(api_url + "/" + os.Getenv("PTS_USER"))
     defer resp.Body.Close()
     contents, _ := ioutil.ReadAll(resp.Body)
     fmt.Printf("%s\n", string(contents))

}

func updateUser(amt string) {
     user := os.Getenv("PTS_USER")
     fmt.Println("Updating " + user)
     fmt.Println("Amount " + amt)
     values := make(url.Values)
     values.Set("user", user)
     values.Set("level", amt)

     resp, _ := http.PostForm(api_url, values)
     defer resp.Body.Close()
}

func autoUpdate() {
     rate := os.Getenv("PTS_AUTO_RATE")
     if rate == "" {
        rate = "0"
     }
     /*amt, err := strconv.Atoi(rate)
     if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
     }*/
     updateUser("+" + rate)
}

func main() {
     fmt.Printf("Hello, world!\n")
     getUser()
     updateUser("10")
     getUser()
     autoUpdate()
     getUser()
}
