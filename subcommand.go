package main

import (
"fmt"
//"net/http"
"flag"
"os"
)

func main() {
   seven := flag.NewFlagSet("w7", flag.ExitOnError)
   sevenfile := seven.String("w", "", "Set Wordlist")
   
   ten := flag.NewFlagSet("w10", flag.ExitOnError)
   tenfile := ten.String("w", "", "Set Wordlist")
   
   if len(os.Args) < 2 {
      fmt.Println("Expected 'w7' and 'w10' subcommand")
      os.Exit(1)
   }
   
   switch os.Args[1] {
   
   case "w7":
      seven.Parse(os.Args[2:])
      if *sevenfile == "" {
         fmt.Println("Error: Wordlist is not set")
      } else {
         fmt.Println("file:", *sevenfile)
      }
   
   case "w10":
      ten.Parse(os.Args[2:])
      if *tenfile == "" {
         fmt.Println("Error: Wordlist is not set")
      } else {
      fmt.Println("file:", *tenfile)
      }

   default:
      fmt.Println("Expected 'w7' and 'w10' subcommand")
      os.Exit(1)
   }
}
