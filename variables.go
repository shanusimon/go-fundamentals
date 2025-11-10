package main

import "fmt"


func variables() {
    var name string = "Shanu" // Basic var declaration. You can also omit the type if you assign a value.
    var age int = 23          // If a type is declared but no value is given, Go assigns the zero value (for int → 0).
    
    city := "Kochi"           // Short variable declaration (syntactic sugar for var). Can only be used *inside* functions.
    version := 1.23           // Type inferred as float64.
    
    x, y, z := 10, 20, 30     // Multiple variable declaration.
    
    const country = "India"   // Constant declaration. Cannot be reassigned.
    
    age = 24                  // Variables (declared with var or :=) can be updated.

    // Printing all values
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(city)
    fmt.Println(version)
    fmt.Println(x, y, z)
    fmt.Println(country)
}
