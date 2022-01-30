package chinese_number

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var numtable = map[rune]int64{
	'零': 0,
	'一': 1,
	'二': 2,
	'三': 3,
	'四': 4,
	'五': 5,
	'六': 6,
	'七': 7,
	'八': 8,
	'九': 9,
	'十': 10,
	'百': 100,
	'千': 1000,
	'万': 10000,
	'亿': 1_0000_0000,
	'壹': 1,
	'贰': 2,
	'叁': 3,
	'肆': 4,
	'伍': 5,
	'陆': 6,
	'柒': 7,
	'捌': 8,
	'玖': 9,
	'拾': 10,
	'佰': 100,
	'仟': 1000,
	'干': 1000,
	'两': 2,
	'兩': 2,
	'〇': 0,
}

var defaultHanTable = map[int64]string{
	0:           "零",
	1:           "一",
	2:           "二",
	3:           "三",
	4:           "四",
	5:           "五",
	6:           "六",
	7:           "七",
	8:           "八",
	9:           "九",
	10:          "十",
	100:         "百",
	1000:        "千",
	1_0000:      "万",
	1_0000_0000: "亿",
}

func init() {
	for i := 0; i < 10; i++ {
		key := rune('0' + i)
		numtable[key] = int64(i)
	}

	for i := 0; i < 26; i++ {
		key := rune('a' + i)
		numtable[key] = 10 + int64(i)

		key = rune('A' + i)
		numtable[key] = 10 + int64(i)
	}

}

func getNumValue(ch rune) (int64, error) {
	n, ok := numtable[ch]
	if !ok {
		return n, fmt.Errorf("unrecognized wangwen-number: %c", ch)
	}

	return n, nil
}

func Parse(han string) (num int64, err error) {
	if len(han) == 0 {
		return
	}

	hans := []rune(han)

	var n1, tem, begin, last, max int64

	for i := range hans {
		if n1, err = getNumValue(hans[i]); err != nil {
			return
		}

		if n1 > max {
			max = n1
		}

		if n1 == 1_0000 || n1 == 1_0000_0000 {
			if last != 1_0000 && last != 1_0000_0000 {
				tem, err = parse10000(string(hans[begin:i]))
				if err != nil {
					return
				}
				if n1 >= max {
					num += tem
					num *= n1
				} else {
					num += tem * n1
				}
			} else {
				num *= n1
			}

			begin = int64(i) + 1
		}
		last = n1

	}

	tem, err = parse10000(string(hans[begin:]))
	if err != nil {
		return
	}
	num += tem

	return
}

func parse10000(han string) (num int64, err error) {
	if len(han) == 0 {
		return
	}

	hans := []rune(han)

	hans = append(hans, '0')
	var n1, n2, tem int64
	var val int64

	val, err = getNumValue(hans[0])
	if err != nil {
		return
	}

	if val >= 10 {
		tem = 1
	}

	for i := range hans[:len(hans)-1] {
		if n1, err = getNumValue(hans[i]); err != nil {
			return
		}

		if n2, err = getNumValue(hans[i+1]); err != nil {
			return
		}

		if n1 < 10 {
			if unicode.IsDigit(hans[i]) {
				tem = tem*10 + n1
			} else {
				tem = n1
			}
		} else {
			tem *= n1
		}

		if n2 < n1 && n1 >= 10 {
			num += tem
			tem = 0
		}
	}
	num += tem
	return
}

func addExceptEmpty(hans []string, han string) []string {
	if len(han) == 0 {
		return hans
	}
	return append([]string{han}, hans...)
}

func Convert(num int64) (han string, err error) {
	if num < 0 {
		return han, errors.New("wangwen-number only support the number greater than or equal to zero")
	}

	var hans []string
	hanTable := defaultHanTable

	if num == 0 {
		han = hanTable[0]
		return
	}

	num1 := num
	count := 0

	num2 := num1 % 10000
	num1 /= 10000
	last := num1 == 0
	tem, _ := convert10000(num2, last, hanTable)
	hans = addExceptEmpty(hans, tem)

	for num1 > 0 {
		if num2 < 1000 && num2 > 0 {
			zero := hanTable[0]
			hans = append([]string{zero}, hans...)
		}

		wan := hanTable[10000]
		if count&1 == 0 {
			// 10000
			hans = append([]string{wan}, hans...)
		} else {
			// 1_0000_0000
			yi := hanTable[1_0000_0000]
			if hans[0] == wan {
				hans[0] = yi
			} else {
				hans = append([]string{yi}, hans...)
			}
		}

		num2 = num1 % 10000
		num1 /= 10000
		last = num1 == 0
		tem, _ := convert10000(num2, last, hanTable)
		hans = addExceptEmpty(hans, tem)

		count++
	}

	han = strings.Join(hans, "")

	return
}

func convert10000(num int64, last bool, hanTable map[int64]string) (han string, err error) {
	var hans []string

	var (
		num1          = num
		base    int64 = 1
		tem     string
		lastNum int64
		total   int64
	)

	for num1 > 0 {
		currentNum := num1 % 10
		total = total*10 + currentNum

		if base >= 10 && currentNum > 0 {
			tem = hanTable[base]
			hans = append([]string{tem}, hans...)
		}

		// 100 1000
		if total != 0 && !(currentNum == 0 && lastNum == 0) {
			tem = hanTable[currentNum]
			hans = append([]string{tem}, hans...)
		}

		lastNum = currentNum
		base *= 10
		num1 /= 10

	}

	if num >= 10 && num < 20 && last {
		hans = hans[1:]
	}

	han = strings.Join(hans, "")

	return
}
