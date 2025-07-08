# 🌀 Salvation

`salvation` is a package containing a generic Go wrapper for optional values.

Inspired by Rust's `Option<T>` and Haskell's `Maybe`, `Possibly[T]` gives you a way to wrap things that may or may not exist, and then reflect on their Nothingness.

## ✨ Features

- Generic support for any type `T`
- Reflection-powered nil checking
- Precomputable "Nothingness" for less overhead
- Configurable rules (empty slices can be something... if you _believe_)
- Matcher flow control
- Fully documented (yes, this is a feature)

## 🧠 Why?

Honestly? I don't know. But it seemed fun enough to build.

## 🚀 Usage

```go
opt := salvation.NewPossibility[*MyStruct](nil)

if opt.IsSomething() {
    // yay!
}

if opt.IsNothing() {
    // aww!
}
```

```go
opt.Match().
    Case(func(v int) bool { return v > 10 }, func(v int) { fmt.Println("Large number", v) }).
    Default(func(_ Possibly[int]) {
        fmt.Println("Nothing matched. Or maybe there was Nothing at all.")
    })
```

## 🧯 License

This project uses the **The Unlicense** license (lol, jk just copy and paste into your code and believe in yourself)
