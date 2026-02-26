/***********************************************************
 * File: ex1.10.go
 *
 * Purpose: copies its input to its output, replacing each
 * tab by \t, each backspace by \b, and each backlash by \\.
 * This makes tabs and backslashes visible in an unambiguous
 * way
 *
 * Author: Alain
 *
 **********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "io"
   "os"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)

   for {
     r, _, err := reader.ReadRune()
     if err != nil {
       if err != io.EOF {     // error, print it and exit
         fmt.Println("error:", err)
         return
       }
       break        // reached end of file, exit
     }

     switch r {
     case '\t':
        fmt.Print("\\t")
     case '\\':
       fmt.Print("\\\\")
     default:
       fmt.Printf("%c", r)
    }
   }
 }

