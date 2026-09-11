# 27 — Pointers

## THEORY

Uy manzilini qog'ozga yozib, do'stingizga berganingizni tasavvur qiling. Bu qog'oz **uyning o'zi emas** — u shunchaki uyni qayerdan topish mumkinligini ko'rsatadi. Lekin muhim narsa shu: agar do'stingiz o'sha manzil bo'yicha borib, uyning eshigini bo'yasa, **haqiqiy uy** o'zgaradi — chunki qog'ozdagi manzil aynan o'sha haqiqiy uyga olib boradi, uyning nusxasiga emas. Go'da **pointer (ko'rsatkich)** aynan shu — bu qiymatning o'zi emas, balki uning xotiradagi **manzili**.

**Nega bu kerak — "Fundamentals"dan eslang.** Siz allaqachon bilasiz: Go'da funksiyaga argument berilganda, u **nusxalanadi** (bu — array va struct'lar uchun ayniqsa aniq ko'rinadi, "Arrays" va "Structs" darslarida ko'rgan edingiz). Demak, agar funksiya ichida oddiy parametrni o'zgartirsangiz, bu **faqat nusxaga** ta'sir qiladi, chaqiruvchi tomondagi asl o'zgaruvchi o'zgarmay qoladi:

```go
func ozgartir(son int) {
	son = 100 // faqat mahalliy nusxa o'zgaradi
}

x := 5
ozgartir(x)
fmt.Println(x) // 5 — o'zgarmadi!
```

Agar funksiya chaqiruvchi tomondagi **asl** o'zgaruvchini o'zgartirishi kerak bo'lsa, unga qiymatning o'zini emas, balki uning **manzilini** berish kerak — mana shu yerda pointer kerak bo'ladi.

**`&` — manzilni olish, `*` — manzil bo'yicha borib, qiymatni o'qish/yozish:**

```go
x := 5
manzil := &x        // &x — x ning xotiradagi manzili, turi *int
fmt.Println(*manzil) // *manzil — o'sha manzildagi qiymatni o'qish: 5

*manzil = 100         // o'sha manzildagi qiymatni yangilash
fmt.Println(x)        // 100 — x HAM o'zgardi! chunki manzil aynan x ga ishora qiladi
```

`*int` — "bu `int`ning manzilini saqlaydigan tur" degani. `&x` — "x ning manzilini ol" degani. `*manzil` — "manzil bo'yicha borib, u yerdagi qiymatni ol (yoki yoz)" degani.

**Pointer orqali funksiyani "chinakam" o'zgartiruvchi qilish:**

```go
func ozgartir(son *int) {
	*son = 100 // manzil bo'yicha borib, ASL qiymatni o'zgartiryapmiz
}

x := 5
ozgartir(&x) // x ning manzilini beramiz, x ning o'zini emas
fmt.Println(x) // 100 — endi chindan ham o'zgardi!
```

**`nil` pointer.** Agar pointer hech qanday manzilga ishora qilmasa, u `nil` bo'ladi — buni "bo'sh qog'oz, hech qanday manzil yozilmagan" deb tasavvur qiling. `nil` pointer orqali `*` bilan qiymat o'qishga yoki yozishga urinish dastur ishlashini to'xtatadi (panic):

```go
var p *int    // nil pointer
fmt.Println(*p) // PANIC: manzil yo'q, "borib qiymat olish" mumkin emas
```

**Struct pointerlari — qulaylik uchun avtomatik "dereference".** Struct'ga pointer bilan ishlaganda, Go sizga har safar `(*p).Maydon` deb yozishga majburlamaydi — `p.Maydon` deb yozsangiz ham, agar `p` pointer bo'lsa, Go buni avtomatik tushunadi:

```go
type Odam struct {
	Yosh int
}

func kattalashtir(o *Odam) {
	o.Yosh++ // aslida (*o).Yosh++ demak, lekin qisqa yozuv ishlaydi
}

ali := Odam{Yosh: 25}
kattalashtir(&ali)
fmt.Println(ali.Yosh) // 26
```

Bu naqsh — struct'ni funksiya ichida **haqiqatan o'zgartirish** kerak bo'lganda (masalan, keyingi darsda ko'radigan pointer receiver'li method'larda) juda ko'p ishlatiladi.

## EXAMPLE

```go
package main

import "fmt"

func ozgartir(son *int) {
	*son = 100
}

type Odam struct {
	Yosh int
}

func kattalashtir(o *Odam) {
	o.Yosh++
}

func main() {
	x := 5
	ozgartir(&x)
	fmt.Println(x) // 100

	ali := Odam{Yosh: 25}
	kattalashtir(&ali)
	fmt.Println(ali.Yosh) // 26
}
```

Natija:

```
100
26
```

## TASK

`oshir(son *int, qiymat int)` funksiyasi berilgan. Uni shunday to'ldiringki, u `son` ko'rsatib turgan qiymatga `qiymat`ni qo'shsin (funksiya hech narsa qaytarmaydi, lekin chaqiruvchi tomondagi asl o'zgaruvchi o'zgarishi kerak).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `*son` — `son` ko'rsatib turgan qiymatning o'zi. Uni o'qish ham, yozish ham mumkin.
2. `*son = *son + qiymat` — joriy qiymatni o'qib, ustiga `qiymat`ni qo'shib, natijani yana o'sha manzilga yozing.
