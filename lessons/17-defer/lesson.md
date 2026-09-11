# 17 — Defer

## THEORY

Uydan chiqishdan oldin eshikka "chiroqni o'chir!" degan qog'oz yopishtirib qo'yganingizni tasavvur qiling. Siz bu eslatmani **hozir** yozasiz, lekin u aslida **eshikdan chiqayotganingizda** amalga oshadi — xona ichida hali qancha ish qilishingizdan qat'iy nazar. Go'da `defer` aynan shunday ishlaydi: `defer` bilan belgilangan chaqiruv **darhol emas**, balki uni o'z ichiga olgan funksiya **tugashidan oldin** ishga tushadi — funksiyadan qanday chiqilishidan qat'iy nazar (oddiy `return` orqalimi, yoki xato tufaylimi):

```go
func salomlash() {
	defer fmt.Println("Xayr!")
	fmt.Println("Salom!")
}
```

Bu funksiya chaqirilganda avval `"Salom!"`, keyin `"Xayr!"` chop etiladi — garchi `defer` qatori kodda birinchi bo'lib yozilgan bo'lsa ham.

`defer` eng ko'p **tozalash (cleanup)** ishlari uchun ishlatiladi — fayl yopish, ulanishni uzish, qulfni bo'shatish. Buning sababi oddiy: siz resursni ochgan zahoti, uni "qanday yopishni" darhol, o'sha yerning o'zida yozib qo'yasiz — shunda uni yopishni unutib qo'yish xavfi yo'qoladi, hatto funksiya o'rtasida kutilmagan joydan chiqib ketsa ham:

```go
fayl, err := os.Open("data.txt")
if err != nil {
	return err
}
defer fayl.Close()
// ... fayl bilan ishlash ...
// fayl.Close() funksiya qanday tugashidan qat'iy nazar avtomatik chaqiriladi
```

**Bir nechta `defer` — teskari tartibda (LIFO).** Agar bir nechta `defer` yozilsa, ular **teskari tartibda** (LIFO — "oxirgi kirgan birinchi chiqadi") bajariladi — xuddi eshikka bir nechta eslatma qog'ozini ketma-ket yopishtirsangiz, chiqayotganda **eng oxirgi yopishtirgan qog'ozni birinchi** o'qiganingizga o'xshaydi:

```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
// chiqadi: 3, 2, 1
```

**Muhim, ko'p odam adashadigan nozik joy: `defer`ning argumentlari darhol hisoblanadi.** `defer` faqat **chaqiruvning o'zini** kechiktiradi, lekin unga beriladigan qiymatlar (argumentlar) `defer` qatori ishga tushgan zahoti, **o'sha lahzadagi** holatda "suratga olib" qo'yiladi:

```go
son := 1
defer fmt.Println("Kechiktirilgan qiymat:", son) // son=1 ni HOZIR "suratga oladi"
son = 99
fmt.Println("Joriy qiymat:", son)
// chiqadi:
// Joriy qiymat: 99
// Kechiktirilgan qiymat: 1     <- 99 emas, 1! chunki son defer chaqirilgan paytda 1 edi
```

Buni eslatma qog'oziga o'xshatish mumkin: qog'ozga "chiroqni o'chir, hozir soat 3" deb yozib qo'ysangiz, keyinroq soat o'zgarsa ham, qog'ozdagi yozuv **o'zgarmaydi** — chunki u yozilgan paytdagi holatni saqlab qolgan. Shu sababli, agar sizga funksiya **tugashi vaqtidagi eng oxirgi** qiymat kerak bo'lsa (masalan, named return qiymatini o'zgartirish uchun — buni oldingi darsda eslatgan edik), oddiy qiymat emas, balki funksiya (`func() { ... }`) ni `defer` qilish kerak bo'ladi — bu ancha ilg'or mavzu, hozircha shuni bilib qo'yish kifoya.

## EXAMPLE

```go
package main

import "fmt"

func ishla() {
	defer fmt.Println("Tozalash bajarildi")
	fmt.Println("Ish boshlandi")
	fmt.Println("Ish tugadi")
}

func sanoq() {
	defer fmt.Println("1-eslatma")
	defer fmt.Println("2-eslatma")
	defer fmt.Println("3-eslatma")
	fmt.Println("Asosiy ish")
}

func main() {
	ishla()
	fmt.Println("---")
	sanoq()
}
```

Natija:

```
Ish boshlandi
Ish tugadi
Tozalash bajarildi
---
Asosiy ish
3-eslatma
2-eslatma
1-eslatma
```

## TASK

`hisobla()` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `defer` yordamida `"Funksiya tugadi"` so'zini chop etuvchi `fmt.Println` chaqiruvini kechiktiring (funksiya boshida yozilsa ham, oxirida chiqishi kerak).
2. Undan keyin oddiy (kechiktirilmagan) `fmt.Println("Hisoblash boshlandi")` yozing.

Kutilgan natija tartibi: avval `"Hisoblash boshlandi"`, keyin `"Funksiya tugadi"`.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `defer fmt.Println("Funksiya tugadi")` qatorini funksiya boshiga yozing.
2. `defer` qatoridan keyin oddiy `fmt.Println("Hisoblash boshlandi")` yozing — tartib muhim: `defer` avval yoziladi, lekin oxirida ishga tushadi.
