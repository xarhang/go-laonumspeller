# go-laonumspeller

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://go.dev/)
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
go install github.com/xarhang/go-laonumspeller@latest
```

### Library
```bash
go get github.com/xarhang/go-laonumspeller
```

## 🚀 ການໃຊ້ງານ CLI / CLI Usage

### ແປງຕົວເລກເປັນຄຳ / Convert Number to Words
```bash
go-laonumspeller 123
# Output: ໜຶ່ງຮ້ອຍຊາວສາມ

go-laonumspeller 123.45
# Output: ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ

go-laonumspeller -50
# Output: ລົບຫ້າສິບ

go-laonumspeller 100000
# Output: ໜຶ່ງແສນ

go-laonumspeller 1000000
# Output: ໜຶ່ງລ້ານ
```

### ແປງຄຳເປັນຕົວເລກ / Convert Words to Number
```bash
go-laonumspeller -r "ໜຶ່ງຮ້ອຍຊາວສາມ"
# Output: 123

go-laonumspeller -r "ໜຶ່ງແສນ"
# Output: 100000

go-laonumspeller -r "ສິບລ້ານ"
# Output: 10000000
```

### ຄຳສັ່ງອື່ນໆ / Other Commands
```bash
go-laonumspeller -h          # ສະແດງຄຳແນະນຳ / Show help
go-laonumspeller --help      # ສະແດງຄຳແນະນຳ / Show help
```

## 💻 ການໃຊ້ງານເປັນ Library / Library Usage

```go
package main

import (
    "fmt"
    "log"
)

// Import functions directly from main package
// (Functions are in main.go: numberToWordsLA, wordsToNumberLA)

func main() {
    // Note: This package is primarily designed as a CLI tool.
    // For library usage, you can copy the conversion functions
    // from main.go into your own package.
    
    // Example of what the functions do:
    // numberToWordsLA(123.45) -> "ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ"
    // wordsToNumberLA("ໜຶ່ງຮ້ອຍຊາວສາມ") -> 123.0
}
```

## 📊 ຕົວຢ່າງຕົວເລກຂະໜາດໃຫຍ່ / Large Number Examples

```bash
# ພັນ (Thousand - 10³)
go-laonumspeller 1000           # ໜຶ່ງພັນ
go-laonumspeller 10000          # ສິບພັນ
go-laonumspeller 20000          # ຊາວພັນ
go-laonumspeller 35000          # ສາມສິບຫ້າພັນ

# ແສນ (Hundred Thousand - 10⁵)
go-laonumspeller 100000         # ໜຶ່ງແສນ
go-laonumspeller 500000         # ຫ້າແສນ
go-laonumspeller 150000         # ໜຶ່ງແສນຫ້າສິບພັນ

# ລ້ານ (Million - 10⁶)
go-laonumspeller 1000000        # ໜຶ່ງລ້ານ
go-laonumspeller 10000000       # ສິບລ້ານ
go-laonumspeller 20000000       # ຊາວລ້ານ
go-laonumspeller 100000000      # ໜຶ່ງຮ້ອຍລ້ານ

# ຕື້ (Billion - 10⁹)
go-laonumspeller 1000000000     # ໜຶ່ງຕື້
go-laonumspeller 10000000000    # ສິບຕື້
go-laonumspeller 20000000000    # ຊາວຕື້
go-laonumspeller 100000000000   # ໜຶ່ງຮ້ອຍຕື້

# ລ້ານລ້ານ (Trillion - 10¹²)
go-laonumspeller 1000000000000  # ໜຶ່ງລ້ານລ້ານ
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
├── go.mod              # Go module definition
├── main.go             # CLI tool & conversion functions
├── LICENSE             # Apache 2.0 License
└── README.md           # Documentation
```

## 🛠️ Development

### ທົດສອບ Local / Test Locally
```bash
# Clone repository
git clone https://github.com/xarhang/go-laonumspeller.git
cd go-laonumspeller

# Run directly
go run main.go 123.45
go run main.go -r "ໜຶ່ງຮ້ອຍຊາວສາມ"

# Build
go build -o go-laonumspeller

# Run built binary
./go-laonumspeller 999999
```

### ສ້າງ Release / Build Release
```bash
# For current platform
go build -o go-laonumspeller

# For Linux
GOOS=linux GOARCH=amd64 go build -o go-laonumspeller-linux

# For Windows
GOOS=windows GOARCH=amd64 go build -o go-laonumspeller.exe

# For macOS
GOOS=darwin GOARCH=amd64 go build -o go-laonumspeller-macos
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