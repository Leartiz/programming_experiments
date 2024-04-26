struct Foo {
    x: u64,
}

impl Foo {

    /// Любой тип, заимствующий экземпляр Foo, может
    /// вызвать этот метод, так как он требует всего лишь `ссылку` на Foo.
    pub fn total(&self) -> u64 {
        self.x
    }

    /// Только по `исключительным ссылкам` на экземпляры Foo
    /// можно вызывать этот метод, так как требуется, чтобы Foo был изменяемым
    pub fn increase(&mut self) {
        self.x += 1;
    }
}

pub fn exec() {
    let mut foo = Foo { x: 10 };
    println!("{}", foo.total());

    foo.increase();

    println!("{}", foo.total());
}