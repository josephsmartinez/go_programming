package main
import (
  "fmt"
  "os"
  //"flag"
)

func main() {

  option := os.Args[1:]

  // Verify that a subcommand has been provided
  // os.Arg[0] is the main command
  // os.Arg[1] will be the subcommand
  if len(os.Args) < 2 {
      fmt.Println("subcommand is required")
      os.Exit(1)
  }

  fmt.Println(option)
  fmt.Println(len(os.Args))

  switch os.Args[1] {
  case "cpu":
    // get os cpu information
    
  case "ram":
    // get ram from system

  }

}
