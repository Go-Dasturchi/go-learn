# 13 — Switch

## THEORY

`switch`ni lift tugmalar paneliga o'xshatish mumkin: qaysi qavat tugmasini bossangiz, faqat o'sha qavatga borasiz — boshqa qavatlarni "tekshirib" o'tirmaysiz. Ko'p `if`/`else if`/`else` zanjiri o'rniga `switch` ishlatish ko'proq o'qilishi oson bo'ladi:

```go
switch kun {
case 1:
	fmt.Println("Dushanba")
case 2:
	fmt.Println("Seshanba")
default:
	fmt.Println("Boshqa kun")
}
```

**Muhim farq boshqa tillardan.** Ko'plab tillarda (masalan C, Java, JavaScript) `switch`ning har bir `case`i tugagach, agar ataylab `break` yozilmasa, kod **keyingi `case`ga ham "quyilib" ketaveradi** — bu ko'p yangi dasturchilarni ushlab qoladigan mashhur xato manbai. Go bu muammoni **butunlay yo'q qilgan**: har bir `case` avtomatik ravishda tugaydi, `break` yozish shart **emas**. Agar chindan ham ataylab keyingi `case`ga o'tishni xohlasangiz (bu juda kam kerak bo'ladi), maxsus `fallthrough` kalit so'zi bor:

```go
switch kun {
case 1:
	fmt.Println("Hafta boshlanishi")
	fallthrough // ataylab keyingi case'ga ham o'tamiz
case 2:
	fmt.Println("Hali hafta boshi")
}
```

**Bir nechta qiymatni bitta `case`da tekshirish.** Vergul bilan ajratib, bir nechta qiymatni bitta `case` ichiga yig'ish mumkin:

```go
switch kun {
case 6, 7:
	fmt.Println("Dam olish kuni")
default:
	fmt.Println("Ish kuni")
}
```

**Shartsiz `switch` — `if/else if` zanjirining chiroyli o'rinbosari.** `switch` faqat aniq bir qiymatni solishtirish uchun emas — uni hech qanday qiymatsiz ham yozib, har bir `case`ga to'liq shart yozish mumkin. Bu ayniqsa uzun `if/else if` zanjirlarini ancha tozaroq ko'rinishga keltiradi:

```go
switch {
case ball >= 90:
	fmt.Println("A'lo")
case ball >= 70:
	fmt.Println("Yaxshi")
default:
	fmt.Println("Qoniqarli")
}
```

**`switch` boshlang'ich amal bilan.** Xuddi `if`dagi kabi (04-if-else darsida ko'rgan edingiz), `switch`dan oldin ham kichik bir amal bajarib, uning natijasini darhol shu `switch` doirasida solishtirish mumkin:

```go
switch kun := hafta.KunNomi(); kun {
case "Shanba", "Yakshanba":
	fmt.Println("Dam olish kuni")
default:
	fmt.Println("Ish kuni")
}
```

**`switch` faqat sonlar bilan emas.** `switch` matnlar (string) bilan ham, hatto boshqa solishtirsa bo'ladigan istalgan tur bilan ham ishlaydi:

```go
switch rang {
case "qizil":
	fmt.Println("Xavf!")
case "yashil":
	fmt.Println("Xavfsiz")
default:
	fmt.Println("Noma'lum rang")
}
```

## EXAMPLE

```go
package main

import "fmt"

func baho(ball int) string {
	switch {
	case ball >= 90:
		return "A'lo"
	case ball >= 70:
		return "Yaxshi"
	default:
		return "Qoniqarli"
	}
}

func kunTuri(kun int) string {
	switch kun {
	case 6, 7:
		return "Dam olish"
	default:
		return "Ish kuni"
	}
}

func main() {
	fmt.Println(baho(95))
	fmt.Println(baho(75))
	fmt.Println(baho(40))
	fmt.Println(kunTuri(7))
}
```

Natija:

```
A'lo
Yaxshi
Qoniqarli
Dam olish
```

## TASK

`kunNomi(kun int) string` funksiyasi berilgan. `switch` yordamida uni shunday to'ldiringki:

- `1` bo'lsa — `"Dushanba"`
- `6` yoki `7` bo'lsa — `"Dam olish kuni"`
- boshqa har qanday qiymatda — `"Oddiy kun"`

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `switch kun { case 1: return "Dushanba" ... }` — `kun` qiymatiga qarab shoxlanadi.
2. Bir nechta qiymatni bitta `case`da vergul bilan yozing: `case 6, 7:`.
