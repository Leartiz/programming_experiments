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
    let num: i32 = rng.gen::<i32>() % 3;

    let number = if num == 1 {
        4
    } else if num == 2 {
        5
    } else {
        10
    };

    println!("number = {}", number);
}

fn example_4() {
    let mut rng = rand::thread_rng();
    let num = rng.gen::<i32>() % 4;

    match num
    {
        1=>println! ("один"),
        2=>println! ("два"),
        3=>println! ("три"),
        _=>println! ("непонятно")
    }
}

fn as_str(data: &u32) -> &str {
    let s = format!("{}", data);
    &s
}

fn example_5() {
    let val :u32 = 55;
    format!("{}", as_str(&val));
}

pub fn exec() {
    example_0();
    example_1();
    example_2();
    example_4();
    example_5();
    //...
}
