/***********************************************************
 * File: fahr_cels2.go
 *
 * Purpose: prints a table of fahrenheit temperatures and 
 * their celsius equivalents using the formula c = 5/9(f-32)
 * This version prints a pretty output.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   for fahr := 0; fahr <= 300; fahr +=20 {
     celsius := 5 * (fahr - 32) / 9
     fmt.Printf("%3d%6d\n", fahr, celsius)
   }
 }
