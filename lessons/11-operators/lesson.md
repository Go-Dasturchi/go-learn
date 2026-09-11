# 11 — Operators

## THEORY

Siz allaqachon arifmetik (`+ - * / %`) va mantiqiy (`&& || !`) amallarni ko'rgansiz. Bundan tashqari Go'da yana bir necha muhim operator guruhi bor.

**Solishtirish operatorlari** (natijasi har doim `bool`):

```go
==   // teng
!=   // teng emas
<    // kichik
>    // katta
<=   // kichik yoki teng
>=   // katta yoki teng
```

**Amallar bajarilish tartibi (precedence).** Xuddi maktabda o'rgangan matematikadagidek — avval ko'paytirish/bo'lish, keyin qo'shish/ayirish bajariladi:

```go
natija := 2 + 3*4   // 14, 20 emas! avval 3*4=12, keyin 2+12=14
```

Agar boshqacha tartibda hisoblanishini xohlasangiz, qavslardan foydalaning — qavs ichidagi amal har doim birinchi bajariladi:

```go
natija := (2 + 3) * 4  // 20
```

Aniq bo'lmagan hollarda, hatto shart bo'lmasa ham, qavs qo'yish kodni o'qishni osonlashtiradi. Solishtirish operatorlari (`==`, `<` va h.k.) esa arifmetik amallardan **keyin** bajariladi:

```go
5 + 3 > 2*3   // avval 5+3=8 va 2*3=6 hisoblanadi, keyin 8 > 6 solishtiriladi → true
```

**Muhim ogohlantirish: zanjirlab solishtirish ishlamaydi.** Ba'zi tillarda `1 < x < 10` deb yozish mumkin, lekin Go'da bu compile bo'lmaydi — har bir solishtirishni alohida yozib, `&&` bilan bog'lash kerak:

```go
// x < 10  // XATO emas, lekin 1 < x < 10 kabi zanjir Go'da yo'q
agar1 := x > 1 && x < 10   // TO'G'RI usul
```

**Qisqartma operatorlar** — o'zgaruvchini o'zgartirib, qayta yozishning qisqa yo'li. Buni "hisobingizga pul qo'shish/ayirish" kabi tasavvur qiling — har safar "hisobim = hisobim + ..." deb to'liq yozish o'rniga, qisqa yo'l ishlatiladi:

```go
son := 5
son += 3   // son = son + 3  →  8
son -= 2   // son = son - 2  →  6
son *= 2   // son = son * 2  →  12
son /= 4   // son = son / 4  →  3
son %= 2   // son = son % 2  →  1
son++      // son = son + 1  →  2
son--      // son = son - 1  →  1
```

`++` va `--` faqat mustaqil qator sifatida ishlatiladi (`x := son++` kabi ifoda ichida ishlatib bo'lmaydi) — bu ko'plab boshqa tillardan farq qiladigan joyi. Bundan tashqari, Go'da faqat `son++` bor, `++son` (qiymatni oshirib, keyin ishlatish) shakli umuman yo'q — bu ba'zi tillarda ikkalasi ham bo'ladigan chalkashlikning oldini oladi.

**Bitli (bitwise) operatorlar — qisqacha tanishuv.** Go'da sonlarning ikkilik (binary) ko'rinishi ustida ishlaydigan operatorlar ham bor. Bular kundalik dasturlashda kamroq uchraydi, lekin bilib qo'yish foydali:

```go
a & b   // AND — ikkala bitning ham 1 bo'lgan joylarini qoldiradi
a | b   // OR  — kamida bittasi 1 bo'lgan joylarni qoldiradi
a ^ b   // XOR — faqat bittasi 1 bo'lgan joylarni qoldiradi
a << 1  // chapga siljitish — 2 ga ko'paytirishga teng
a >> 1  // o'ngga siljitish — 2 ga bo'lishga teng
```

## EXAMPLE

```go
package main

import "fmt"

func main() {
	balans := 100
	balans -= 30
	balans += 10

	fmt.Println(balans)          // 80
	fmt.Println(balans >= 50)    // true

	natija := 2 + 3*4
	fmt.Println(natija)          // 14

	yosh := 25
	fmt.Println(yosh > 18 && yosh < 65) // true — ikkala shart ham to'g'ri
}
```

Natija:

```
80
true
14
true
```

## TASK

`hisob` nomli o'zgaruvchi `10` qiymati bilan berilgan. Quyidagilarni ketma-ket bajaring:

1. `+=` yordamida unga `5` qo'shing.
2. `*=` yordamida uni `2` ga ko'paytiring.

Natija ekranga `30` deb chiqishi kerak (`fmt.Println(hisob)` allaqachon yozilgan).

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `hisob += 5` keyin `hisob *= 2` — ikkita alohida qatorda, aynan shu tartibda.
2. Tartib muhim: avval qo'shish, keyin ko'paytirish — aks holda natija boshqacha chiqadi.
