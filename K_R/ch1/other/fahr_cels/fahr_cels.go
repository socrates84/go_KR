/***********************************************************
 * File: fahr_cels.go
 *
 * Purpose: prints a table of fahrenheit temperatures and 
 * their celsius equivalents using the formula 
 * c = 5/9(f-32).
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   for fahr := 0; fahr <= 300; fahr +=20 {
     celsius := 5 * (fahr - 32) / 9
     fmt.Println(fahr, celsius)
   }
 }
