/***********************************************************
 * File: ex1.9.go
 *
 * Purpose: copies its input to its output, replacing each 
 * string of one or more blanks by a single blank.
 *
 * Author: socrates
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

   blankSeen := false
   for {
     if r, _, err := reader.ReadRune(); err != nil {
       if err != io.EOF { 
         fmt.Println("error:", err)
         return
       }
       break                            // reached end of file
     } else if r == rune(' ') || r == rune('\t') {
         blankSeen = true
     } else {
         if blankSeen {
           fmt.Printf("%c", ' ')        // print previous blank
           blankSeen = false
         }
         fmt.Printf("%c", r)
     } 
   }
 }
