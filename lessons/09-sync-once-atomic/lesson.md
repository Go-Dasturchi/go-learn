# 09 — sync.Once and Atomic

## THEORY

**`sync.Once` — "bu faqat bir marta bo'lishi kerak" kafolati.** Raketani ishga tushirish tugmasini tasavvur qiling — uni necha marta bossangiz ham, raketa faqat **bir marta** uchishi kerak, ikkinchi bosishda hech narsa bo'lmasligi kerak. `sync.Once` — Go'da aynan shu kafolatni beradi: `Do(...)` ichidagi kod, uni **qancha marta va nechta goroutine'dan** chaqirishingizdan qat'iy nazar, faqat **bir marta** bajariladi:

```go
var once sync.Once

func sozlamalarniYukla() {
	once.Do(func() {
		fmt.Println("Sozlamalar yuklanmoqda...") // bu faqat BIR MARTA chop etiladi
	})
}

sozlamalarniYukla()
sozlamalarniYukla()
sozlamalarniYukla() // uch marta chaqirilsa ham, ichkarisi faqat bir marta ishlaydi
```

Bu ayniqsa "qimmat" (uzoq vaqt oladigan yoki resurs talab qiladigan) boshlang'ich sozlashni, ko'p joydan chaqirilishi mumkin bo'lgan holatlarda, **faqat bir marta** bajarilishini kafolatlash uchun foydali — hatto bir nechta goroutine bir vaqtda birinchi marta chaqirsa ham, `sync.Once` ularning faqat bittasiga ishlashga ruxsat beradi, qolganlari kutib, natijada hech kim ikkinchi marta ishlamaydi.

**Atomic — Mutex'siz, "bo'linmas" amallar.** "Mutex" darsida `Lock`/`Unlock` orqali umumiy o'zgaruvchini himoyalashni ko'rgan edik. Ba'zan, oddiy son ustida oddiy amallar (qo'shish, o'qish) uchun, to'liq Mutex o'rniga **atomik (bo'linmas) amallar** yetarli va tezroq bo'ladi — `sync/atomic` paketi buni ta'minlaydi:

```go
import "sync/atomic"

type HavfsizHisoblagichAtomic struct {
	son atomic.Int64
}

func (h *HavfsizHisoblagichAtomic) Oshir() {
	h.son.Add(1) // bo'linmas — hech qanday Lock/Unlock shart emas
}

func (h *HavfsizHisoblagichAtomic) Qiymat() int64 {
	return h.son.Load()
}
```

`atomic.Int64` — o'zida "bo'linmas o'qish/yozish" kafolatini olib yuradigan maxsus tur (Go'ning yangiroq versiyalarida qo'shilgan, qulay usul). `Add(1)` — qiymatni bo'linmas ravishda bittaga oshiradi: hatto minglab goroutine bir vaqtda `Add(1)` chaqirsa ham, **hech qanday oshirish "yo'qolmaydi"** — xuddi Mutex bilan himoyalangandek, lekin ancha yengilroq mexanizm bilan.

**Qachon Mutex, qachon Atomic?** Agar himoyalanishi kerak bo'lgan narsa — **bitta oddiy son yoki bayroq** bo'lsa (masalan, hisoblagich, "tugadimi" bayrog'i), `atomic` odatda yetarli va tezroq. Agar bir nechta o'zgaruvchini **birgalikda, izchil holatda** o'zgartirish kerak bo'lsa (masalan, struct'ning bir nechta maydonini bir vaqtda yangilash), Mutex to'g'riroq tanlov — chunki u butun "kritik bo'lim"ni, nechta amal bo'lishidan qat'iy nazar, yagona blok sifatida himoyalaydi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type HavfsizHisoblagichAtomic struct {
	son atomic.Int64
}

func (h *HavfsizHisoblagichAtomic) Oshir() {
	h.son.Add(1)
}

func main() {
	h := &HavfsizHisoblagichAtomic{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.Oshir()
		}()
	}
	wg.Wait()

	fmt.Println(h.son.Load()) // har doim aniq 1000
}
```

Natija:

```
1000
```

## TASK

`HavfsizHisoblagichAtomic` struct'i (`son atomic.Int64` maydoni bilan) berilgan. `Oshir()` methodini shunday to'ldiringki, u `atomic.Int64`ning `Add` methodi yordamida `son`ni bittaga xavfsiz oshirsin (Mutex ishlatmasdan).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `atomic.Int64` turida `Add(delta int64)` methodi bor — u qiymatga `delta` qo'shadi va bo'linmas ishlaydi.
2. `h.son.Add(1)` — funksiya tanasi shu bitta qatordan iborat bo'lishi mumkin.
