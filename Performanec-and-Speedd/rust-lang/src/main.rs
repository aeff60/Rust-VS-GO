use std::time::Instant;

fn main() {
    let start = Instant::now();

    let sum = (1..=10000).sum::<u32>();

    let duration = start.elapsed();
    println!("Sum: {}", sum);
    println!("Time elapsed: {:?}", duration);
}