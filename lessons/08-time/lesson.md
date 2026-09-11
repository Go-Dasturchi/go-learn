# 08 — Time and Duration

## THEORY

Sportchi tayyorgarlik uchun necha daqiqa mashq qilganini stopwatch bilan o'lchaganini tasavvur qiling — stopwatch faqat **o'tgan vaqtni** ko'rsatadi, aniq kalendar sanasini emas. Go'da vaqt bilan ishlash uchun ikkita asosiy tur bor: **`time.Time`** (aniq bir lahza — kalendar sanasi va soat) va **`time.Duration`** (vaqt oralig'i, ya'ni "necha vaqt davomida" — stopwatch ko'rsatgan narsa).

**`time.Duration` — vaqt oralig'i.** Bu, aslida, oddiy son (nanosekundlar sonida saqlanadi), lekin unga o'z nomi va qulay birliklari berilgan:

```go
besDaqiqa := 5 * time.Minute
ikkiSoat := 2 * time.Hour
uchYuzMillisekund := 300 * time.Millisecond
```

`time.Minute`, `time.Hour`, `time.Second` — bular tayyor konstantalar, ular bilan sonlarni ko'paytirib, kerakli davomiylikni hosil qilish mumkin.

**Duration'lar ustida amallar.** Duration'lar oddiy sonlardek qo'shiladi, ayiriladi, solishtiriladi:

```go
jami := 2*time.Hour + 30*time.Minute
fmt.Println(jami) // 2h30m0s

fmt.Println(jami.Minutes()) // 150 — umumiy daqiqalarga aylantirilgan (float64)
fmt.Println(jami.Hours())   // 2.5 — umumiy soatlarga aylantirilgan
```

Diqqat qiling: `jami` ni `fmt.Println` bilan chiqarganda, u avtomatik ravishda **o'qish uchun qulay** ko'rinishda (`2h30m0s`) chiqadi — chunki `time.Duration` o'zining maxsus `String()` methodiga ega (bu — "Custom Types" darsida ko'rgan, o'z turingizga method qo'shish g'oyasining aynan Go standart kutubxonasidagi haqiqiy namunasi).

**`time.Now()` va `time.Since()` — real vaqtni o'lchash.** Amalda, biror amal qancha vaqt olganini bilish uchun:

```go
boshlanish := time.Now()
// ... vaqt oluvchi ish ...
otganVaqt := time.Since(boshlanish) // time.Duration turida
fmt.Println("Ketgan vaqt:", otganVaqt)
```

Muhim ogohlantirish: `time.Now()` chaqirilgan payt haqiqiy, real dunyodagi soatga bog'liq — shuning uchun uni ishlatuvchi kod odatda **avtomatik testlarda ishlatilmaydi** (natija har safar, har kompyuterda boshqacha bo'ladi). Shu sababli, bu darsning mashqi haqiqiy vaqtga emas, balki `time.Duration` qiymatlarining o'zi bilan ishlashga qaratilgan — bu ko'proq uchraydigan, "berilgan davomiyliklarni hisoblash" turidagi vazifa.

## EXAMPLE

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	besDaqiqa := 5 * time.Minute
	ikkiSoat := 2 * time.Hour

	jami := besDaqiqa + ikkiSoat
	fmt.Println(jami)          // 2h5m0s
	fmt.Println(jami.Minutes()) // 125
}
```

Natija:

```
2h5m0s
125
```

## TASK

`jamiDavomiylik(daqiqalar []int) time.Duration` funksiyasi berilgan. Uni shunday to'ldiringki, u berilgan `daqiqalar` slice'idagi barcha sonlarni (har biri daqiqa sifatida) qo'shib, umumiy `time.Duration` qaytarsin (masalan, `jamiDavomiylik([]int{30, 45, 15})` → `1h30m0s`, ya'ni jami 90 daqiqa).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval barcha daqiqalarni oddiy `int` sifatida yig'ib oling: `jami := 0`, keyin `for _, d := range daqiqalar { jami += d }`.
2. Yig'ilgan `int`ni `time.Duration`ga aylantiring va `time.Minute`ga ko'paytiring: `return time.Duration(jami) * time.Minute`.
