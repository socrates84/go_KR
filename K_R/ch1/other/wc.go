/**********************************************************
 * File: wc.c
 *
 * Purpose: counts lines, words and characters, This is a
 * bare-bones version of the UNIX program wc
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

   var (
     nl = 0
     nw = 0
     nc = 0
     inWord = false
   )

   reader := bufio.NewReader(os.Stdin)

   for {
     if r, _, err := reader.ReadRune(); err != nil {
       if err != io.EOF {
         // encountered a real real, exit
         fmt.Println("error:", err)
         return 
       }
       // reached end of file, exit
       break
       
     } else if r == ' ' || r == '\t' || r == '\n' {
        if r == '\n' {
          nl++
        }
        inWord = false
     } else if inWord == false {
        inWord = true
        nw++
     }
     nc++
   }

   fmt.Printf("nl = %d, nw = %d, nc = %d\n", nl, nw, nc)
 }
