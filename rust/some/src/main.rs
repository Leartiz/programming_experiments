mod n2;
mod n1;
mod n3;
mod n4;

fn main() {

    let idx = 0;

    let fns = [
        n4::exec,
    ];

    fns[idx]();
}
