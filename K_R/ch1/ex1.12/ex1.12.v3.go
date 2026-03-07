/**********************************************************
 * File: ex1.12.v3.go
 *
 * Purpose: prints its input one word per line. This version
 * increases the default bufio.Scanner max token size from
 * 64 KB to 1MB to take care of very long input.
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
   scanner.Buffer(make([]byte, 1024), 1024*1024)

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
