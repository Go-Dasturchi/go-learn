# 07 — Constants

## THEORY

`var` va `:=` bilan yaratilgan o'zgaruvchi — yorliqlangan tortma, ichidagi narsani istalgan payt almashtirsa bo'ladi. Lekin ba'zi qiymatlar hech qachon o'zgarmasligi kerak — masalan, doiraning `pi` soni, haftadagi kunlar soni, yoki bankdagi eng past foiz stavkasi. Bunday qiymatlarni **zavodda quyilgan tanga** kabi tasavvur qiling: bir marta zarb qilingandan keyin, uni endi qayta shakllantirib bo'lmaydi. Go'da bu — `const`:

```go
const pi = 3.14
```

`const` bilan e'lon qilingan qiymatni keyin o'zgartirishga urinsangiz, dastur **compile bo'lmaydi** — bu xatoni dastur ishga tushishidan oldin, hali yozish bosqichidayoq ushlab qolish imkonini beradi. Bu ayniqsa katta dasturlarda foydali: agar kimdir tasodifan "o'zgarmas bo'lishi kerak" qiymatni o'zgartirib qo'ysa, Go buni darrov xato deb ko'rsatadi.

**Bir nechta doimiyni birga e'lon qilish:**

```go
const (
	Kichik = 1
	Katta  = 100
)
```

**Muhim farq:** `var` singari, `const` ham funksiyadan tashqarida (paket darajasida) yoki ichida ishlatilishi mumkin — lekin `:=` konstantalar uchun ishlamaydi, faqat `const` kaliti orqali e'lon qilinadi.

**Konstanta qiymati kompilyatsiya vaqtida ma'lum bo'lishi kerak.** Bu — muhim farqlovchi jihat: `var` bilan o'zgaruvchiga funksiya natijasi yoki hisoblangan biror narsani berish mumkin (masalan `var x = hisobla()`), lekin `const` faqat dastur yozilayotgan paytdayoq aniq bo'lgan qiymatlarni qabul qiladi (sonlar, matnlar, boshqa konstantalar bilan qilingan oddiy arifmetika). `const soat = vaqtOl()` kabi yozish — xato beradi, chunki `vaqtOl()` natijasi faqat dastur ishlaganda ma'lum bo'ladi.

**`iota` — avtomatik raqamlash.** Go'da ketma-ket nomerlangan konstantalar yaratish uchun maxsus, juda qulay vosita bor — `iota`. U `const ( ... )` blokining ichida har qatorda avtomatik ravishda `0`dan boshlab bittadan oshib boradi — xuddi navbat mashinasi har mijozga avtomatik ketma-ket raqam bosib berganidek:

```go
const (
	Dushanba = iota // 0
	Seshanba         // 1
	Chorshanba       // 2
	Payshanba        // 3
	Juma             // 4
)
```

Bu yerda faqat birinchi qatorga `= iota` yozilgan, qolganlari esa uni avtomatik meros qiladi va navbatma-navbat oshib boradi. `iota` ayniqsa "shu ro'yxatdan bittasi" turidagi qiymatlarni (masalan, hafta kunlari, holat darajalari — "faol", "kutilmoqda", "bekor qilingan") ifodalashda juda qulay, chunki har biriga qo'lda raqam yozib o'tirishning hojati yo'q.

## EXAMPLE

```go
package main

import "fmt"

const soatlarKuni = 24

const (
	Kichik = iota // 0
	Orta          // 1
	Katta         // 2
)

func main() {
	fmt.Println(soatlarKuni)
	fmt.Println(Kichik, Orta, Katta)
}
```

Natija:

```
24
0 1 2
```

## TASK

1. `haftaKunlari` nomli konstanta yarating, qiymati `7` bo'lsin.
2. Uni `fmt.Println` orqali ekranga chiqaring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `const haftaKunlari = 7` qatorini `main()` ichida yozing.
2. Konstanta ham oddiy o'zgaruvchidek `fmt.Println(haftaKunlari)` orqali chiqariladi — faqat uni keyin qayta o'zgartirib bo'lmaydi.
