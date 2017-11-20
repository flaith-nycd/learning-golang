package main

import (
    "fmt"
    "strconv"
)

func FizzBuzz1() {
    for i := 1; i < 101; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

/* Python's way
fizzbuzz = []

for i in range(1, 101):
    val = ''
    if i % 3 == 0:
        val += 'fizz'
    if i % 5 == 0:
        val += 'buzz'
    if i % 3 != 0 and i % 5 != 0:
        val = i
    fizzbuzz.append(val)

print(fizzbuzz)
*/

func FizzBuzz2() {
    fizzbuzz := make([]string, 0)
    //fizzbuzz := []string{}

    for i := 1; i < 101; i++ {
        val := ""

        if i%3 == 0 {
            val += "Fizz"
        }
        if i%5 == 0 {
            val += "Buzz"
        }
        if i%3 != 0 && i%5 != 0 {
            val = strconv.Itoa(i)
        }
        fizzbuzz = append(fizzbuzz, val)
    }

    for _, value := range fizzbuzz {
        fmt.Println(value)
    }
}

func main() {
    fmt.Println("--- FizzBuzz v1 ---")
    FizzBuzz1()
    fmt.Println("--- FizzBuzz v2 ---")
    FizzBuzz2()
}
