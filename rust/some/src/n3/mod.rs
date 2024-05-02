use std::mem::size_of;
use std::sync::Arc;
use rand::Rng;

fn example_0() {
    let mut age = 36;
    println!("Начальное значение: {}", age);

    age = 25;
    println!("Конечное значение: {}", age);
}

/* shadowing */
fn example_1() {
    let mut number = 10;
    println!("number = {}", number);

    let number = 15.8;
    println!("number = {}", number);

    let number = "hello";
    println!("number = {}", number);
}

fn example_2() {
    let ch = 'a';
    println!("char: {}", ch);
    println!("size_of<char>: {}", size_of::<char>());
}

fn example_3() {
    let mut rng = rand::thread_rng();
    let cond: i32 = rng.gen::<i32>() % 3;

    let number = if cond == 0 {
        4
    } else if cond == 1 {
        5
    } else {
        10
    };

    println!("number = {}", number);
}

pub fn exec() {
    example_0();
    example_1();
    example_2();
    //...
}
