# 28 — Methods

## THEORY

Mashinani tasavvur qiling — uning "dvigatelni yoqish" tugmasi bor. Bu tugma mavhum, "kimdir uchun" emas — u aynan **shu mashinaga** tegishli, va bosilganda aynan shu mashinaning dvigateli yonadi. Go'da **method** aynan shunday ishlaydi: bu — muayyan bir turga "tegishli" bo'lgan funksiya.

**Method qanday e'lon qilinadi.** Oddiy funksiyadan farqi — nom oldida qavs ichida **receiver** (qabul qiluvchi) ko'rsatiladi:

```go
type Hisob struct {
	Balans int
}

func (h Hisob) Holat() string {
	return fmt.Sprintf("Balans: %d", h.Balans)
}
```

Bu yerda `(h Hisob)` — "bu method `Hisob` turiga tegishli, va uning ichida joriy nusxaga `h` nomi bilan murojaat qilinadi" degani. Chaqirish esa nuqta orqali bo'ladi, xuddi struct maydoniga murojaat qilgandek:

```go
mening := Hisob{Balans: 1000}
fmt.Println(mening.Holat()) // "Balans: 1000"
```

**Ikki xil receiver — qiymat va pointer.** Bu — methodlar bilan ishlashda eng muhim tushunish kerak bo'lgan narsa, va u to'g'ridan-to'g'ri oldingi "Pointers" darsiga bog'liq:

- **Qiymat receiver** (`func (h Hisob) ...`) — struct'ning **nusxasi** bilan ishlaydi. Method ichida `h`ni o'zgartirsangiz, bu faqat nusxaga ta'sir qiladi, asl obyekt o'zgarmaydi.
- **Pointer receiver** (`func (h *Hisob) ...`) — struct'ning **manziliga** ishora qiladi. Method ichida `h`ni o'zgartirsangiz, bu **asl obyektning o'zini** o'zgartiradi.

```go
func (h *Hisob) Tolda(summa int) {
	h.Balans += summa // *h ga yozish — asl obyekt o'zgaradi
}

mening := Hisob{Balans: 1000}
mening.Tolda(500)
fmt.Println(mening.Balans) // 1500 — chindan ham o'zgardi!
```

Diqqat qiling: `mening.Tolda(500)` deb chaqirganimizda, biz `&mening`ni qo'lda yozmadik — Go buni **avtomatik** qiladi, chunki `mening` — o'zgaruvchi (manzili bor). Bu — "Pointers" darsida ko'rgan struct pointerlarining avtomatik "dereference" qulayligining aynan teskarisi: bu yerda Go kerak bo'lganda avtomatik `&` qo'yadi.

**Qachon qaysi birini tanlash kerak?** Oddiy qoida: agar method obyektni **o'zgartirishi kerak bo'lsa** (masalan, balansni oshirish), pointer receiver ishlatiladi. Agar method faqat **o'qish** uchun ishlatilsa (masalan, joriy holatni matn qilib qaytarish), qiymat receiver ham yetarli. Amalda, ko'p Go dasturchilari bitta struct uchun barcha methodlarni **bir xil turdagi** receiver bilan yozishga harakat qiladi (izchillik uchun), hatto ba'zilari o'zgartirmasa ham.

## EXAMPLE

```go
package main

import "fmt"

type Hisob struct {
	Balans int
}

func (h Hisob) Holat() string {
	return fmt.Sprintf("Balans: %d", h.Balans)
}

func (h *Hisob) Tolda(summa int) {
	h.Balans += summa
}

func main() {
	mening := Hisob{Balans: 1000}
	fmt.Println(mening.Holat()) // Balans: 1000

	mening.Tolda(500)
	fmt.Println(mening.Holat()) // Balans: 1500
}
```

Natija:

```
Balans: 1000
Balans: 1500
```

## TASK

`Hisob` struct'i (`Balans int` maydoni bilan) berilgan. `Tolda(summa int)` methodini shunday to'ldiringki, u `Balans`ga `summa`ni qo'shsin (asl `Hisob` obyektini o'zgartirishi kerak — pointer receiver allaqachon yozilgan).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Receiver allaqachon pointer (`*Hisob`) sifatida yozilgan — siz shunchaki `h.Balans += summa` deb yozsangiz kifoya.
2. `h.Balans = h.Balans + summa` ham xuddi shu narsa, faqat uzunroq yozilgan — ikkalasi ham to'g'ri ishlaydi.
