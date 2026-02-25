/***********************************************************
 * File: fahr_cels3.go
 *
 * Purpose: prints a table of fahrenheit temperatures and 
 * their celsius equivalents using the formula c = 5/9(f-32)
 * This version prints a pretty output and converted the
 * fahrenheit to a float before printing.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   for fahr := 0; fahr <= 300; fahr += 20 {
     celsius := (5.0/9.0) * (float64(fahr) - 32) 
     fmt.Printf("%3d %6.1f\n", fahr, celsius)
   }
 }
