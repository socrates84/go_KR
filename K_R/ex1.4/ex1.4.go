/***********************************************************
 * File: ex1.4.go
 *
 * Purpose: prints the corresponding celsius to fahrenheit
 * conversion table.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   fmt.Println("Celsius  Fahrenheit")
   fmt.Println("===================")

   for celsius := 0; celsius <= 300; celsius += 20 {
     fahr := (9.0/5.0) * float64(celsius) + 32 
     fmt.Printf("%4d %11.1f\n", celsius, fahr)
   }
 }
