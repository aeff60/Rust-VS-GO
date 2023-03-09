use std::io;

fn main() {
    println!("โปรแกรมคำนวน Factorial");

    let mut n = String::new();

    println!("กรุณากรอกจำนวนเต็มบวก: ");

    io::stdin()
        .read_line(&mut n)
        .expect("ไม่สามารถอ่านค่าจากแป้นพิมพ์ได้");

    let n: u64 = n.trim().parse().expect("กรุณากรอกตัวเลข");

    let mut result = 1;

    for i in 1..=n {
        result *= i;
    }

    println!("{}! = {}", n, result);
}