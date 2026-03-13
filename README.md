# Vld

Go validation library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/vld.svg)](https://pkg.go.dev/github.com/pierrre/vld)

## Features

- Declarative syntax, type-safe (using generics), no reflection, and no code generation.
- Composable validation rules with `And`, `Or`, `All`, `Not`, `If`, and `IfElse`.
- Built-in validators for common checks such as equality, ordering, ranges, membership, zero values, and required values.
- Specialized validators for strings, slices, maps, pointers, and regular expressions, including length, emptiness, containment, uniqueness, and per-item validation.
- Helpers for validating derived and nested values with `Get`, `Field`, `Wrap`, and `Message`.
- Error helpers to wrap and join validation errors while keeping messages structured and readable.

## Usage

[Example](https://pkg.go.dev/github.com/pierrre/vld#example-package)
