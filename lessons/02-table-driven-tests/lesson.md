# 02 — Table Driven Tests

## THEORY

Diqqat bilan qarasangiz, ushbu kursning deyarli har bir darsida hidden test fayllari bir xil naqshda yozilgan: `TestFunksiyaCase1`, `TestFunksiyaCase2`, `TestFunksiyaCase3` — har biri deyarli bir xil kodni takrorlaydi, faqat kirish va kutilgan qiymatlar boshqacha. Bu — ataylab, o'rganish bosqichida **soddalik** uchun shunday qilingan. Lekin professional Go kodida, bunday takrorlanishni kamaytirish uchun **table driven tests** (jadval asosidagi testlar) deb ataladigan naqsh keng qo'llaniladi.

**G'oya — barcha holatlarni bitta "jadval"ga yig'ish.** Har bir holat uchun alohida funksiya yozish o'rniga, barcha kirish/kutilgan-natija juftliklarini **bitta slice**ga yig'ib, ustidan bitta `for` tsikli bilan yuriladi:

```go
func TestBelgi(t *testing.T) {
	holatlar := []struct {
		nomi     string
		kirish   int
		kutilgan string
	}{
		{"musbat", 5, "musbat"},
		{"manfiy", -3, "manfiy"},
		{"nol", 0, "nol"},
	}

	for _, h := range holatlar {
		t.Run(h.nomi, func(t *testing.T) {
			got := Belgi(h.kirish)
			if got != h.kutilgan {
				t.Errorf("Belgi(%d) = %q, kutilgan %q", h.kirish, got, h.kutilgan)
			}
		})
	}
}
```

`holatlar` — anonim struct'lardan iborat slice ("Structs" va "Slices" darslarini eslang, bu ikkalasini birlashtirgan holat). Har bir element — bitta test holatini (nomi, kirish, kutilgan natija) ifodalaydi.

**`t.Run` — ichki subtest yaratish.** `t.Run(nomi, func(t *testing.T) {...})` — jadvaldagi har bir qatorni **alohida, nomlangan subtest** sifatida ishga tushiradi. Bu, `go test -v` natijasida har bir holatni alohida ko'rish imkonini beradi:

```
=== RUN   TestBelgi
=== RUN   TestBelgi/musbat
=== RUN   TestBelgi/manfiy
=== RUN   TestBelgi/nol
--- PASS: TestBelgi (0.00s)
    --- PASS: TestBelgi/musbat (0.00s)
    --- PASS: TestBelgi/manfiy (0.00s)
    --- PASS: TestBelgi/nol (0.00s)
```

**Nega bu foydali.** Yangi holat qo'shish uchun endi **yangi funksiya yozish shart emas** — shunchaki jadvalga bitta yangi qator qo'shiladi. Bu kodni qisqartiradi, takrorlanishni kamaytiradi, va barcha holatlarni **bitta joyda**, ko'rish/solishtirish oson bo'lgan shaklda saqlaydi — ayniqsa o'nlab holatlarni tekshirish kerak bo'lganda, bu "Nested Loops" darsida ko'rgan samaradorlik g'oyasiga o'xshab, kodni ancha ixchamlashtiradi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"testing"
)

func Belgi(n int) string {
	if n > 0 {
		return "musbat"
	} else if n < 0 {
		return "manfiy"
	}
	return "nol"
}

func TestBelgi(t *testing.T) {
	holatlar := []struct {
		nomi     string
		kirish   int
		kutilgan string
	}{
		{"musbat", 5, "musbat"},
		{"manfiy", -3, "manfiy"},
		{"nol", 0, "nol"},
	}

	for _, h := range holatlar {
		t.Run(h.nomi, func(t *testing.T) {
			got := Belgi(h.kirish)
			if got != h.kutilgan {
				t.Errorf("Belgi(%d) = %q, kutilgan %q", h.kirish, got, h.kutilgan)
			}
		})
	}
}

func main() {
	fmt.Println(Belgi(5), Belgi(-3), Belgi(0))
}
```

Natija:

```
musbat manfiy nol
```

## TASK

`Belgi(n int) string` funksiyasi berilgan. Uni shunday to'ldiringki, u `n` musbat bo'lsa `"musbat"`, manfiy bo'lsa `"manfiy"`, nolga teng bo'lsa `"nol"` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if n > 0 { return "musbat" }` bilan boshlang.
2. `else if n < 0 { return "manfiy" }`, aks holda `return "nol"`.
