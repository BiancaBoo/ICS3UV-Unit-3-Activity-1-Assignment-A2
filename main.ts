/**
 * Author: Bianca Bertinato
 * Version: 1.0.0
 * Date: 2025-11-08
 * This program calculates and displays the length, width, and height of a cube given its volume.
 */

// INITIALIZE 
const VOLUME: number = 1000;  
let side: number;            

// INPUT - none 

// PROCESS
// length = width = height = cube root of volume
side = Math.cbrt(VOLUME);

// OUTPUT
// display the result
console.log("The length and width and height of a cube with a volume of " + VOLUME + " mm³ is " + side + " mm.");

console.log("\nDone.");