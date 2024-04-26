mod n2;
mod n1;

fn main() {

    let idx = 1;

    let fns = [
        n1::exec,
        n2::exec
    ];

    fns[idx]();
}
