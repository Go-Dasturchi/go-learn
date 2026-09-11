# 25 — Structs

## THEORY

Pasportingizni tasavvur qiling — unda "Ism", "Familiya", "Tug'ilgan yil" kabi turli-tuman ma'lumotlar bitta hujjat ichida, har biri o'z yorlig'i bilan joylashgan. Go'da bunday "bir nechta turli ma'lumotni bitta nom ostida birlashtirish" — **struct** deb ataladi.

**E'lon qilish.** `type` kalit so'zi bilan o'zingizning yangi turingizni yaratasiz:

```go
type Odam struct {
	Ism  string
	Yosh int
}
```

Bu — "Odam degan yangi tur bor, u ichida `Ism` (matn) va `Yosh` (butun son) degan ikkita maydon (field) saqlaydi" degani. Odatda struct maydonlari **katta harf** bilan yoziladi (`Ism`, `Yosh`) — bu boshqa fayllardan ham ko'rinishi (eksport qilinishi) mumkinligini bildiradi (06-darsda funksiyalar haqida gaplashganimizdek), va Go dasturchilari orasida struct maydonlari uchun keng tarqalgan odat.

**Struct qiymati yaratish (literal) — ikki usul bor:**

```go
// 1. Nomlangan maydonlar bilan — tavsiya etiladi, o'qishga aniqroq
ali := Odam{Ism: "Ali", Yosh: 25}

// 2. Faqat qiymatlar bilan, e'lon qilingan tartibda — qisqaroq, lekin xato qilish osonroq
vali := Odam{"Vali", 30}
```

Nomlangan usul afzalroq, chunki agar kelajakda struct'ga yangi maydon qo'shsangiz yoki maydonlar tartibini o'zgartirsangiz, "faqat qiymatlar" usuli bilan yozilgan kod kutilmagan tarzda buzilishi mumkin, nomlangan usul esa har doim to'g'ri ishlaydi.

**Maydonlarga murojaat — nuqta (`.`) orqali:**

```go
fmt.Println(ali.Ism)  // "Ali"
fmt.Println(ali.Yosh) // 25

ali.Yosh = 26 // maydonni yangilash
```

**Standart (zero) qiymatli struct.** Agar struct'ni hech qanday qiymat bermay yaratsangiz, uning har bir maydoni o'z turining standart qiymatiga ega bo'ladi (xuddi "Variables" darsida ko'rgan `var` kabi):

```go
var bosh Odam
fmt.Println(bosh) // {  0} — Ism bo'sh string, Yosh esa 0
```

**Muhim: struct — "qiymat turi" (value type), xuddi array kabi.** "Arrays" darsida array'ni tenglashtirish uning **to'liq nusxasini** yaratishini ko'rgan edik. Struct ham xuddi shunday ishlaydi:

```go
a := Odam{Ism: "Ali", Yosh: 25}
b := a         // b — a ning TO'LIQ, mustaqil nusxasi
b.Yosh = 99
fmt.Println(a.Yosh) // 25 — a o'zgarmagan!
fmt.Println(b.Yosh) // 99 — faqat b o'zgargan
```

Bu — map va slice'dan (ular "yorliq" kabi ishlaydi, "Slice Copy" darsida ko'rganingizdek) muhim farq: struct'ni tenglashtirish yoki funksiyaga argument sifatida berish, har doim **butun nusxasini** yaratadi. (Keyingi darslarda ko'radigan **pointer**lar aynan shu xususiyatni "aylanib o'tish" — ya'ni struct'ning nusxasini emas, aynan o'zini o'zgartirish — uchun ishlatiladi.)

## EXAMPLE

```go
package main

import "fmt"

type Odam struct {
	Ism  string
	Yosh int
}

func main() {
	ali := Odam{Ism: "Ali", Yosh: 25}
	vali := Odam{Ism: "Vali", Yosh: 30}

	fmt.Println(ali.Ism, ali.Yosh)
	fmt.Println(vali)

	nusxa := ali
	nusxa.Yosh = 100
	fmt.Println(ali.Yosh)   // 25 — o'zgarmagan
	fmt.Println(nusxa.Yosh) // 100
}
```

Natija:

```
Ali 25
{Vali 30}
25
100
```

## TASK

`Odam` struct'i (`Ism string`, `Yosh int` maydonlari bilan) va `kattaroq(a, b Odam) Odam` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u ikkita `Odam`dan **yoshi kattarog'ini** qaytarsin (agar yoshlari teng bo'lsa, birinchisini — `a`ni — qaytaring).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Maydonga nuqta orqali murojaat qiling: `a.Yosh`, `b.Yosh`.
2. `if a.Yosh >= b.Yosh { return a }` keyin `return b` — teng bo'lgan holatni ham `>=` o'zi to'g'ri hal qiladi.
