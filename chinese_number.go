package chinese_number

import (
	"bytes"
	"errors"
	"fmt"
	"unicode"
)

// Parse map
var numTable = map[rune]int64{
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
	'白': 100,
	'干': 1000,
	'两': 2,
	'兩': 2,
	'〇': 0,
}

// Convert map
var defaultHanTable = map[int64]rune{
	0:           '零',
	1:           '一',
	2:           '二',
	3:           '三',
	4:           '四',
	5:           '五',
	6:           '六',
	7:           '七',
	8:           '八',
	9:           '九',
	10:          '十',
	100:         '百',
	1000:        '千',
	1_0000:      '万',
	1_0000_0000: '亿',
}

func init() {
	// add 0-9 to Parse map
	for i := 0; i < 10; i++ {
		key := rune('0' + i)
		numTable[key] = int64(i)
	}

	for i := 0; i < 26; i++ {
		// add a-z to Parse map
		key := rune('a' + i)
		numTable[key] = 10 + int64(i)

		// add A-Z to Parse map
		key = rune('A' + i)
		numTable[key] = 10 + int64(i)
	}

}

// Get Value from Parse map. This function is Only using in Parse.
func getNumValue(ch rune) (int64, error) {
	n, ok := numTable[ch]
	if !ok {
		return n, fmt.Errorf("unrecognized chinese-number: %c", ch)
	}

	return n, nil
}

// Parse chinese number into integer(int64).
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

		// record the biggest chinese number
		if n1 > max {
			max = n1
		}

		// if current chinese number is 万 or 亿
		if n1 == 1_0000 || n1 == 1_0000_0000 {
			// handle conner case '万亿' and '亿亿'
			if last != 1_0000 && last != 1_0000_0000 {
				// parse a part of chinese number ahead of '万' or '亿'
				// such as 一千九百八十五万, it should parse part 一千九百八十五
				tem, err = parse10000(string(hans[begin:i]))
				if err != nil {
					return
				}
				// such as '一亿零五万'
				if n1 >= max {
					// part 一亿
					num += tem
					num *= n1
				} else {
					// part 零五万
					num += tem * n1
				}
			} else {
				// conner case 万亿 and 亿亿
				num *= n1
			}

			begin = int64(i) + 1
		}
		last = n1
	}

	// parse rest part that less than 万
	// such as 一万五千八百零五, it will parse the rest part 五千八百零五
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

	// handle conner case 十 十一 十二 ... 十九
	if val >= 10 {
		tem = 1
	}

	for i := range hans[:len(hans)-1] {
		// get current chinese number
		if n1, err = getNumValue(hans[i]); err != nil {
			return
		}

		// get next chinese number
		if n2, err = getNumValue(hans[i+1]); err != nil {
			return
		}

		if n1 < 10 {
			if unicode.IsDigit(hans[i]) {
				// just a number
				tem = tem*10 + n1
			} else {
				// chinese number
				tem = n1
			}
		} else {
			// n1 is 十 or n1 is 百 or n1 is 千
			tem *= n1
		}

		// 一千一百一十一
		if n2 < n1 && n1 >= 10 {
			// plus 一千 first
			// then plus 一百
			// then plus 一十
			num += tem
			tem = 0
		}
	}
	// finally, plus 一
	num += tem
	return
}

func addExceptEmpty(han string) string {
	if len(han) == 0 {
		return ""
	}
	return han
}

func Convert(num int64) (han string, err error) {
	if num < 0 {
		return han, errors.New("chinese-number only support the number greater than or equal to zero")
	}

	var buf bytes.Buffer
	hanTable := defaultHanTable

	if num == 0 {
		han = string(hanTable[0])
		return
	}

	num1 := num

	// count 10000
	count := 0

	num2 := num1 % 10000
	num1 /= 10000
	last := num1 == 0
	tem, _ := convert10000(num2, last, hanTable)
	buf.WriteString(addExceptEmpty(tem))

	for num1 > 0 {
		// case 一万(零)一百, 一万(零)一十
		if num2 < 1000 && num2 > 0 {
			zero := hanTable[0]
			buf.WriteRune(zero)
		}

		wan := hanTable[10000]
		if count&1 == 0 {
			// 10000
			buf.WriteRune(wan)
		} else {

			yi := hanTable[1_0000_0000]
			tStr := buf.String()
			tRunes := []rune(tStr)

			if tRunes[len(tRunes)-1] == wan {
				// if there is a 万(1_0000) ahead of current chinese number
				// it should turn to 亿 (1_0000_0000)
				tRunes[len(tRunes)-1] = yi
				buf.Reset()
				buf.WriteString(string(tRunes))
			} else {
				buf.WriteRune(yi)
			}
		}

		num2 = num1 % 10000
		num1 /= 10000
		last = num1 == 0
		tem, _ := convert10000(num2, last, hanTable)
		buf.WriteString(addExceptEmpty(tem))

		count++
	}

	han = buf.String()

	han = reverse(han)

	return
}

func convert10000(num int64, last bool, hanTable map[int64]rune) (han string, err error) {
	var buf bytes.Buffer

	var (
		num1          = num
		base    int64 = 1
		tem     rune
		lastNum int64
		total   int64
	)

	for num1 > 0 {
		currentNum := num1 % 10
		total = total*10 + currentNum

		if base >= 10 && currentNum > 0 {
			tem = hanTable[base]
			buf.WriteRune(tem)
		}

		// such as 10 100 1000 should write 十 百 千 first
		if total != 0 && !(currentNum == 0 && lastNum == 0) {
			tem = hanTable[currentNum]
			buf.WriteRune(tem)
		}

		lastNum = currentNum
		base *= 10
		num1 /= 10

	}

	han = buf.String()

	// conner case 一十 一十一 一十二 ... 一十九
	// there should be 十 十一 十二 ... 十九
	if num >= 10 && num < 20 && last {
		hans := []rune(han)
		han = string(hans[:len(hans)-1])
	}

	return
}

func reverse(han string) string {
	hans := []rune(han)
	length := len(hans)
	times := length / 2
	for i := 0; i < times; i++ {
		hans[i], hans[length-1-i] = hans[length-1-i], hans[i]
	}

	return string(hans)
}
