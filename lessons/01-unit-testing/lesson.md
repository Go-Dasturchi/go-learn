# 01 — Unit Testing

## THEORY

Mashina zavodida har bir g'ildirak, yig'ilgan mashinaga o'rnatilishidan oldin, **alohida** sinovdan o'tkaziladi — aylanadimi, mustahkammi, o'lchami to'g'rimi. Butun mashinani yig'ib bo'lib, keyin "ishlamayapti" deb hayron bo'lgandan ko'ra, har bir qismni alohida tekshirish ancha oson va tezroq muammoni topadi. **Unit test (birlik sinovi)** — dasturlashda aynan shu: kodning **eng kichik, mustaqil qismini** (odatda bitta funksiyani), qolgan dastur bilan aralashtirmasdan, alohida sinash.

**Aslida, siz bu kursning boshidan beri unit test bilan ishlab kelyapsiz!** Har bir dars, sizning yechimingizni tekshirish uchun, aynan **unit test** ishlatadi — hozir buni ochiq-oydin ko'rib chiqamiz.

**Test funksiyasi qanday yoziladi.** Go'da test — bu, nomi **`Test`** bilan boshlanadigan, va yagona parametri `*testing.T` bo'lgan oddiy funksiya:

```go
func TestKvadrat(t *testing.T) {
	got := Kvadrat(5)
	want := 25
	if got != want {
		t.Errorf("Kvadrat(5) = %d, kutilgan %d", got, want)
	}
}
```

Bu fayl **`_test.go`** bilan tugashi shart (masalan, `main_test.go`) — Go kompilyatori bunday fayllarni oddiy dastur qismidan ajratib, faqat `go test` buyrug'i orqali ishga tushiradi.

**`t.Errorf` va `t.Fatalf` — farqi.** Ikkalasi ham testni "muvaffaqiyatsiz" deb belgilaydi, lekin `t.Errorf` shundan keyin **testning qolgan qismini davom ettiradi** (agar keyingi tekshiruvlar ham bo'lsa), `t.Fatalf` esa **darhol to'xtatadi** — odatda, agar keyingi tekshiruvlar avvalgisiga bog'liq bo'lib, davom etish ma'nosiz bo'lsa (masalan, `nil` bo'lishi mumkin bo'lgan qiymatga murojaat qilishdan oldin).

**`go test` buyrug'i.** Terminalda `go test ./...` (yoki oddiygina `go test`) buyrug'i — joriy paketdagi (yoki barcha pastki paketlardagi) **barcha** `TestXxx` funksiyalarini avtomatik topib, ishga tushiradi va natijani ko'rsatadi:

```
$ go test -v
=== RUN   TestKvadrat
--- PASS: TestKvadrat (0.00s)
PASS
```

`-v` (verbose) bayrog'i — har bir test alohida qanday o'tganini ko'rsatadi; bayroqsiz esa faqat umumiy `PASS`/`FAIL` chiqadi.

**Nega bu muhim.** Qo'lda, `fmt.Println` bilan natijani ko'zdan kechirish — kichik dasturlar uchun ishlaydi, lekin loyiha kattalashgani sayin, **har safar** o'zgartirish kiritganda, hamma narsa hali ham to'g'ri ishlashini qo'lda tekshirib chiqish imkonsiz bo'lib qoladi. Testlar — bu tekshiruvni **avtomatlashtiradi**: bir marta yozib qo'yilgan test, kod necha marta o'zgarsa ham, bir zumda qayta ishga tushirilishi mumkin.

## EXAMPLE

```go
package main

import (
	"fmt"
	"testing"
)

func Kvadrat(x int) int {
	return x * x
}

func TestKvadratPositive(t *testing.T) {
	got := Kvadrat(5)
	want := 25
	if got != want {
		t.Errorf("Kvadrat(5) = %d, kutilgan %d", got, want)
	}
}

func TestKvadratZero(t *testing.T) {
	got := Kvadrat(0)
	want := 0
	if got != want {
		t.Errorf("Kvadrat(0) = %d, kutilgan %d", got, want)
	}
}

func main() {
	fmt.Println(Kvadrat(6))
}
```

Natija (`go test -v` orqali):

```
=== RUN   TestKvadratPositive
--- PASS: TestKvadratPositive (0.00s)
=== RUN   TestKvadratZero
--- PASS: TestKvadratZero (0.00s)
PASS
```

## TASK

`TubSonmi(n int) bool` funksiyasi berilgan. Uni shunday to'ldiringki, u `n` tub son (faqat 1 ga va o'ziga bo'linadigan, 1dan katta son) bo'lsa `true`, aks holda `false` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `1` va undan kichik sonlar tub emas: `if n <= 1 { return false }`. `2` dan `n-1`gacha (yoki samaraliroq, `n`ning kvadrat ildizigacha) bo'linuvchi qidiring.
2. `for i := 2; i*i <= n; i++ { if n%i == 0 { return false } }`, tsikldan keyin `return true`.
