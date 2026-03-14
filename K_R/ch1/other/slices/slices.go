/**********************************************************
 * File: slices.go
 *
 * Purpose: counts the number of occurrences of each digit,
 * of white space characters (blanks, tabs, newlines) and 
 * of all other characters.
 *
 * Author: socrates
 *
 *********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "io"
   "os"
 )

 func main() {

   reader := bufio.NewReader(os.Stdin)
   digits := make([]int, 10)
   nwhite := 0
   nother := 0

   for {
     r, _, err := reader.ReadRune()
     if err == io.EOF {
       break
     }

     if err != nil {
       fmt.Fprintln(os.Stderr, "error reading input:", err)
       os.Exit(1)
     }

     if r >= '0' && r <= '9' {
       digits[r - '0']++
     } else if r == ' ' || r == '\t' || r == '\n' {
       nwhite++
     } else {
       nother++
     }
   }

   fmt.Println("digits =", digits)
   fmt.Println("white spaces =", nwhite)
   fmt.Println("other characters =", nother)
 }

