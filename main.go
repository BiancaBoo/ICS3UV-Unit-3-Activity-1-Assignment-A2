/*
* Author: Bianca Bertinato
* Version: 1.0.0
* Date: 2025-11-08
* This program calculates and displays the length, width, and height of a cube given its volume.
*/

package main

import (
    "fmt"
    "math"
)

func main() {
    // INITIALIZE 
    const VOLUME float64 = 1000   
    var side float64              

    // INPUT - none 

    // PROCESS
    // For a cube, length = width = height = cube root of volume
    side = math.Cbrt(VOLUME)

    // OUTPUT
    // display the result
    fmt.Println("The length and width and height of a cube with a volume of", VOLUME, "mm³ is", side, "mm.")

    fmt.Println("\nDone.")
}