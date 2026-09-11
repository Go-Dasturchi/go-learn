# 26 — Nested Structs

## THEORY

Pasportingizni yana bir bor tasavvur qiling — unda "Manzil" degan alohida bo'lim bor, va bu bo'limning o'zi ham ichida "Shahar", "Ko'cha" kabi o'z maydonlariga ega. Ya'ni bir hujjat ichida **kichikroq hujjat** joylashgan. Go'da bir struct'ning maydoni sifatida **boshqa struct**ni ishlatish mumkin — bunga **ichma-ich struct (nested struct)** deyiladi.

**E'lon qilish va ishlatish:**

```go
type Manzil struct {
	Shahar string
	Kocha  string
}

type Odam struct {
	Ism    string
	Manzil Manzil // Manzil turidagi maydon
}
```

Struct yaratganda, ichki struct ham o'z navbatida to'ldiriladi:

```go
ali := Odam{
	Ism: "Ali",
	Manzil: Manzil{
		Shahar: "Toshkent",
		Kocha:  "Amir Temur",
	},
}
```

**Ichki maydonlarga murojaat — zanjirlangan nuqta bilan:**

```go
fmt.Println(ali.Manzil.Shahar) // "Toshkent"
ali.Manzil.Shahar = "Samarqand" // ichki maydonni ham yangilash mumkin
```

Buni pasportni ochib, "Manzil" bo'limiga o'tib, u yerdan "Shahar" qatorini o'qishga o'xshatish mumkin — har bir nuqta, hujjat ichida bir bosqich chuqurroqqa kirishni bildiradi.

**Anonim (embedded) maydonlar — ichki struct'ni "yashirin" qilib qo'shish.** Go'da yana bir qiziq imkoniyat bor: ichki struct'ga alohida nom bermasdan, faqat uning **turini** yozib qo'yish mumkin — bunga "embedding" (joylashtirish) deyiladi:

```go
type Odam struct {
	Ism string
	Manzil // nomi yo'q — faqat turi yozilgan
}
```

Bunday holatda, `Manzil`ning maydonlariga **to'g'ridan-to'g'ri**, `ali.Manzil.Shahar` deb emas, `ali.Shahar` deb ham murojaat qilish mumkin bo'ladi — Go bu maydonlarni avtomatik "ko'tarib" (promote qilib) chiqaradi:

```go
ali := Odam{Ism: "Ali", Manzil: Manzil{Shahar: "Toshkent"}}
fmt.Println(ali.Shahar) // "Toshkent" — to'g'ridan-to'g'ri, oraliq nomsiz
```

Bu naqsh keyingi darslarda (Methods, Interfaces) struct'lar orasida umumiy xatti-harakatlarni "meros qilib olish"ga o'xshash effekt yaratish uchun juda ko'p ishlatiladi.

## EXAMPLE

```go
package main

import "fmt"

type Manzil struct {
	Shahar string
	Kocha  string
}

type Odam struct {
	Ism    string
	Manzil Manzil
}

func main() {
	ali := Odam{
		Ism: "Ali",
		Manzil: Manzil{
			Shahar: "Toshkent",
			Kocha:  "Amir Temur",
		},
	}

	fmt.Println(ali.Ism)
	fmt.Println(ali.Manzil.Shahar)
	fmt.Println(ali.Manzil.Kocha)

	ali.Manzil.Shahar = "Samarqand"
	fmt.Println(ali.Manzil.Shahar)
}
```

Natija:

```
Ali
Toshkent
Amir Temur
Samarqand
```

## TASK

`Manzil` (`Shahar string`) va `Odam` (`Ism string`, `Manzil Manzil`) struct'lari berilgan. `toliqMalumot(o Odam) string` funksiyasini shunday to'ldiringki, u `"Ism, Shahar"` ko'rinishidagi matnni qaytarsin (masalan, `Odam{Ism: "Ali", Manzil: Manzil{Shahar: "Toshkent"}}` uchun — `"Ali, Toshkent"`).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ichki struct maydoniga zanjirlangan nuqta orqali murojaat qiling: `o.Manzil.Shahar`.
2. `return o.Ism + ", " + o.Manzil.Shahar` — ikkita matnni vergul va bo'shliq bilan birlashtiring.
