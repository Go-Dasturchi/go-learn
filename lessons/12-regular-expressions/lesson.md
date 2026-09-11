# 12 — Regular Expressions

## THEORY

Tasavvur qiling, sizda katta bir qog'oz bor va siz undan faqat **ma'lum shakldagi** bo'laklarni (masalan, aynan doira shaklidagilarni) kesib olishingiz kerak — buning uchun maxsus shablon-kesuvchi (stencil) ishlatasiz. **Regular expression (regex)** — matn ichidan ma'lum "shakl"ga (naqshga) mos keladigan qismlarni topish uchun ishlatiladigan maxsus til. Go'da bu — `regexp` paketi orqali ishlatiladi.

**Oddiy tekshirish — `MatchString`.** Matn berilgan naqshga mos keladimi yo'qmi, shuni tekshirish uchun:

```go
mos, err := regexp.MatchString(`^[0-9]+$`, "12345")
fmt.Println(mos) // true — faqat raqamlardan iborat
```

Naqsh (pattern) — bu maxsus belgilar bilan yozilgan "qoida": `^` — matn boshini bildiradi, `[0-9]` — istalgan bitta raqam, `+` — undan oldingi narsa **bir yoki undan ko'p marta** takrorlanishi mumkinligini, `$` — matn oxirini bildiradi. Ya'ni `^[0-9]+$` — "boshidan oxirigacha, faqat raqamlardan iborat" degani.

**Naqshni oldindan "tayyorlash" — `regexp.MustCompile`.** Agar bir xil naqshni bir necha marta ishlatmoqchi bo'lsangiz, uni har safar qaytadan "tekshirib" o'tirmasdan, bir marta **kompilyatsiya qilib** olish samaraliroq:

```go
naqsh := regexp.MustCompile(`[0-9]+`)
```

`MustCompile` — agar naqsh noto'g'ri yozilgan bo'lsa, xatolikni qaytarish o'rniga darhol **panic** qiladi ("Panic and Recover" darsini eslang) — bu odatda naqshlar dastur kodida qat'iy, o'zgarmas (litereal) yozilganda ishlatiladi, chunki bunday holda naqshning noto'g'ri ekanligi darhol, dastur ishga tushishi bilanoq ma'lum bo'lishi kerak.

**Barcha mosliklarni topish — `FindAllString`.** Matn ichidan naqshga mos keladigan **barcha** qismlarni ro'yxat qilib olish uchun:

```go
naqsh := regexp.MustCompile(`[0-9]+`)
natija := naqsh.FindAllString("Uy 12, xona 5, qavat 3", -1)
fmt.Println(natija) // [12 5 3]
```

`-1` — "cheksiz, hammasini top" degani (agar musbat son bersangiz, faqat shuncha tagacha topadi).

**Boshqa foydali funksiyalar:**

```go
naqsh.MatchString("abc123")           // true/false — mos keladimi
naqsh.FindString("abc123")             // "123" — birinchi moslikni topadi
naqsh.ReplaceAllString("a1b2c3", "#") // "a#b#c#" — har bir moslikni almashtiradi
```

Regex — kuchli, lekin murakkab bo'lishi mumkin bo'lgan vosita; oddiy holatlar (ma'lum belgi turkumi, takrorlanish) uchun juda qulay, lekin murakkab matn tuzilmalarini tahlil qilish uchun ba'zan boshqa yondashuvlar (masalan, `strings` paketi funksiyalari) soddaroq bo'lishi mumkin.

## EXAMPLE

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	naqsh := regexp.MustCompile(`[0-9]+`)

	fmt.Println(naqsh.MatchString("xona 5"))
	fmt.Println(naqsh.FindAllString("Uy 12, xona 5, qavat 3", -1))
	fmt.Println(naqsh.ReplaceAllString("a1b22c333", "#"))
}
```

Natija:

```
true
[12 5 3]
a#b#c#
```

## TASK

`raqamlarniTop(matn string) []string` funksiyasi berilgan. Uni `regexp` paketi yordamida shunday to'ldiringki, u berilgan `matn` ichidagi barcha raqamlar ketma-ketligini (bir yoki undan ko'p ketma-ket raqam) topib, slice qilib qaytarsin (masalan, `raqamlarniTop("Uy 12, xona 5")` → `["12", "5"]`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `regexp.MustCompile` funksiyasiga `[0-9]+` naqshini bering — bu "bir yoki undan ko'p ketma-ket raqam"ni bildiradi.
2. `return naqsh.FindAllString(matn, -1)` — barcha mosliklarni topib qaytaring.
