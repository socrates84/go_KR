/**********************************************************
 * File: wc2.go
 *
 * Purpose: counts lines, words and characters. This is a
 * bare-bones version of the UNIX program wc. This version
 * uses a switch.
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
     r, _, err := reader.ReadRune()
     if err != nil {
       if err != io.EOF {
         // encountered a real error, exit
         fmt.Println("error:", err)
         return 
       }
       // reached end of file, exit
       break
     }
       
     switch r {
     case ' ':
       inWord = false
     case '\t':
       inWord = false
     case '\n': 
        if r == '\n' {
          nl++
        }
        inWord = false
     default:
        if inWord == false {
          inWord = true
          nw++
        }
     }
     
     if r != '\n' {
        nc++
     } 
  }

  fmt.Printf("nl = %d, nw = %d, nc = %d\n", nl, nw, nc)
}
