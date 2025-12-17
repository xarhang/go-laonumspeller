package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xarhang/go-laonumspeller"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-r", "--reverse":
		// ແປງຄຳເປັນຕົວເລກ
		if len(os.Args) < 3 {
			fmt.Println("ຂໍ້ຜິດພາດ: ກະລຸນາໃສ່ຄຳພາສາລາວ")
			printUsage()
			os.Exit(1)
		}

		result, err := laonumspeller.WordsToNumberLA(os.Args[2])
		if err != nil {
			fmt.Printf("ຂໍ້ຜິດພາດ: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)

	case "-h", "--help":
		printUsage()

	case "-v", "--version":
		fmt.Println("go-laonumspeller v1.0.0")

	default:
		// ແປງຕົວເລກເປັນຄຳ
		num := os.Args[1]

		floatNum, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Printf("ຂໍ້ຜິດພາດ: '%v' ບໍ່ແມ່ນຕົວເລກທີ່ຖືກຕ້ອງ\n", num)
			os.Exit(1)
		}

		result, err := laonumspeller.NumberToWordsLA(floatNum)
		if err != nil {
			fmt.Printf("ຂໍ້ຜິດພາດ: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}

func printUsage() {
	fmt.Println("Lao Numeric Speller - ແປງຕົວເລກເປັນພາສາລາວ")
	fmt.Println()
	fmt.Println("ວິທີໃຊ້:")
	fmt.Println("  laonumspell <number>              ແປງຕົວເລກເປັນຄຳ")
	fmt.Println("  laonumspell -r <lao-words>        ແປງຄຳເປັນຕົວເລກ")
	fmt.Println("  laonumspell -h, --help            ສະແດງຄຳແນະນຳນີ້")
	fmt.Println("  laonumspell -v, --version         ສະແດງເວີຊັນ")
	fmt.Println()
	fmt.Println("ຕົວຢ່າງ:")
	fmt.Println("  laonumspell 123.45")
	fmt.Println("  laonumspell -r \"ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ\"")
	fmt.Println()
}
