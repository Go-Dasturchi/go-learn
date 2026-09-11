# 21 — Slice Append

## THEORY

Har biringiz daftarga yozib borasiz — sahifa tugasa, yangi sahifa qo'shiladi, daftar "o'sadi". Go'da slice'ga yangi element qo'shishning rasmiy usuli — `append` degan tayyor (built-in) funksiya:

```go
mevalar := []string{"olma", "banan"}
mevalar = append(mevalar, "uzum")
fmt.Println(mevalar) // [olma banan uzum]
```

**Eng muhim qoida: `append` natijasini albatta qayta o'zgaruvchiga yozib qo'yish kerak.** `append(mevalar, "uzum")` — bu `mevalar`ni o'zgartirmaydi, balki **yangi** (yoki kengaytirilgan) slice qaytaradi. Agar natijani `mevalar = ...` deb qayta yozmasangiz, qo'shilgan element "yo'qolib qoladi":

```go
append(mevalar, "shaftoli") // XATO: natija hech qayerga yozilmadi, shaftoli "yo'qoladi"
mevalar = append(mevalar, "shaftoli") // TO'G'RI
```

**Bir nechta elementni birdaniga qo'shish:**

```go
mevalar = append(mevalar, "nok", "gilos", "olcha")
```

**Bir slice'ni boshqasiga qo'shish — `...` yoyish (spread) belgisi bilan:**

```go
yangi := []string{"anor", "hurmo"}
mevalar = append(mevalar, yangi...) // yangi ichidagi HAR BIR elementni alohida qo'shadi
```

Agar `...` ni tashlab ketsangiz, Go butun `yangi` slice'ini bitta element sifatida qo'shishga urinadi va bu compile xatosiga sabab bo'ladi (`yangi`ning turi `[]string`, `mevalar`ning har bir elementi esa `string` — ular mos kelmaydi).

**Nima uchun bo'sh slice'ga append qilish ishlaydi.** Yodingizda bo'lsa, "Slices" darsida nil slice'ga to'g'ridan-to'g'ri indeks orqali yozish (`royxat[0] = 5`) xato berishini aytgan edik. Lekin `append` bunda **butunlay boshqacha** ishlaydi — u nil (yoki bo'sh) slice bilan ham muammosiz ishlaydi, chunki kerak bo'lganda **o'zi yangi joy ajratadi**:

```go
var royxat []int          // nil slice
royxat = append(royxat, 1) // ishlaydi! royxat endi [1]
```

Shu sababli, tsikl ichida elementlarni to'plab boradigan slice'ni odatda bo'sh e'lon qilib (`royxat := []int{}` yoki hatto `var royxat []int`), keyin `append` bilan to'ldirib boriladi — bu Go'da eng ko'p uchraydigan naqshlardan biri.

**Sig'im (capacity) — sahifalarni oldindan qoldirib qo'yish.** Slice ichida, uzunlikdan (`len`) tashqari, yana **sig'im** (`cap`) degan narsa ham bor — bu, aslida, orqa fonda ajratilgan joyning umumiy hajmi. Daftarni tasavvur qiling: ba'zan bir necha bo'sh sahifa oldindan qoldirilgan bo'ladi — yangi yozuv qo'shganingizda, agar bo'sh sahifa bo'lsa, o'sha joyning o'ziga yoziladi (tez); agar bo'sh joy qolmagan bo'lsa, Go **butunlay yangi, kattaroq** massiv ajratadi va eski ma'lumotlarning hammasini u yerga ko'chiradi (sekinroq). Bu jarayon avtomatik va shaffof — siz buni ko'rmaysiz, `append` har doim to'g'ri natija qaytaradi, lekin nima uchun katta hajmdagi ma'lumotlarga ko'p marta `append` qilish ba'zan boshqa usullardan sekinroq bo'lishi mumkinligini tushunish uchun buni bilib qo'yish foydali.

## EXAMPLE

```go
package main

import "fmt"

func main() {
	mevalar := []string{"olma", "banan"}
	mevalar = append(mevalar, "uzum")
	fmt.Println(mevalar) // [olma banan uzum]

	mevalar = append(mevalar, "nok", "gilos")
	fmt.Println(mevalar) // [olma banan uzum nok gilos]

	qoshimcha := []string{"anor", "hurmo"}
	mevalar = append(mevalar, qoshimcha...)
	fmt.Println(mevalar) // [olma banan uzum nok gilos anor hurmo]

	var royxat []int
	for i := 1; i <= 5; i++ {
		royxat = append(royxat, i*i)
	}
	fmt.Println(royxat) // [1 4 9 16 25]
}
```

Natija:

```
[olma banan uzum]
[olma banan uzum nok gilos]
[olma banan uzum nok gilos anor hurmo]
[1 4 9 16 25]
```

## TASK

`juftSonlar(n int) []int` funksiyasi berilgan. Uni `append` yordamida shunday to'ldiringki, u `1` dan `n` gacha bo'lgan barcha **juft** sonlarni o'z ichiga olgan slice qaytarsin (masalan, `juftSonlar(10)` → `[2 4 6 8 10]`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh slice bilan boshlang: `natija := []int{}`, keyin `1` dan `n` gacha `for` tsikli yozing.
2. Har bir son uchun `son%2 == 0` shartini tekshirib, to'g'ri bo'lsa `natija = append(natija, son)` qiling.
