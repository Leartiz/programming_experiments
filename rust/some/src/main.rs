mod n2;
mod n1;
mod n3;

fn main() {

    let idx = 2;

    let fns = [
        n1::exec,
        n2::exec,
        n3::exec
    ];

    fns[idx]();
}
