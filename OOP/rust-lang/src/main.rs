trait Animal {
    fn speak(&self);
}

struct Dog {
    name: String,
}

impl Animal for Dog {
    fn speak(&self) {
        println!("{} โฮ่งโฮ่ง", self.name);
    }
}

struct Cat {
    name: String,
}

impl Animal for Cat {
    fn speak(&self) {
        println!("{} เมี๊ยวเมี๊ยว", self.name);
    }
}

fn main() {
    let dog = Dog {
        name: String::from("Rufus"),
    };
    let cat = Cat {
        name: String::from("Garfield"),
    };

    let animals: Vec<Box<dyn Animal>> = vec![Box::new(dog), Box::new(cat)];

    for animal in animals {
        animal.speak();
    }
}