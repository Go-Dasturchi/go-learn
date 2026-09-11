# 01 — Errors

## THEORY

Bankka borib pul yechmoqchi bo'lganingizni, lekin hisobingizda yetarli mablag' yo'qligini tasavvur qiling. Kassir sizga pulni bermaydi, aksincha bir varaqcha yozib beradi: "Yetarli mablag' yo'q". Bu — muvaffaqiyatsizlik, lekin dastur qulab tushmadi, sizga **nima uchun** ishlamaganini tushuntirib berildi. Go'da bunday "muvaffaqiyatsizlik haqidagi xabar" — **error** deb ataladi.

**Go'da xatoliklar qanday ishlaydi.** Ko'p tillarda xatolik yuz berganda dastur "otib yuboriladi" (throw/catch mexanizmi). Go'da bunday emas — xatolik oddiy **qiymat**, va u odatda funksiyaning **oxirgi qaytariladigan qiymati** sifatida beriladi ("Multiple Return Values" darsida ko'rgan naqshni eslang):

```go
func bolish(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nolga bo'lib bo'lmaydi")
	}
	return a / b, nil
}
```

`errors.New("...")` — matndan oddiy xatolik yaratadi. Xatolik yo'q bo'lsa, `nil` qaytariladi — bu "hammasi joyida" degani.

**Xatolikni tekshirish — Go'ning eng ko'p uchraydigan naqshi.** Har safar xato qaytarishi mumkin bo'lgan funksiyani chaqirganda, natijadan foydalanishdan oldin xatolikni tekshirish odat tusiga kirgan:

```go
natija, err := bolish(10, 0)
if err != nil {
	fmt.Println("Xato:", err)
	return
}
fmt.Println(natija)
```

Bu naqsh shunchalik ko'p uchraydiki, Go dasturchilari buni deyarli avtomatik ravishda yozadi — chunki xatolikni e'tiborsiz qoldirib, natijadan foydalanishga urinish (masalan, yuqoridagi holatda `b == 0` bo'lganida `natija = 0` bilan ishlashda davom etish) kutilmagan xatoliklarga olib kelishi mumkin.

**`error` — bu shunchaki interfeys.** "Interfaces" darsida ko'rgan tushunchani eslang: `error` — Go'da o'rnatilgan, juda oddiy interfeys:

```go
type error interface {
	Error() string
}
```

Ya'ni, `Error() string` methodiga ega bo'lgan **har qanday** tur `error` sifatida ishlatilishi mumkin — bu keyingi darsda ("Custom Errors") o'zimizning maxsus xatolik turlarimizni yaratish imkonini beradi.

**`fmt.Errorf` — formatlangan xatolik yaratish.** Ko'pincha xatolik matni ichida qandaydir qiymatni ham ko'rsatish kerak bo'ladi — buning uchun `errors.New` o'rniga `fmt.Errorf` qulayroq, chunki u `fmt.Sprintf` kabi joy egallovchilarni (`%d`, `%s` va h.k.) qo'llab-quvvatlaydi:

```go
func yoshTekshir(yosh int) error {
	if yosh < 0 {
		return fmt.Errorf("yosh manfiy bo'lishi mumkin emas: %d", yosh)
	}
	return nil
}
```

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

func bolish(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("nolga bo'lib bo'lmaydi")
	}
	return a / b, nil
}

func main() {
	natija, err := bolish(10, 2)
	if err != nil {
		fmt.Println("Xato:", err)
	} else {
		fmt.Println("Natija:", natija)
	}

	natija, err = bolish(10, 0)
	if err != nil {
		fmt.Println("Xato:", err)
	} else {
		fmt.Println("Natija:", natija)
	}
}
```

Natija:

```
Natija: 5
Xato: nolga bo'lib bo'lmaydi
```

## TASK

`bolish(a, b int) (int, error)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Agar `b == 0` bo'lsa, `0` va `errors.New("nolga bo'lib bo'lmaydi")` qaytarsin.
2. Aks holda, `a / b` va `nil` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `errors` paketini `import` qilishni unutmang — xatolik yaratish uchun `errors.New(...)` kerak bo'ladi.
2. `if b == 0 { return 0, errors.New("nolga bo'lib bo'lmaydi") }` keyin `return a / b, nil`.
