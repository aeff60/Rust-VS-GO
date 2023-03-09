use std::io;

fn main() {
    println!("โปรแกรมคำนวนเกรด");

    let mut score = String::new();
    println!("กรุณากรอกคะแนนของคุณ: ");

    io::stdin()
        .read_line(&mut score)
        .expect("ไม่สามารถอ่านค่าได้");

    let score_float = match score.trim().parse::<f32>() {
        Ok(num) => num,
        Err(_) => {
            println!("กรุณากรอกตัวเลข");
            return;
        }
    };

    let grade = match score_float {
        x if x >= 80.0 => "A",
        x if x >= 75.0 => "B+",
        x if x >= 70.0 => "B",
        x if x >= 65.0 => "C+",
        x if x >= 60.0 => "C",
        x if x >= 55.0 => "D+",
        x if x >= 50.0 => "D",
        _ => "F",
    };

    println!("คะแนนของคุณคือ {:.2} และได้รับเกรด {}", score_float, grade);
}