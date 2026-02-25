/***********************************************************
 * File: ex1.8.go
 *
 * Purpose: counts blanks, tabs, and newlines.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import (
   "bufio"
   "os"
   "fmt"
   "io"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)

   var (
     bl = 0
     tb = 0
     nl = 0
   )

   for {
     if r, _, err := reader.ReadRune(); err != nil {
       if err != io.EOF { 
         fmt.Println("error:", err)
         return
       }
       break          // reached end of file
     } else if r == rune(' ') {
            bl++
     } else if r == rune('\t') {
            tb++
     } else if r == rune('\n') {
            nl++
     }
   }

   fmt.Printf("\nbl = %d, tb = %d, nl = %d\n", bl, tb, nl)
 }


