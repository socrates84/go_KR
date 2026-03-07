/**********************************************************
 * File: ex1.12.go
 *
 * Purpose: prints its input one word per line.
 *
 * Author: socrates
 *
 *********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "os"
   "strings"
 )

 func main() {

   scanner := bufio.NewScanner(os.Stdin)

   for scanner.Scan() {
      line := scanner.Text()
      for _, word := range strings.Fields(line) {
        fmt.Println(word)
      }
   }

   // verify errors
   if err := scanner.Err(); err != nil {
     fmt.Fprintln(os.Stderr, "error reading input:", err)
     os.Exit(1)
   }
 }
