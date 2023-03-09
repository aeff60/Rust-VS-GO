use std::io;

fn main() {
    println!("Welcome to the Rust Calculator!");

    let mut input = String::new();
    println!("Please enter the first number:");
    io::stdin().read_line(&mut input).unwrap();
    let num1: f64 = input.trim().parse().unwrap();

    input.clear();
    println!("Please enter the second number:");
    io::stdin().read_line(&mut input).unwrap();
    let num2: f64 = input.trim().parse().unwrap();

    println!("Please choose an operation (+, -, *, /):");
    input.clear();
    io::stdin().read_line(&mut input).unwrap();
    let op = input.trim();

    let result = match op {
        "+" => num1 + num2,
        "-" => num1 - num2,
        "*" => num1 * num2,
        "/" => num1 / num2,
        _ => panic!("Invalid operator"),
    };

    println!("Result: {}", result);
}