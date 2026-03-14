/**********************************************************
 * File: slices2.go
 *
 * Purpose: counts the number of occurrences of each digit,
 * of white space characters (blanks, tabs, newlines) and 
 * of all other characters.
 *
 * Author: socrates
 *
 ********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "io"
   "os"
   "unicode"
 )

 func main() {
   digits, whitespace, other := count(os.Stdin)

   fmt.Println("Digit counts:")
   for i, v := range digits {
     fmt.Printf("%d: %d\n", i, v)
   }

   fmt.Println("Whitespace:", whitespace)
   fmt.Println("Other:", other)
 }

func count(r io.Reader)([10]int, int, int) {

   reader := bufio.NewReader(r)
   
   var digits [10]int
   var whitespace int 
   var other int

   for {
     ch, _, err := reader.ReadRune()
     
     if err == io.EOF {
       break
     }

     if err != nil {
       fmt.Fprintln(os.Stderr, "error reading input:", err)
       os.Exit(1)
     }

     switch {
     case ch >= '0' && ch <= '9':
       digits[ch - '0']++
     case unicode.IsSpace(ch):
       whitespace++
     default:
       other++
     }
   }

   return digits, whitespace, other
 }
