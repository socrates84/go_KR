/**********************************************************
 * File: slices3.go
 *
 * Purpose: counts the number of occurrences of each digit,
 * of white space characters (blanks, tabs, newlines) and 
 * of all other characters.
 *
 * This version is a high performance version
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

   buf := make([]byte, 32*1024)

   for {
     n, err := reader.Read(buf)
     
     if err == io.EOF {
       break
     }

     if err != nil {
       fmt.Fprintln(os.Stderr, "error reading input:", err)
       os.Exit(1)
     }

     for _, b := range buf[:n] {
       switch {
       case b >= '0' && b <= '9':
         digits[b - '0']++
       case b == ' ' || b == '\t' || b == '\n':
         whitespace++
       default:
         other++
       }
     }
   }

   return digits, whitespace, other
 }
