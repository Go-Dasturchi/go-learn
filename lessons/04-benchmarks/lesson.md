# 04 — Benchmarks

## THEORY

Ikki yo'lchi bir manzilga borishning ikki xil yo'lini sinab ko'rmoqchi — kim tezroq yetib borishini bilish uchun, ikkalasi ham soatini yoqib, vaqtni o'lchaydi. Unit test dasturning **to'g'ri ishlashini** tekshirsa, **benchmark (o'lchov sinovi)** — dasturning **qanchalik tez** ishlashini o'lchaydi.

**Benchmark funksiyasi qanday yoziladi.** Nomi `Benchmark` bilan boshlanadi, va parametri `*testing.T` emas, **`*testing.B`**:

```go
func BenchmarkStringQoshish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		natija := ""
		for j := 0; j < 100; j++ {
			natija += "a"
		}
		_ = natija
	}
}
```

**`b.N` — Go'ning o'zi aniqlaydigan takrorlash soni.** Siz `b.N`ni o'zingiz belgilamaysiz — Go benchmark'ni turli `N` qiymatlari bilan bir necha marta ishga tushirib, natija **statistik jihatdan ishonchli** bo'lguncha (juda tez ishlaydigan kod uchun ko'proq marta, sekinroq kod uchun kamroq marta) avtomatik moslashtiradi.

**Ishga tushirish — `go test -bench`.** Oddiy `go test` benchmark'larni **ishga tushirmaydi** (ular sekin bo'lishi mumkin) — buning uchun maxsus bayroq kerak:

```bash
go test -bench=.
```

Natija taxminan shunday ko'rinadi:

```
BenchmarkStringQoshish-8   	  123456	      9532 ns/op
```

`9532 ns/op` — har bir takrorlash o'rtacha necha nanosekund vaqt olganini bildiradi. Bu son qanchalik **kichik** bo'lsa, kod shunchalik **tezroq** ishlaydi.

**Amaliy misol — nega `strings.Builder` tezroq.** "Strings" darsida ko'rgan `+` orqali birlashtirish, har safar **yangi** string yaratadi (stringlar o'zgarmas, eslaysizmi?) — bu ko'p marta takrorlansa, sekinlashadi. `strings.Builder` esa ichki buferni **qayta ishlatib**, ancha samaraliroq ishlaydi:

```go
var b strings.Builder
for i := 0; i < 100; i++ {
	b.WriteString("a")
}
natija := b.String()
```

Aynan shu ikki yondashuvni benchmark qilib solishtirsangiz, `strings.Builder` versiyasi sezilarli darajada tezroq chiqishini ko'rasiz — bu, benchmark'larning haqiqiy amaliy foydasini yaqqol ko'rsatadi: "qaysi yechim tezroq" degan savolga **taxmin** emas, **o'lchangan raqam** bilan javob berish.

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
	"testing"
)

func StringQurish(qismlar []string) string {
	var b strings.Builder
	for _, q := range qismlar {
		b.WriteString(q)
	}
	return b.String()
}

func BenchmarkStringQurish(b *testing.B) {
	qismlar := []string{"Salom", ", ", "dunyo", "!"}
	for i := 0; i < b.N; i++ {
		StringQurish(qismlar)
	}
}

func main() {
	fmt.Println(StringQurish([]string{"Salom", ", ", "dunyo", "!"}))
}
```

Natija (`go test -bench=.` orqali, taxminiy):

```
BenchmarkStringQurish-8   	 5000000	       234 ns/op
```

## TASK

`StringQurish(qismlar []string) string` funksiyasi berilgan. Uni `strings.Builder` yordamida shunday to'ldiringki, u `qismlar` slice'idagi barcha satrlarni ketma-ket birlashtirib, yagona string qilib qaytarsin (masalan, `StringQurish([]string{"a", "b", "c"})` → `"abc"`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `var b strings.Builder` bilan boshlang, keyin `for _, q := range qismlar { b.WriteString(q) }`.
2. `return b.String()` — buferga yig'ilgan barcha qismlarni yagona string qilib qaytaradi.
