/***********************************************************
 * File: char_count.go
 *
 * Purpose: counts and prints the number of characters in 
 * the input.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import (
   "io"
   "os"
   "bufio"
   "fmt"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)

   nc := 0
   for {
     if _, _, err := reader.ReadRune();  err != nil {
       if err != io.EOF {   // only print valid errors 
          fmt.Println("error:", err)
          return
       }
       break                // reached end of file
     }
     nc++
   }

   fmt.Printf("\nNumber of characters read: %d\n", nc)
 }
