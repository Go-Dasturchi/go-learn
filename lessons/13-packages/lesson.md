# 13 — Packages

## THEORY

Katta kutubxonani tasavvur qiling — kitoblar tasodifiy tartibsiz uyilib yotmaydi, balki mavzu bo'yicha bo'limlarga ("tarix", "matematika", "badiiy adabiyot") joylashtirilgan. Go dasturi ham xuddi shunday — kod tasodifiy fayllar to'plami emas, balki **paketlar (packages)**ga bo'lingan bo'ladi, va har bir paket ma'lum bir vazifaga mas'ul bo'ladi.

**Har bir Go fayli — bitta paketga tegishli.** Fayl boshida yozilgan `package nomi` qatori, shu fayl qaysi "bo'lim"ga tegishli ekanligini bildiradi ("Hello World" darsida ko'rgan `package main`ni eslang):

```go
package hisoblash

func Kop(a, b int) int {
	return a * b
}
```

Bu — `hisoblash` nomli paket, va u boshqa fayllardan **import** qilinishi mumkin. Bitta papkadagi **barcha** `.go` fayllari, agar bir xil `package` qatoriga ega bo'lsa, **bitta paketning qismlari** hisoblanadi — ular bir-birining funksiyalariga, hatto alohida import qilmasdan ham, to'g'ridan-to'g'ri murojaat qila oladi.

**Real loyihada fayllar qanday tuzilishi mumkin:**

```
loyiha/
├── main.go              (package main)
└── hisoblash/
    └── hisoblash.go     (package hisoblash)
```

`main.go` ichida, `hisoblash` paketidan foydalanish uchun uni import qilish kerak (import yo'li odatda modul nomi + papka nomidan iborat bo'ladi — "Modules" darsida bu haqda batafsil gaplashamiz):

```go
import "loyiha/hisoblash"

func main() {
	fmt.Println(hisoblash.Kop(3, 4)) // paket nomi orqali murojaat
}
```

**Nega bu kursda hamma narsa bitta faylda edi.** Diqqat qilgan bo'lsangiz, shu paytgacha barcha mashqlar bitta `package main` ichida, bitta faylda bo'lgan — bu shunchaki, har bir darsni **mustaqil va sodda** qilib saqlash uchun qilingan qulaylik. Haqiqiy, katta loyihalarda esa kod har doim mantiqiy jihatdan bir nechta paketga bo'linadi — masalan, aynan shu kursning o'zi (`go-learn` dasturi) ham `internal/cli`, `internal/lesson`, `internal/runner` kabi alohida paketlarga bo'lingan holda yozilgan.

**Nega paketlarga bo'lish foydali:**
1. **Tashkiliy tozalik** — bog'liq kod bir joyda, aloqasiz kod boshqa joyda turadi.
2. **Qayta ishlatish** — bir paketni bir nechta boshqa loyihada import qilib ishlatish mumkin.
3. **Kapsulyatsiya** — "Exported/Unexported" darsida ko'radigan naqsh orqali, paket faqat kerakli qismini "ochiq" qiladi, ichki tafsilotlarni yashiradi.

**Paket darajasidagi e'lonlar.** "Scope" darsida ko'rgan narsani eslang — funksiyadan tashqarida, fayl boshida e'lon qilingan `const`, `var`, yoki `func` — **butun paket bo'ylab**, hatto bir xil paketning boshqa fayllarida ham ko'rinadi:

```go
package main

const versiya = "1.0.0" // paket darajasidagi konstanta — butun paket bo'ylab ko'rinadi

func versiyaChiqar() string {
	return "Versiya: " + versiya
}
```

## EXAMPLE

```go
package main

import "fmt"

const versiya = "1.0.0"

func versiyaChiqar() string {
	return "Versiya: " + versiya
}

func main() {
	fmt.Println(versiyaChiqar())
}
```

Natija:

```
Versiya: 1.0.0
```

## TASK

`versiya` nomli paket darajasidagi konstanta (`"2.5.0"` qiymati bilan, allaqachon yozilgan) berilgan. `versiyaChiqar() string` funksiyasini shunday to'ldiringki, u `"Versiya: <versiya>"` (masalan, `"Versiya: 2.5.0"`) matnini qaytarsin, shu paket darajasidagi `versiya` konstantasidan foydalanib.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `versiya` allaqachon paket darajasida e'lon qilingan — funksiya ichida uni parametr sifatida qabul qilishning hojati yo'q, to'g'ridan-to'g'ri ishlatish mumkin.
2. `return "Versiya: " + versiya` — ikkita stringni `+` bilan birlashtiring.
