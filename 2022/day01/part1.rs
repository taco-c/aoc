use std::io;
use std::cmp;

fn main() {
    let mut max = 0;
    let mut elf = 0;

    for line in io::stdin().lines() {
        match line.unwrap().parse::<u32>() {
            Ok(n) => {
                elf += n
            },
            Err(_) => {
                max = cmp::max(max, elf);
                elf = 0;
            },
        };
    }

    max = cmp::max(max, elf);
    println!("{max}");
}

