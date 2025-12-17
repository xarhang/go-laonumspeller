# go-laonumspeller

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

ແປງຕົວເລກເປັນພາສາລາວ ແລະ ແປງຄຳພາສາລາວກັບຄືນເປັນຕົວເລກ

Convert numbers to Lao words and vice versa.

## ✨ ຄຸນສົມບັດ / Features

✅ **ຮອງຮັບຕົວເລກລົບ** / Support negative numbers  
✅ **ຮອງຮັບທົດສະນິຍົມ** / Support decimal numbers  
✅ **ຮອງຮັບຕົວເລກຂະໜາດໃຫຍ່** / Support large numbers:
  - 10³: ພັນ (Thousand) → ສິບພັນ, ຊາວພັນ, ສາມສິບພັນ...
  - 10⁵: ແສນ (Hundred Thousand)
  - 10⁶: ລ້ານ (Million) → ສິບລ້ານ, ຊາວລ້ານ, ຮ້ອຍລ້ານ...
  - 10⁹: ຕື້ (Billion) → ສິບຕື້, ຊາວຕື້, ຮ້ອຍຕື້...
  - 10¹²: ລ້ານລ້ານ (Trillion)

✅ **ໃຊ້ກົດລະບຽບພາສາລາວທີ່ຖືກຕ້ອງ** / Uses correct Lao language rules:
  - **ເອັດ** (for ones place after tens: ສິບເອັດ, ຊາວເອັດ)
  - **ຊາວ** (for twenty: ຊາວ, ຊາວເອັດ)
  - **ຈຸດ** (decimal point separator)
  - **ແສນ** (100,000 - not ໜຶ່ງຮ້ອຍພັນ)

## 📦 ຕິດຕັ້ງ / Installation

### CLI Tool
```bash
go install github.com/xarhang/go-laonumspeller/cmd/laonumspell@latest
```

### Library
```bash
go get github.com/xarhang/go-laonumspeller
```

## 🚀 ການໃຊ້ງານ CLI / CLI Usage

### ແປງຕົວເລກເປັນຄຳ / Convert Number to Words
```bash
laonumspell 123
# Output: ໜຶ່ງຮ້ອຍຊາວສາມ

laonumspell 123.45
# Output: ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ

laonumspell -50
# Output: ລົບຫ້າສິບ

laonumspell 100000
# Output: ໜຶ່ງແສນ

laonumspell 1000000
# Output: ໜຶ່ງລ້ານ
```

### ແປງຄຳເປັນຕົວເລກ / Convert Words to Number
```bash
laonumspell -r "ໜຶ່ງຮ້ອຍຊາວສາມ"
# Output: 123

laonumspell -r "ໜຶ່ງແສນ"
# Output: 100000

laonumspell -r "ສິບລ້ານ"
# Output: 10000000
```

### ຄຳສັ່ງອື່ນໆ / Other Commands
```bash
laonumspell -h          # ສະແດງຄຳແນະນຳ / Show help
laonumspell --help      # ສະແດງຄຳແນະນຳ / Show help
laonumspell -v          # ສະແດງເວີຊັນ / Show version
laonumspell --version   # ສະແດງເວີຊັນ / Show version
```

## 💻 ການໃຊ້ງານເປັນ Library / Library Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/xarhang/go-laonumspeller"
)

func main() {
    // ແປງຕົວເລກເປັນຄຳ / Number to Words
    result, err := laonumspeller.NumberToWordsLA(123.45)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
    // Output: ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ
    
    // ແປງຄຳເປັນຕົວເລກ / Words to Number
    num, err := laonumspeller.WordsToNumberLA("ໜຶ່ງແສນ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(num)
    // Output: 100000
    
    // ໃຊ້ງານກັບຕົວເລກຂະໜາດໃຫຍ່
    words, _ := laonumspeller.NumberToWordsLA(1000000000)
    fmt.Println(words) // Output: ໜຶ່ງຕື້
    
    // ແປງຕົວເລກລົບ
    negative, _ := laonumspeller.NumberToWordsLA(-50.5)
    fmt.Println(negative) // Output: ລົບຫ້າສິບຈຸດຫ້າ
}
```

### API Reference

#### `NumberToWordsLA(num float64) (string, error)`
ແປງຕົວເລກເປັນຄຳພາສາລາວ

**Parameters:**
- `num`: ຕົວເລກທີ່ຕ້ອງການແປງ (ຮອງຮັບທົດສະນິຍົມ ແລະ ເລກລົບ)

**Returns:**
- `string`: ຄຳພາສາລາວ
- `error`: ຂໍ້ຜິດພາດ (ຖ້າມີ)

#### `WordsToNumberLA(words string) (float64, error)`
ແປງຄຳພາສາລາວເປັນຕົວເລກ

**Parameters:**
- `words`: ຄຳພາສາລາວທີ່ຕ້ອງການແປງ

**Returns:**
- `float64`: ຕົວເລກ
- `error`: ຂໍ້ຜິດພາດ (ຖ້າມີ)

#### `IntToWordsLA(n int64) string`
ແປງຈຳນວນເຕັມເປັນຄຳພາສາລາວ (exported function ສຳລັບການໃຊ້ພິເສດ)

**Parameters:**
- `n`: ຈຳນວນເຕັມ

**Returns:**
- `string`: ຄຳພາສາລາວ

## 📊 ຕົວຢ່າງຕົວເລກຂະໜາດໃຫຍ່ / Large Number Examples

```bash
# ພັນ (Thousand - 10³)
laonumspell 1000           # ໜຶ່ງພັນ
laonumspell 10000          # ສິບພັນ
laonumspell 20000          # ຊາວພັນ
laonumspell 35000          # ສາມສິບຫ້າພັນ

# ແສນ (Hundred Thousand - 10⁵)
laonumspell 100000         # ໜຶ່ງແສນ
laonumspell 500000         # ຫ້າແສນ
laonumspell 150000         # ໜຶ່ງແສນຫ້າສິບພັນ

# ລ້ານ (Million - 10⁶)
laonumspell 1000000        # ໜຶ່ງລ້ານ
laonumspell 10000000       # ສິບລ້ານ
laonumspell 20000000       # ຊາວລ້ານ
laonumspell 100000000      # ໜຶ່ງຮ້ອຍລ້ານ

# ຕື້ (Billion - 10⁹)
laonumspell 1000000000     # ໜຶ່ງຕື້
laonumspell 10000000000    # ສິບຕື້
laonumspell 20000000000    # ຊາວຕື້
laonumspell 100000000000   # ໜຶ່ງຮ້ອຍຕື້

# ລ້ານລ້ານ (Trillion - 10¹²)
laonumspell 1000000000000  # ໜຶ່ງລ້ານລ້ານ
```

## 🔢 ຕາຕະລາງຫົວໜ່ວຍ / Unit Reference Table

| ເລກ (Number) | ພາສາລາວ (Lao) | ຄ່າ (Value) |
|---------------|----------------|-------------|
| 10¹ | ສິບ | 10 |
| 10² | ຮ້ອຍ | 100 |
| 10³ | ພັນ | 1,000 |
| 10⁴ | ສິບພັນ | 10,000 |
| 10⁵ | ແສນ | 100,000 |
| 10⁶ | ລ້ານ | 1,000,000 |
| 10⁷ | ສິບລ້ານ | 10,000,000 |
| 10⁸ | ຮ້ອຍລ້ານ | 100,000,000 |
| 10⁹ | ຕື້ | 1,000,000,000 |
| 10¹⁰ | ສິບຕື້ | 10,000,000,000 |
| 10¹¹ | ຮ້ອຍຕື້ | 100,000,000,000 |
| 10¹² | ລ້ານລ້ານ | 1,000,000,000,000 |

## 📁 ໂຄງສ້າງໂປຣເຈັກ / Project Structure

```
go-laonumspeller/
├── go.mod                    # Go module definition
├── speller.go                # Main library (package laonumspeller)
├── cmd/
│   └── laonumspell/
│       └── main.go           # CLI tool
├── LICENSE                   # Apache 2.0 License
└── README.md                 # Documentation
```

## 🛠️ Development

### ທົດສອບ Local / Test Locally
```bash
# Clone repository
git clone https://github.com/xarhang/go-laonumspeller.git
cd go-laonumspeller

# Run CLI directly
go run cmd/laonumspell/main.go 123.45
go run cmd/laonumspell/main.go -r "ໜຶ່ງຮ້ອຍຊາວສາມ"

# Build CLI
go build -o laonumspell cmd/laonumspell/main.go

# Run built binary
./laonumspell 999999

# Test as library
go test -v
```

### ສ້າງ Release / Build Release
```bash
# Build CLI for current platform
go build -o laonumspell cmd/laonumspell/main.go

# For Linux
GOOS=linux GOARCH=amd64 go build -o laonumspell-linux cmd/laonumspell/main.go

# For Windows
GOOS=windows GOARCH=amd64 go build -o laonumspell.exe cmd/laonumspell/main.go

# For macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o laonumspell-macos-amd64 cmd/laonumspell/main.go

# For macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o laonumspell-macos-arm64 cmd/laonumspell/main.go
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**xarhang**

- GitHub: [@xarhang](https://github.com/xarhang)

## 🙏 Acknowledgments

- Thanks to all contributors who help improve this project
- Inspired by the need for proper Lao language number conversion tools

## 📮 Support

If you have any questions or run into issues, please [open an issue](https://github.com/xarhang/go-laonumspeller/issues).

---

**Made with ❤️ for the Lao language community** 🇱🇦
