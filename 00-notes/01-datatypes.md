# Datatypes in Go (by LLM)

## 🔹 Numeric Types

stick to *int, uint, float64, complex128* amd following default types for most general usage.

## Default Types
- bool
- string
- int
- uint32
- byte
- rune
- float64
- complex128

### 🟠 Integer Types

| Type     | Size     | Signed | Description                              |
|----------|----------|--------|------------------------------------------|
| `int`    | 32 or 64 | ✅     | Default integer type (platform-dependent) |
| `int8`   | 8-bit    | ✅     | -128 to 127                              |
| `int16`  | 16-bit   | ✅     | -32,768 to 32,767                        |
| `int32`  | 32-bit   | ✅     | -2,147,483,648 to 2,147,483,647          |
| `int64`  | 64-bit   | ✅     | -9.2×10¹⁸ to 9.2×10¹⁸                   |

| Type     | Size     | Signed | Description                              |
|----------|----------|--------|------------------------------------------|
| `uint`   | 32 or 64 | ❌     | Unsigned integer                         |
| `uint8`  | 8-bit    | ❌     | 0 to 255                                 |
| `uint16` | 16-bit   | ❌     | 0 to 65,535                              |
| `uint32` | 32-bit   | ❌     | 0 to 4.29×10⁹                           |
| `uint64` | 64-bit   | ❌     | 0 to 1.8×10¹⁹                           |
| `uintptr`| Arch-dep | ❌     | Stores pointer addresses as uint         |

---

### 🟠 Floating-Point Types

| Type      | Precision | Description               |
|-----------|-----------|---------------------------|
| `float32` | 32-bit    | IEEE-754 floating point   |
| `float64` | 64-bit    | Default float type        |

---

### 🟠 Complex Number Types

| Type        | Components       | Description                           |
|-------------|------------------|---------------------------------------|
| `complex64` | 2× float32       | Real + Imag part using float32        |
| `complex128`| 2× float64       | Real + Imag part using float64        |

---

## 🔹 Text and Character Types

| Type    | Underlying Type | Description                            |
|---------|------------------|----------------------------------------|
| `string`| -                | Immutable UTF-8 encoded text sequence  |
| `rune`  | `int32`          | Alias used to represent Unicode chars  |
| `byte`  | `uint8`          | Alias used for raw 8-bit binary data   |

> `byte` is a synonym for `uint8`  
> `rune` is a synonym for `int32` (typically a Unicode code point)

---

## 🔹 Boolean Type

| Type   | Description           |
|--------|-----------------------|
| `bool` | Only `true` or `false`|

---

## 🔹 Special / Utility Types

| Type     | Description                              |
|----------|------------------------------------------|
| `error`  | Built-in interface for error handling    |
| `any`    | Alias for `interface{}` (Go 1.18+)       |

---

## 📌 Summary of Aliases

| Alias  | Actual Type | Common Use                           |
|--------|-------------|--------------------------------------|
| `byte` | `uint8`     | Working with bytes, files, buffers   |
| `rune` | `int32`     | Unicode character representation     |
| `any`  | `interface{}` | Generic data (type-agnostic)       |

---

## Example Usages (Assignment)

```go
var b byte = 65         // same as var b uint8 = 65
var r rune = '世'       // same as var r int32 = Unicode code point
var i int = 42          // default integer
var f float64 = 3.1415  // floating point
var s string = "Hello"  // UTF-8 string
var ok bool = true      // boolean
var e error = nil       // error interface
var anything any = 5.5  // generic type holding a float
