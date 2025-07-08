# 🤝 Contributing to `salvation`

**TL;DR**: Do stuff. Break things (just not in `main`). Gaslight the tests until they agree with you. It's _open source_, not a cathedral.

First of all: why?

But seriously, I could say things like "Hey, we appreciate the help and consideration!" but let's not pretend that's what this is. If you've found yourself here, considering contributing to a library that brings `Option<T>` to **Golang**, welcome. I won't ask what went wrong in your life to get here.

You're allowed to break stuff, just not in `main`. That's where the lies live.

## 🧘‍♂️ Ground Rules

- Yes, it's a _meh_ lib, but let's try to keep it from collapsing in on itself.
- All contributions must maintain the illusion of seriousness.
- Write idiomatic Go. Unless idiomatic Go gets in the way of the bit.
- PRs with actual utility will be merged. As long as they don't look like a fever dream.
- All code should be unit tested. All tests should pass. You're allowed to change the tests until they agree with you. That's how software works.
- Try not to break the already existing API.

## 📦 Feature Suggestions

Before suggesting a feature, ask yourself:

- Why?
- Does this increase the emotional value of the API?
- Will someone depend on this in production and live to regret it?
- Does your spec match the current API?

If you somehow answered _"yes"_ to all of these (even the "Why?"): submit the issue/PR/whatever.

## 🧪 Testing

Run `go test ./...` and pretend you care.

## 🌀 Before You Open a PR

1. Format your code with `go fmt`. I need to be able to read it without crying.
2. Document your functions like they'll be read by a team of ten (they won't).
3. Ask yourself if you're okay being associated with a project named `salvation`. Forever. It will come back to haunt you (and your GitHub).

## ✨ Bonus Point

- Add APIs that read like a philosophy: `NewPossibility.MustReveal`
- Submit serious-sounding PRs with names like `"Implement QuantumFold logic"` or `"Refactor soul-crushing ambiguity module"`

## 🙏 Final Note

Thanks, btw. It's weird that you're here. But thanks.
