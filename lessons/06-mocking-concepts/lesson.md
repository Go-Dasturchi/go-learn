# 06 — Mocking Concepts

## THEORY

Uchuvchini tayyorlash uchun uni darhol haqiqiy samolyotga o'tqazishmaydi — avval **trenajyor (simulyator)**da mashq qildirishadi: u haqiqiy samolyotning boshqaruv paneliga o'xshaydi, lekin real parvoz qilmaydi, va instruktor xohlagan holatni (bo'ron, dvigatel nosozligi) darhol "yaratib" bera oladi. **Mock (soxta obyekt)** — dasturlashda aynan shu vazifani bajaradi: sinov paytida, haqiqiy (sekin, ishonchsiz, yoki nazoratsiz) qismni, uning **o'rnini bosuvchi, to'liq nazorat qilinadigan soxta versiya** bilan almashtirish.

**Nega mock kerak.** Tasavvur qiling, sizning funksiyangiz xabar yuborish xizmatiga (masalan, SMS yoki email) bog'liq. Buni test qilishda, **haqiqiy** SMS yuborishni istamaysiz — bu sekin, pulga tushishi mumkin, va internetga bog'liq (testlar esa har doim, hamma joyda, tez ishlashi kerak). Yechim — **interfeys** orqali "xabar yuborish" imkoniyatini mavhumlashtirish ("Interfaces" darsini eslang):

```go
type Xabarchi interface {
	Yubor(matn string) error
}
```

**Haqiqiy va soxta implementatsiya.** Ishlab chiqarishda (production) haqiqiy SMS/email yuboruvchi struct ishlatiladi. Testda esa — **soxta (mock)** versiya, u hech narsani chindan yubormaydi, faqat "menga nima berilgandi" deb **eslab qoladi**:

```go
type SoxtaXabarchi struct {
	YuborilganXabarlar []string
}

func (s *SoxtaXabarchi) Yubor(matn string) error {
	s.YuborilganXabarlar = append(s.YuborilganXabarlar, matn)
	return nil
}
```

**Interfeys orqali "almashtirish" — Go'ning mock uchun tabiiy usuli.** Funksiya haqiqiy struct'ga emas, balki **interfeysga** bog'liq bo'lsa, unga ishlab chiqarishda haqiqiy, testda esa soxta implementatsiyani "almashtirib" berish mumkin — funksiyaning o'zi ikkalasining farqini bilmaydi ham:

```go
func Bildirish(x Xabarchi, matn string) error {
	return x.Yubor(matn)
}

// Ishlab chiqarishda:
Bildirish(haqiqiySMS, "Salom!")

// Testda:
soxta := &SoxtaXabarchi{}
Bildirish(soxta, "Salom!")
// endi soxta.YuborilganXabarlar ichida "Salom!" borligini tekshirish mumkin — hech narsa chindan yuborilmadi!
```

**Nega bu Go'da ayniqsa tabiiy ishlaydi.** "Interfaces" darsida ko'rgan **implicit implementatsiya**ni eslang — `SoxtaXabarchi` hech qayerda "men `Xabarchi`man" demasdan, faqat `Yubor(string) error` methodiga ega bo'lgani uchun, avtomatik ravishda `Xabarchi` sifatida ishlatilishi mumkin. Bu, Go'da mocking uchun maxsus tashqi kutubxona shart emasligining sababi — oddiy interfeys va struct orqali, to'liq nazorat qilinadigan soxta obyektlar yaratish mumkin.

## EXAMPLE

```go
package main

import "fmt"

type Xabarchi interface {
	Yubor(matn string) error
}

type SoxtaXabarchi struct {
	YuborilganXabarlar []string
}

func (s *SoxtaXabarchi) Yubor(matn string) error {
	s.YuborilganXabarlar = append(s.YuborilganXabarlar, matn)
	return nil
}

func Bildirish(x Xabarchi, matn string) error {
	return x.Yubor(matn)
}

func main() {
	soxta := &SoxtaXabarchi{}
	Bildirish(soxta, "Salom!")
	Bildirish(soxta, "Xayr!")

	fmt.Println(soxta.YuborilganXabarlar)
}
```

Natija:

```
[Salom! Xayr!]
```

## TASK

`Xabarchi` interfeysi va `Bildirish` funksiyasi berilgan. `SoxtaXabarchi` struct'ining `Yubor(matn string) error` methodini shunday to'ldiringki, u xabarni chindan yubormasdan, uni `YuborilganXabarlar` slice'iga qo'shib, `nil` (xatolik yo'q) qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `s.YuborilganXabarlar = append(s.YuborilganXabarlar, matn)` — xabarni "yuborish" o'rniga, ro'yxatga qo'shib qo'yasiz.
2. `return nil` — hech qanday xatolik yo'q, chunki bu shunchaki soxta (mock) versiya.
