# 01 — Hello World

## THEORY

Tasavvur qiling, siz katta bir zavodda ishlaysiz. Zavodda yuzlab bo'lim bor — har birida o'z vazifasi bilan shug'ullanadigan ishchilar guruhi. Lekin har qanday zavodning bitta narsasi bo'ladi: **bosh kirish eshigi**, ya'ni ishning boshlanadigan aniq nuqtasi. Go dasturi ham xuddi shunday ishlaydi.

**Paket (package) — bu bo'lim.** Go kodining har bir fayli qaysi "bo'lim"ga tegishli ekanligini fayl boshida ko'rsatishi shart:

```go
package main
```

`main` — bu oddiy bo'lim emas, balki maxsus nomlangan bo'lim: aynan shu nomdagi paket ichidagi kod **ishga tushiriladigan dastur** ekanligini bildiradi. Agar siz `package hisoblash` yoki `package vositalar` deb yozsangiz, bu — boshqa dasturlar foydalanadigan "yordamchi bo'lim" bo'lar edi, lekin uni to'g'ridan-to'g'ri ishga tushirib bo'lmaydi. Faqat `package main` — ishga tushiriladigan dastur ekanligining belgisi.

**`func main()` — bu bosh kirish eshigi.** `package main` ichida albatta `main()` nomli funksiya bo'lishi kerak — dastur ishga tushganda, Go operatsion tizimga: "qayerdan boshlashni bilaman, mana shu joydan boshla" deydi, va aynan shu `main()` funksiyasi ichidagi kod birinchi bo'lib ishlaydi. Zavodning har bir ishchisi o'z ish joyiga o'zi keladi, lekin ishning o'zi darvozadan boshlanadi — `main()` aynan shu darvoza.

```go
func main() {
	// dastur shu yerdan boshlanadi
}
```

**Ekranga chiqarish — bu ovoz balandkarligini yoqish kabi.** Dastur ichida hisoblab chiqarilgan narsani odam ko'rishi uchun uni ekranga (terminalga) chiqarish kerak. Buning uchun `fmt` nomli tayyor "asboblar qutisi" (paket) bor — u yerda matnni chiroyli formatlab chiqarish uchun tayyor vositalar mavjud. `fmt` — inglizcha "format" so'zidan qisqartma.

Eng ko'p ishlatiladigan vosita — `Println` ("print line", ya'ni "qatorni chiqar"):

```go
fmt.Println("Salom, Go!")
```

**Boshqa bo'limdan asbob olib kelish — bu `import`.** Zavodning o'zingiz ishlamaydigan bo'limidan biror asbobni ishlatmoqchi bo'lsangiz, avval o'sha bo'limdan ruxsat olib, uni o'z ish joyingizga olib kelishingiz kerak. Go'da bu — `import` deb ataladi: siz `fmt` bo'limidagi vositalardan foydalanmoqchi bo'lsangiz, avval uni faylning boshida "chaqirib olishingiz" kerak:

```go
import "fmt"
```

Buni unutsangiz, Go kompilyatori kodni umuman ishga tushirmaydi va aniq xato ko'rsatadi — bu Go'ning yaxshi tomoni: xatoni dastur ishga tushishidan oldin, hali yozish bosqichidayoq topib beradi.

**Izohlar (comments) — o'zingiz uchun eslatma.** Kod ichiga, kompyuter e'tiborsiz qoldiradigan, faqat odamlar uchun mo'ljallangan matn yozish mumkin. Bir qatorlik izoh `//` bilan boshlanadi:

```go
// Bu — izoh, Go uni umuman o'qimaydi
fmt.Println("Salom") // qatorning davomiga ham yozsa bo'ladi
```

Bir nechta qatorlik izoh esa `/* ... */` ichiga olinadi. Izohlar kodni tushunarli qilish uchun juda foydali, ayniqsa keyinroq o'zingiz yozgan kodga qaytganingizda.

**Butun dastur qanday ko'rinadi.** Endi hammasini birlashtiramiz: qaysi bo'limga tegishli ekanligini aytamiz (`package main`), kerakli asbobni chaqirib olamiz (`import "fmt"`), va darvozani ochamiz (`func main() { ... }`) — ichida esa xohlagan ishimizni qilamiz.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	// Bu dastur ishga tushganda birinchi bo'lib shu qator bajariladi
	fmt.Println("Salom, Go!")
	fmt.Println("Bu — ikkinchi qator")
}
```

Natija (terminalga chiqadigan matn, yuqoridan pastga, aynan shu tartibda):

```
Salom, Go!
Bu — ikkinchi qator
```

Diqqat qiling: `fmt.Println` har chaqirilganda **yangi qatordan** boshlab chiqaradi — xuddi daftarga har safar yangi satrdan yozgandek.

## TASK

1. `main()` funksiyasi ichida `fmt.Println` yordamida ekranga aynan `Hello, Go!` degan matnni chiqaring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Println("...")` qatorini `main()` funksiyasi ichiga yozing.
2. Qatorlar orasidagi qo'shtirnoq belgilariga e'tibor bering — matn aynan `Hello, Go!` bo'lishi kerak, katta-kichik harflar va tinish belgilari ham mos kelishi shart.
