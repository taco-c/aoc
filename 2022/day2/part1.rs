use std::io;

// Rock     A X 1
// Paper    B Y 2
// Scissors C Z 3

fn main() {
    let mut score = 0;

    for line in io::stdin().lines() {
        let s = line.unwrap();
        let a = s.split(' ').collect::<Vec<&str>>();

        score += match a[1] {
            "X" => 1,
            "Y" => 2,
            "Z" => 3,
            _ => 0,
        };

        score += if a[0] == "A" && a[1] == "X" { 3 }
            else if a[0] == "A" && a[1] == "Y" { 6 }
            else if a[0] == "A" && a[1] == "Z" { 0 }
            else if a[0] == "B" && a[1] == "X" { 0 }
            else if a[0] == "B" && a[1] == "Y" { 3 }
            else if a[0] == "B" && a[1] == "Z" { 6 }
            else if a[0] == "C" && a[1] == "X" { 6 }
            else if a[0] == "C" && a[1] == "Y" { 0 }
            else if a[0] == "C" && a[1] == "Z" { 3 }
            else { 0 };
    }

    println!("{score}");
}

