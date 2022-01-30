package chinese_number

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var testDataSet = map[int64]string{
	0:                       "零",
	1:                       "一",
	10:                      "十",
	19:                      "十九",
	20:                      "二十",
	86:                      "八十六",
	100:                     "一百",
	101:                     "一百零一",
	110:                     "一百一十",
	111:                     "一百一十一",
	1000:                    "一千",
	1001:                    "一千零一",
	1010:                    "一千零一十",
	1011:                    "一千零一十一",
	1100:                    "一千一百",
	1101:                    "一千一百零一",
	1110:                    "一千一百一十",
	1111:                    "一千一百一十一",
	10000:                   "一万",
	10001:                   "一万零一",
	10010:                   "一万零一十",
	11000:                   "一万一千",
	100000:                  "十万",
	108001:                  "十万八千零一",
	1_0000_0000:             "一亿",
	1_0000_0001:             "一亿零一",
	1_0000_0000_0000_0000:   "一亿亿",
	10_0000_0000_0000_0000:  "十亿亿",
	100_0000_0000_0000_0000: "一百亿亿",
	math.MaxInt16:           "三万二千七百六十七",
	math.MaxInt32:           "二十一亿四千七百四十八万三千六百四十七",
	math.MaxInt64:           "九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零七",
}

func TestParseWangwenNumber(t *testing.T) {
	for num, han := range testDataSet {
		ret, err := Parse(han)
		if err != nil {
			t.Error(err)
			return
		}

		if ret != num {
			t.Errorf("the parsed result(%v) is not equal to %v", ret, num)
		}
	}
}

func TestConvert2WangwenNumber(t *testing.T) {
	for num, han := range testDataSet {
		ret, err := Convert(num)
		if err != nil {
			t.Error(err)
			return
		}

		if ret != han {
			t.Errorf("the parsed result(%v) is not the same as %v", ret, han)
		}
	}
}

func TestRandomNumber(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var num int64 = rand.Int63()
	han, err := Convert(num)
	if err != nil {
		t.Error(err)
		return
	}

	num2, err := Parse(han)
	if err != nil {
		t.Error(err)
		return
	}

	if num != num2 {
		t.Errorf("num: %v != %v", num, num2)
		return
	}

	han2, err := Convert(num2)
	if err != nil {
		t.Error(err)
		return
	}

	if han != han2 {
		t.Errorf("han: %v != %v", han, han2)
		return
	}
}

func BenchmarkParse(b *testing.B) {
	rand.Seed(time.Now().Unix())
	var num int64 = rand.Int63()
	han, err := Convert(num)
	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(han)
	}

}

func BenchmarkConvert(b *testing.B) {
	rand.Seed(time.Now().Unix())
	var num int64 = rand.Int63()
	han, err := Convert(num)
	if err != nil {
		b.Error(err)
		return
	}

	num2, err := Parse(han)
	if err != nil {
		b.Error(err)
		return
	}

	if num != num2 {
		b.Errorf("num: %v != %v", num, num2)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Convert(num2)
	}

}
