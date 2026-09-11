# 06 — Mutex

## THEORY

Bitta hammom va uning eshigida bitta kalit borligini tasavvur qiling — kalitni olgan odam ichkariga kirib, ishini bitirib, kalitni qaytarguncha, boshqa hech kim kira olmaydi. `sync.Mutex` ("mutual exclusion" — o'zaro chetlashtirish) — Go'da xuddi shu "yagona kalit" vazifasini bajaradi: bir vaqtning o'zida faqat **bitta** goroutine "kalitni ushlab turishi" (ya'ni himoyalangan kodni bajarishi) mumkin.

**Nega bu kerak — musobaqa holati (race condition).** Agar bir nechta goroutine **bir xil** o'zgaruvchini, hech qanday himoyasiz, bir vaqtda o'zgartirsa, natija **bashorat qilib bo'lmaydigan** bo'lib qoladi:

```go
son := 0
var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		son++ // XAVFLI: bir nechta goroutine bir vaqtda o'qib-yozadi
	}()
}
wg.Wait()
fmt.Println(son) // 1000 bo'lishi KERAK, lekin ko'pincha kamroq chiqadi!
```

Muammo shunda: `son++` aslida uchta qadam — "o'qi, bittaga oshir, qayta yoz". Agar ikkita goroutine buni **bir vaqtda** qilsa, ikkalasi ham bir xil "eski" qiymatni o'qib, ikkalasi ham bittaga oshirib qaytarsa — bitta oshirish "yo'qolib qoladi".

**Yechim — `Lock()` va `Unlock()`:**

```go
type HavfsizHisoblagich struct {
	mu  sync.Mutex
	son int
}

func (h *HavfsizHisoblagich) Oshir() {
	h.mu.Lock()
	h.son++
	h.mu.Unlock()
}
```

`Lock()` — "kalitni ol" (agar band bo'lsa, bo'shaguncha kut). `Unlock()` — "kalitni qaytar". `Lock()` va `Unlock()` orasidagi kod — **kritik bo'lim (critical section)** deyiladi, va faqat bitta goroutine bir vaqtning o'zida shu bo'limni bajarishi mumkin.

**`defer` bilan birga ishlatish odati.** Amalda, `Unlock()`ni unutmaslik uchun, uni ko'pincha `defer` bilan yozadilar ("Defer" darsini eslang — bu funksiya qanday tugashidan qat'iy nazar `Unlock` chaqirilishini kafolatlaydi):

```go
func (h *HavfsizHisoblagich) Oshir() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.son++
}
```

**Mutex struct'ning maydoni sifatida.** E'tibor bering — `sync.Mutex` odatda himoya qilinayotgan ma'lumot bilan **bir struct ichida** saqlanadi ("Structs" va "Methods" darslarini eslang). Bu — "qaysi qulf qaysi eshikka tegishli"ni aniq ko'rsatadi, va tasodifan boshqa ma'lumotni himoyalash uchun noto'g'ri mutex ishlatib qo'yish xavfini kamaytiradi.

**`Channel` yoki `Mutex` — qachon qaysi birini tanlash?** Ikkalasi ham "musobaqa holati"ning oldini oladi, lekin turli vaziyatlar uchun mos: agar goroutine'lar orasida **ma'lumot uzatish** kerak bo'lsa (masalan, natijani "topshirish") — channel qulayroq. Agar bir nechta goroutine **bitta umumiy holatni** (masalan, hisoblagichni) birgalikda o'zgartirishi kerak bo'lsa — Mutex to'g'ridan-to'g'riroq va tushunarliroq bo'ladi. Go jamoasining mashhur maqoli: "Do not communicate by sharing memory; instead, share memory by communicating" ("Xotirani baham ko'rish orqali muloqot qilmang; aksincha, muloqot qilish orqali xotirani baham ko'ring") — imkon qadar channel'ga moyil bo'lish tavsiya etiladi, lekin Mutex ham to'liq huquqli, ko'p ishlatiladigan vosita.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
)

type HavfsizHisoblagich struct {
	mu  sync.Mutex
	son int
}

func (h *HavfsizHisoblagich) Oshir() {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.son++
}

func main() {
	h := &HavfsizHisoblagich{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.Oshir()
		}()
	}
	wg.Wait()

	fmt.Println(h.son) // har doim aniq 1000
}
```

Natija:

```
1000
```

## TASK

`HavfsizHisoblagich` struct'i (`mu sync.Mutex`, `son int` maydonlari bilan) berilgan. `Oshir()` methodini shunday to'ldiringki, u `Mutex` yordamida **xavfsiz** ravishda `son`ni bittaga oshirsin — ya'ni ko'plab goroutine bir vaqtda `Oshir()`ni chaqirsa ham, natija hech qachon "yo'qolmasin".

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `h.mu.Lock()` bilan boshlang, `defer h.mu.Unlock()` qiling.
2. Undan keyin oddiygina `h.son++` yozing — `Lock`/`Unlock` orasida bo'lgani uchun, bu amal endi xavfsiz.
