package decimal_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/robaho/fixed"
	"github.com/shopspring/decimal"
)

func ApdTest() {
	s1 := rand.NewSource(time.Now().UnixMicro())
	r1 := rand.New(s1)
	randVal := r1.Float64() + float64(rand.Int63n(1999888777))

	randVal1 := r1.Float64() * 5

	a, _, _ := apd.NewFromString(fmt.Sprintf("%f", randVal))
	b, _, _ := apd.NewFromString(fmt.Sprintf("%f", randVal1))

	apd.BaseContext.WithPrecision(20).Mul(a, a, b)
}

func ShopspringMulTest() {
	s1 := rand.NewSource(time.Now().UnixMicro())
	r1 := rand.New(s1)
	randVal := r1.Float64() + float64(rand.Int63n(1999888777))

	randVal1 := r1.Float64() * 5

	a, _ := decimal.NewFromString(fmt.Sprintf("%f", randVal))
	b, _ := decimal.NewFromString(fmt.Sprintf("%f", randVal1))

	a.Mul(b)
}

func FixedMulTest() {
	s1 := rand.NewSource(time.Now().UnixMicro())
	r1 := rand.New(s1)
	randVal := r1.Float64() + float64(rand.Int63n(1999888777))

	randVal1 := r1.Float64() * 5

	a := fixed.NewS(fmt.Sprintf("%f", randVal))
	b := fixed.NewS(fmt.Sprintf("%f", randVal1))

	a.Mul(b)
}

func BenchmarkMulShopspring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShopspringMulTest()
	}
}

func BenchmarkMulApd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ApdTest()
	}
}

func BenchmarkMulFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FixedMulTest()
	}
}
