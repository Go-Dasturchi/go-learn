# 15 — Exported / Unexported

## THEORY

Bankka kirganingizda, mijozlar zali **hammaga ochiq** — istalgan kishi kirib, kassaga murojaat qila oladi. Lekin bankning **omonat xonasi** (seyf turgan joy) — faqat bank xodimlariga ochiq, tashqaridagilar u yerga kira olmaydi. Go'da paket ichidagi narsalar ham xuddi shunday ikki toifaga bo'linadi — bu, "Functions" va "Structs" darslarida qisqacha eslatilgan qoidaning, endi to'liq tushuntirilishi.

**Qoida — juda oddiy: birinchi harf.** Agar nom (funksiya, tur, struct maydoni, konstanta — **hammasi**ga tegishli) **katta harf** bilan boshlansa — u **eksport qilingan (exported)**, ya'ni boshqa paketlardan ham ko'rinadi va ishlatilishi mumkin. Kichik harf bilan boshlansa — **eksport qilinmagan (unexported)**, faqat **shu paket ichida** ko'rinadi:

```go
package hisoblash

func Kop(a, b int) int { // katta harf — Exported, boshqa paketlar chaqira oladi
	return yordamchi(a, b)
}

func yordamchi(a, b int) int { // kichik harf — unexported, faqat shu paket ichida
	return a * b
}
```

Boshqa paketdan `hisoblash.Kop(3, 4)` deb chaqirish mumkin, lekin `hisoblash.yordamchi(3, 4)` deb chaqirishga **urinib ko'rilsa, compile xatosi** chiqadi — Go bunga umuman yo'l qo'ymaydi.

**Bu qoida struct maydonlariga ham tegishli.** "Structs" darsida ko'rgan `Ism`, `Yosh` kabi katta harfli maydonlar — aynan shu sababdan katta harf bilan yozilgan edi: ular boshqa paketlardan ham o'qilishi/yozilishi mumkin bo'lishi kerak edi.

**Nega bu foydali — kapsulyatsiya (encapsulation).** Bank seyfi tashqariga yopiq bo'lgani kabi, paketning **ichki tafsilotlari** (masalan, qandaydir hisoblash formulasi, yoki vaqtinchalik holat) ni tashqi dunyodan **yashirish** — kodni ancha xavfsizroq va boshqarish osonroq qiladi. Agar biror narsa unexported bo'lsa, siz uni istalgan vaqtda, boshqa hech kimga ta'sir qilmasdan, o'zgartirishingiz mumkin — chunki uni **faqat siz**, shu paket ichida ishlatasiz.

**Amaliy naqsh — unexported maydon + exported method (getter).** Ko'p hollarda, struct'ning maydoni **to'g'ridan-to'g'ri** o'zgartirilishini xohlamaysiz (masalan, balansni tashqaridan ixtiyoriy songa "sozlab qo'yish" xavfli) — buning uchun maydonni **unexported** qilib, unga faqat maxsus, **exported** method orqali (masalan, o'qish uchun) murojaat qilinadi:

```go
type Hisob struct {
	balans int // unexported — tashqaridan to'g'ridan-to'g'ri o'zgartirib bo'lmaydi
}

func YangiHisob(boshlangich int) *Hisob {
	return &Hisob{balans: boshlangich}
}

func (h *Hisob) Balans() int { // exported "getter" — faqat o'qish uchun ochiq yo'l
	return h.balans
}
```

Boshqa paketdan `hisob.balans` deb to'g'ridan-to'g'ri murojaat qilib bo'lmaydi (compile xatosi), lekin `hisob.Balans()` orqali qiymatni **nazoratli** tarzda olish mumkin — bu, "Methods" darsida ko'rgan g'oyaning, kapsulyatsiya bilan birlashgan holati.

## EXAMPLE

```go
package main

import "fmt"

type Hisob struct {
	balans int
}

func YangiHisob(boshlangich int) *Hisob {
	return &Hisob{balans: boshlangich}
}

func (h *Hisob) Balans() int {
	return h.balans
}

func main() {
	h := YangiHisob(1000)
	fmt.Println(h.Balans())
}
```

Natija:

```
1000
```

## TASK

`Hisob` struct'i (`balans int` — unexported maydon) va `YangiHisob` funksiyasi berilgan. `Balans() int` methodini shunday to'ldiringki, u unexported `balans` maydonining qiymatini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Metod ichida `h` — struct'ning o'ziga (pointer orqali) ishora qiladi, uning ichidagi maydonlarga `h.balans` orqali murojaat qilinadi.
2. `return h.balans` — funksiya tanasi shu bitta qatordan iborat bo'lishi mumkin.
