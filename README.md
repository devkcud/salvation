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

Usage examples below. A full Wiki is coming soon (probably just so I can pretend it matters).

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

## 🫠 Contributing

Thinking of contributing? That's adorable.

Before submitting a PR, feature request, or unhinged philosophical improvement, please consult the [`CONTRIBUTING.md`](CONTRIBUTING.md). It's full of helpful guidelines, emotional red flags, and vague threats.

If you're still interested after reading it, I can't stop you. No one can.

## ⭐️ Validation

If this repo made you feel something (anything at all) you can click the star. It won't fix your life, but it'll make this repo slightly more visible to other souls.

> I made this section against my own will...

## 🧯 License

This project uses **The Unlicense** license (just copy and paste into your code and believe in yourself)
