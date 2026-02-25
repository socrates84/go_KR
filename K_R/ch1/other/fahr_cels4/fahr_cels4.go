/***********************************************************
 * File: fahr_cels4.go
 *
 * Purpose: prints the fahrenheit to celsius conversion 
 * table. This version uses Go's untyped consts.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   const (
     lower = 0
     upper = 300
     step = 20
   )

   fmt.Println("Fahrenheit celsius")
   fmt.Println("==================")

   for fahr := lower; fahr <= upper; fahr += step {
     celsius := (5.0/9.0) * (float64(fahr) - 32) 
     fmt.Printf("%6d %9.1f\n", fahr, celsius)
   }
 }
