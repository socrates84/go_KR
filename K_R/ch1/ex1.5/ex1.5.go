/***********************************************************
 * File: ex1.5.go
 *
 * Purpose: prints the fahrenheit to celsius conversion 
 * table in reverse order.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   fmt.Println("Fahrenheit celsius")
   fmt.Println("==================")

   for fahr := 300; fahr >= 0; fahr -= 20 {
     celsius := (5.0/9.0) * (float64(fahr) - 32) 
     fmt.Printf("%6d %9.1f\n", fahr, celsius)
   }
 }
