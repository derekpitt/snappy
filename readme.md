# Snappy API Client in Go

    package main

    import (
      "fmt"
      "github.com/derekpitt/snappy"
    )

    func main() {
      client := snappy.WithAPIKey("<your api key here>")
      ticket, err := client.Ticket(12345)

      if err != nil {
        return
      }

      fmt.Println(ticket.Summary)
    }

# Documentation

[http://godoc.org/github.com/derekpitt/snappy](http://godoc.org/github.com/derekpitt/snappy)

## TODO

  - Uploading documents
  - Some more deep comparing in tests

## Random Thoughts

I have the actions that are account specific to also take in an ```accountId``` parameter rather than constructing an ```Account{}``` and/or having reciever functions on ```Account{}```, but I may change that (if it feels good).
