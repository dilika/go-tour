# The Go closure trap every developer should know

```go
var fns []func()
for i := 0; i < 3; i++ {
    fns = append(fns, func() { fmt.Println(i) })
}
for _, f := range fns {
    f()
}
```

What does this print?

Most experienced Go developers answer without hesitation: **`3 3 3`**. That was the correct answer for over a decade. It's also, today, wrong.

## The Go version decides

On Go 1.21 and earlier, this prints:

```
3
3
3
```

On Go 1.22 and later, it prints:

```
0
1
2
```

Same code. Different behavior. And it's not a bug — it's a deliberate change to the language specification.

## Why `3 3 3` for ten years

For Go's entire first decade, the loop variable `i` was **a single variable**, declared once and reused each iteration. Every closure created inside the loop captured **a reference to that same variable**.

When the loop ends, `i` equals `3`. The closures only run afterward, in the second `for`. At that point they all read the same `i` — which is `3`. Hence `3 3 3`.

This trap was so classic it appeared in nearly every Go onboarding guide. The universal workaround was explicit:

```go
for i := 0; i < 3; i++ {
    i := i // local copy, a fresh variable per iteration
    fns = append(fns, func() { fmt.Println(i) })
}
// → 0 1 2
```

`i := i` — the strangest line in the language, repeated by generations of developers to work around a questionable specification.

## What Go 1.22 changed

In 2024, Go 1.22 changed `for` loop semantics: **each iteration creates a new variable**. The `i := i` copy is no longer needed. The closure now captures a distinct `i` on each pass, retaining the value at the moment of that iteration.

`0 1 2` becomes the default — and the intuitive behavior most beginners had always expected.

This is one of the rare cases where Go broke semantic compatibility in the name of correctness. Migrations aren't trivial for codebases that depended, intentionally or not, on the old behavior.

## How I discovered I didn't know

I predicted `1 2 3`. Neither the old trap (`3 3 3`) nor the new behavior (`0 1 2`).

I had a partial grasp: I knew closures capture, I knew there was a version story, but I had never checked. My mental model was a fuzzy superposition of two behaviors, and I had never run the code.

The method I now apply to learn Go: **predict before reading, then run, then confront**.

1. **Predict**: write what I think the code does, without looking at the answer.
2. **Run**: execute the real code, observe the output.
3. **Confront**: compare prediction and reality, note the gap precisely.

Step 1 matters most. An exposed prediction can no longer hide. It can't slide into "oh, I knew that" — it's written, dated, wrong, and undeniable.

`1 2 3` in black and white forced me to admit I didn't understand, where `3 3 3` would have let me believe I'd mastered the topic. A precise wrong prediction beats a vague correct intuition.

## The larger lesson

The loop-closure trap isn't anecdotal. It illustrates three things that define Go mastery:

1. **Capture is by reference, not by value.** A closure doesn't copy the variable; it references its location. Until you know where that variable lives — stack or heap, shared or per-iteration — you're not predicting, you're guessing.

2. **The specification evolves.** Go 1.22 changed a semantic in place. The code you write today doesn't behave like the code you read in a 2019 tutorial. Rereading the spec is a senior reflex, not a beginner's admission.

3. **An untested intuition is not knowledge.** The gap between "I think I know" and "I know because I ran it" is exactly where production bugs live.

## Go further

- The exact change: [Go 1.22 release notes — loop variable scoping](https://go.dev/doc/go1.22)
- The original proposal: [issue #60078](https://github.com/golang/go/issues/60078)
- Verify it yourself: copy the snippet at the top, run `go run` on Go 1.21 then Go 1.22. Four lines of code, two decades of history.

---

Predict before you execute. The first time your prediction is wrong is the first time you actually start learning.
