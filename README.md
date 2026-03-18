# Vld

Go validation library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/vld.svg)](https://pkg.go.dev/github.com/pierrre/vld)

## Features

- Simple value validation with declarative syntax and readable error messages.
- Type-safe (using generics), no reflection, no code generation, no allocation while validating, and no dependency.
- Composable validation rules with `And`, `Or`, `All`, `Not`, `If`, and `IfElse`.
- Built-in validators for common checks such as equality, ordering, ranges, membership, zero values, and required values.
- Specialized validators for strings, slices, maps, pointers, iterators, and regular expressions, including length, emptiness, containment, uniqueness, and per-item validation.
- Localized error messages (en/fr) with support for custom localizations.
- Extract path from errors.


## Usage

[Example](https://pkg.go.dev/github.com/pierrre/vld#example-package)
