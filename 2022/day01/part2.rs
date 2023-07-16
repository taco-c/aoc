use std::io;

fn main() {
    let mut max = [0; 3];
    let mut elf = 0;

    for line in io::stdin().lines() {
        match line.unwrap().trim().parse::<u32>() {
            Ok(n) => {
                elf += n
            },
            Err(_) => {
                if elf > max[0] {
                    max[2] = max[1];
                    max[1] = max[0];
                    max[0] = elf;
                } else if elf > max[1] {
                    max[2] = max[1];
                    max[1] = elf;
                } else if elf > max[2] {
                    max[2] = elf;
                }
                elf = 0;
            },
        };
    }

    if elf > max[0] {
        max[2] = max[1];
        max[1] = max[0];
        max[0] = elf;
    } else if elf > max[1] {
        max[2] = max[1];
        max[1] = elf;
    } else if elf > max[2] {
        max[2] = elf;
    }

    println!("{}", max.iter().sum::<u32>());
}

