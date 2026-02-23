/***********************************************************
 * File: ex1.3.go
 *
 * Purpose: prints a table of fahrenheit temperatures and 
 * their celsius equivalents using the formula c = 5/9(f-32)
 * This version prints a pretty output and used floating 
 * point numbers. Also adds a heading to the table.
 *
 * Author: socrates
 *
 **********************************************************/
 package main

 import "fmt"

 func main() {

   fmt.Println("Fahrenheit  Celsius")
   fmt.Println("===================")

   for fahr := 0; fahr <= 300; fahr += 20 {
     celsius := (5.0/9.0) * (float64(fahr) - 32) 
     fmt.Printf("%6d %10.1f\n", fahr, celsius)
   }
 }
