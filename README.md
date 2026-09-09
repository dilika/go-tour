# Go Tour Practice

This repository contains small Go programs and exercises used to practice language fundamentals and software-engineering discipline. I work through the examples hands-on to strengthen my own understanding and judgment.

## Structure

Exercises are grouped into topic-based directories such as arrays, constants, functions, pointers, slices, structs, and variables. Some directories contain independent executable programs.

## Validation

Run the repository-wide checks from the project root:

```bash
  test -z "$(gofmt -l .)"
  go vet ./...
  go test ./...
  go test -race ./...

```

## Current status

The repository is being repaired on the chore/green-baseline branch. Some repository-wide checks are currently expected to fail while the known build, formatting, and test gaps are addressed.

Progress and acceptance criteria are tracked in issue #1.
