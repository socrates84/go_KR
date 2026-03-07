/**********************************************************
 * File: ex1.12.v2.go
 *
 * Purpose: prints its input one word per line. This version
 * scans one word at at ime, instead of a line at a time.
 *
 * Author: socrates
 *
 *********************************************************/
 package main

 import (
   "bufio"
   "fmt"
   "os"
 )

 func main() {

   scanner := bufio.NewScanner(os.Stdin)

   // change scanner to read words instead of lines
   scanner.Split(bufio.ScanWords)

   for scanner.Scan() {
      word := scanner.Text()
      fmt.Println(word)
   }

   // verify errors
   if err := scanner.Err(); err != nil {
     fmt.Fprintln(os.Stderr, "error reading input:", err)
     os.Exit(1)
   }
 }
