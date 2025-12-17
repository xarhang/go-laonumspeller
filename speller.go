package laonumspeller

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// NumberToWordsLA converts a float64 value into a Lao word string.
// Example: 123.45 -> "ໜຶ່ງຮ້ອຍຊາວສາມຈຸດສີ່ສິບຫ້າ"
func NumberToWordsLA(num float64) (string, error) {
	if math.IsInf(num, 0) || math.IsNaN(num) {
		return "", fmt.Errorf("not a finite number: %v", num)
	}

	if num == 0 {
		return "ສູນ", nil
	}

	negative := num < 0
	num = math.Abs(num)

	intPart := int64(num)
	decimalPart := num - float64(intPart)

	result := IntToWordsLA(intPart)

	if decimalPart > 0 {
		decimalStr := fmt.Sprintf("%.10f", decimalPart)
		decimalStr = strings.TrimPrefix(decimalStr, "0.")
		decimalStr = strings.TrimRight(decimalStr, "0")

		if decimalStr != "" {
			result += "ຈຸດ"
			decimalNum, _ := strconv.ParseInt(decimalStr, 10, 64)
			result += IntToWordsLA(decimalNum)
		}
	}

	if negative {
		result = "ລົບ" + result
	}

	return result, nil
}

// IntToWordsLA converts an integer to Lao words.
func IntToWordsLA(n int64) string {
	if n == 0 {
		return "ສູນ"
	}

	// ລ້ານລ້ານ (Trillion - 10^12)
	if n >= 1000000000000 {
		trillions := n / 1000000000000
		remainder := n % 1000000000000
		result := IntToWordsLA(trillions) + "ລ້ານລ້ານ"
		if remainder > 0 {
			result += IntToWordsLA(remainder)
		}
		return result
	}

	// ຕື້ (Billion - 10^9)
	if n >= 1000000000 {
		billions := n / 1000000000
		remainder := n % 1000000000
		billionPart := formatSegment(billions, "ຕື້")
		result := billionPart
		if remainder > 0 {
			result += IntToWordsLA(remainder)
		}
		return result
	}

	// ລ້ານ (Million - 10^6)
	if n >= 1000000 {
		millions := n / 1000000
		remainder := n % 1000000
		millionPart := formatSegment(millions, "ລ້ານ")
		result := millionPart
		if remainder > 0 {
			result += IntToWordsLA(remainder)
		}
		return result
	}

	// ແສນ (Hundred Thousand - 10^5)
	if n >= 100000 {
		hundredThousands := n / 100000
		remainder := n % 100000
		hundredThousandPart := ""
		if hundredThousands == 1 {
			hundredThousandPart = "ໜຶ່ງແສນ"
		} else {
			hundredThousandPart = convertSmallInt(hundredThousands) + "ແສນ"
		}
		result := hundredThousandPart
		if remainder > 0 {
			result += IntToWordsLA(remainder)
		}
		return result
	}

	// ພັນ (Thousand - 10^3)
	if n >= 1000 {
		thousands := n / 1000
		remainder := n % 1000
		thousandPart := formatSegment(thousands, "ພັນ")
		result := thousandPart
		if remainder > 0 {
			result += IntToWordsLA(remainder)
		}
		return result
	}

	return convertSmallInt(n)
}

// formatSegment formats number segments (for ພັນ, ລ້ານ, ຕື້)
func formatSegment(val int64, unit string) string {
	if val < 10 {
		if val == 1 && unit == "ພັນ" {
			return "ໜຶ່ງພັນ"
		}
		return convertSmallInt(val) + unit
	}
	return convertSmallInt(val) + unit
}

// convertSmallInt converts numbers under 1000 to Lao words
func convertSmallInt(n int64) string {
	if n == 0 {
		return ""
	}

	s := strconv.FormatInt(n, 10)
	numtext := []string{"", "ໜຶ່ງ", "ສອງ", "ສາມ", "ສີ່", "ຫ້າ", "ຫົກ", "ເຈັດ", "ແປດ", "ເກົ້າ"}
	var res strings.Builder
	et := false

	for i := 0; i < len(s); i++ {
		val := int(s[i] - '0')
		pos := len(s) - i - 1

		if val > 0 {
			switch pos {
			case 2: // ຮ້ອຍ
				if val == 1 {
					res.WriteString("ໜຶ່ງຮ້ອຍ")
				} else {
					res.WriteString(numtext[val] + "ຮ້ອຍ")
				}
			case 1: // ສິບ
				switch val {
				case 1:
					res.WriteString("ສິບ")
				case 2:
					res.WriteString("ຊາວ")
				default:
					res.WriteString(numtext[val] + "ສິບ")
				}
				et = true
			case 0: // ຫົວໜ່ວຍ
				if et && val == 1 {
					res.WriteString("ເອັດ")
				} else {
					res.WriteString(numtext[val])
				}
				et = false
			}
		} else if pos == 1 {
			et = false
		}
	}
	return res.String()
}

// WordsToNumberLA converts Lao words into a float64.
func WordsToNumberLA(words string) (float64, error) {
	words = strings.TrimSpace(words)
	if words == "ສູນ" {
		return 0, nil
	}

	negative := false
	if strings.HasPrefix(words, "ລົບ") {
		negative = true
		words = strings.TrimPrefix(words, "ລົບ")
	}

	parts := strings.Split(words, "ຈຸດ")
	intPart, err := wordsToIntLA(parts[0])
	if err != nil {
		return 0, err
	}

	result := float64(intPart)
	if len(parts) > 1 {
		decimalInt, err := wordsToIntLA(parts[1])
		if err == nil && decimalInt > 0 {
			decimalStr := strconv.FormatInt(decimalInt, 10)
			decimalValue, _ := strconv.ParseFloat("0."+decimalStr, 64)
			result += decimalValue
		}
	}

	if negative {
		result = -result
	}
	return result, nil
}

// wordsToIntLA converts Lao words to an integer
func wordsToIntLA(words string) (int64, error) {
	words = strings.TrimSpace(words)
	if words == "" || words == "ສູນ" {
		return 0, nil
	}

	numMap := map[string]int64{
		"ໜຶ່ງ": 1, "ສອງ": 2, "ສາມ": 3, "ສີ່": 4, "ຫ້າ": 5,
		"ຫົກ": 6, "ເຈັດ": 7, "ແປດ": 8, "ເກົ້າ": 9,
		"ສິບ": 10, "ຊາວ": 20, "ເອັດ": 1,
		"ຮ້ອຍ": 100, "ພັນ": 1000, "ແສນ": 100000,
		"ລ້ານ": 1000000, "ຕື້": 1000000000,
		"ລ້ານລ້ານ": 1000000000000,
	}

	var total, current int64
	runes := []rune(words)

	for i := 0; i < len(runes); {
		matched := false
		for length := 10; length >= 1; length-- {
			if i+length <= len(runes) {
				substr := string(runes[i : i+length])
				if val, ok := numMap[substr]; ok {
					if val >= 1000 {
						if current == 0 {
							current = 1
						}
						if val >= 1000000 {
							total += current * val
							current = 0
						} else {
							current *= val
						}
					} else if val >= 100 {
						if current == 0 {
							current = 1
						}
						current *= val
					} else if val == 10 || val == 20 {
						current += val
					} else {
						current += val
					}
					i += length
					matched = true
					break
				}
			}
		}
		if !matched {
			i++
		}
	}
	return total + current, nil
}
