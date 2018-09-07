package main

import (
    "fmt"
    "log"

    "github.com/mileusna/crontab"
)

func main() {

    ctab := crontab.New() // create cron table

    // AddJob and test the errors
    err := ctab.AddJob("0 12 1 * *", myFunc) // on 1st day of month
    if err != nil {
        log.Println(err)
        return
    }

    // MustAddJob is like AddJob but panics on wrong syntax or problems with func/args
    // This aproach is similar to regexp.Compile and regexp.MustCompile from go's standard library,  used for easier initialization on startup
    ctab.MustAddJob("* * * * *", myFunc) // every minute
    ctab.MustAddJob("0 12 * * *", myFunc3) // noon lauch

    // fn with args
    ctab.MustAddJob("0 0 * * 1,2", myFunc2, "Monday and Tuesday midnight", 123)
    ctab.MustAddJob("*/5 * * * *", myFunc2, "every five min", 0)

    // all your other app code as usual, or put sleep timer for demo
    // time.Sleep(10 * time.Minute)
}

func myFunc() {
  fmt.Println("Helo, world every second")
}

func myFunc3() {
  fmt.Println("Noon!")
}

func myFunc2(s string, n int) {
  fmt.Println("We have params here, string", s, "and number", n)
}
