/***********************************************************
 * File: line_count.go 
 *
 * Purpose: counts the number of lines in the input.
 *
 * Author: socrates
 *
 ***********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "io"
   "os"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)

   nl := 0
   for {
     if r, _, err := reader.ReadRune(); err != nil {
       if err != io.EOF { // encountered an error, exit
          fmt.Println("error:", err)
          return
        }
        break             // reached of file
     } else if r == rune('\n') {
          nl++
     }
   }

   fmt.Printf("\nNumber of lines: %d\n", nl)
 }



