use std::io;

// Rock     A 1
// Paper    B 2
// Scissors C 3

// Lose     X 0
// Draw     Y 3
// Win      Z 6

const ROCK:    u8 = 1;
const PAPER:   u8 = 2;
const SCISSOR: u8 = 3;

const WIN:  u8 = 6;
const DRAW: u8 = 3;
const LOSE: u8 = 0;

fn main() {
    let mut score = 0;

    for line in io::stdin().lines() {
        let s = line.unwrap();
        let a = s.split(' ').collect::<Vec<&str>>();

        score += if a[0] == "A" /* ROCK    */ && a[1] == "X" { LOSE + SCISSOR }
            else if a[0] == "A" /* ROCK    */ && a[1] == "Y" { DRAW + ROCK }
            else if a[0] == "A" /* ROCK    */ && a[1] == "Z" { WIN  + PAPER }
            else if a[0] == "B" /* PAPER   */ && a[1] == "X" { LOSE + ROCK }
            else if a[0] == "B" /* PAPER   */ && a[1] == "Y" { DRAW + PAPER }
            else if a[0] == "B" /* PAPER   */ && a[1] == "Z" { WIN  + SCISSOR }
            else if a[0] == "C" /* SCISSOR */ && a[1] == "X" { LOSE + PAPER }
            else if a[0] == "C" /* SCISSOR */ && a[1] == "Y" { DRAW + SCISSOR }
            else if a[0] == "C" /* SCISSOR */ && a[1] == "Z" { WIN  + ROCK }
            else { 0 };

    }

    println!("{score}");
}

