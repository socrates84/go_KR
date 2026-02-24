/***********************************************************
 * File: file_copy.go
 *
 * Purpose: prints its input to its output. Go does not have
 * a built-in function getchar() and instead characters are
 * represented using Unicode. You have two options:
 * read runes or raw bytes. I have chosen the former as it
 * makes reading input more efficient and convenient
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "os"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)

   for {
     r, _, err := reader.ReadRune()
     if err != nil {
       break
     }
     fmt.Printf("%c", r)
   }
 }
