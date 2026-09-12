# 02 — Dependency Injection

## THEORY

"Clean Architecture" darsida ko'rdik: `FoydalanuvchiUseCase`, `FoydalanuvchiOmbori` interfeysiga muhtoj. Lekin savol tug'iladi — **qaysi aniq implementatsiya** (`XotiraOmbori`, `PostgresOmbori`, ...) ishlatilishini **kim va qayerda** hal qiladi?

**Yomon yechim — struct ichida yaratish:**

```go
type FoydalanuvchiUseCase struct {
	ombor *PostgresOmbori
}

func YangiUseCase() *FoydalanuvchiUseCase {
	return &FoydalanuvchiUseCase{ombor: &PostgresOmbori{db: ochUlanish()}} // qattiq bog'langan!
}
```

Bu yerda `FoydalanuvchiUseCase`, **o'zi** `PostgresOmbori` yaratadi — bu, uni sinashda haqiqiy bazasiz ishlatib bo'lmaydigan qiladi, chunki ombor **tashqaridan berilmagan**, ichkarida "qattiq bog'langan" (hardcoded).

**Yaxshi yechim — Dependency Injection (bog'liqlikni kiritish).** Bog'liqlikni struct **o'zi yaratmaydi** — uni **tashqaridan, konstruktor orqali qabul qiladi**:

```go
func YangiUseCase(ombor FoydalanuvchiOmbori) *FoydalanuvchiUseCase {
	return &FoydalanuvchiUseCase{ombor: ombor}
}
```

Endi, **kim chaqirsa** — o'zi qaysi ombor kerakligini hal qiladi:

```go
// Haqiqiy dasturda:
useCase := YangiUseCase(&PostgresOmbori{db: db})

// Testda:
useCase := YangiUseCase(&XotiraOmbori{malumotlar: soxtaMalumot})
```

**Bu — g'oyaning o'zi, "Interfaces" va "Repository Pattern" darslarida ko'rgan naqshning davomi**, lekin endi e'tibor **qayerda ulanish sodir bo'lishiga** qaratilgan: barcha bog'liqliklar, odatda dasturning **eng yuqori nuqtasida** (`main()` funksiyasida, yoki alohida "wiring" faylida) yig'iladi va pastga uzatiladi — bu, "Dependency Injection" nomining o'zidan kelib chiqadi: bog'liqlik **tashqaridan kiritiladi** (inject qilinadi), ichkarida yaratilmaydi.

**Nega bu muhim:**

1. **Test qulayligi** — har qanday komponentni, haqiqiy bog'liqliklarsiz, soxta implementatsiyalar bilan **izolyatsiyalab** sinash mumkin.
2. **Moslashuvchanlik** — bir xil komponent, turli konfiguratsiyalarda (masalan, rivojlanish muhitida xotiradagi baza, production'da PostgreSQL) qayta yozilmasdan ishlaydi.
3. **Aniqlik** — bir struct nimaga bog'liqligi, uning konstruktoriga qarab **darhol ko'rinadi** — struct ichiga chuqur kirib, qaysi paketlar chaqirilayotganini qidirishga hojat qolmaydi.

**Katta loyihalarda**, bu "wiring" jarayoni ba'zan maxsus kutubxonalar (masalan, `google/wire` yoki `uber-go/fx`) yordamida avtomatlashtiriladi, lekin **asosiy g'oya** — bog'liqlikni konstruktor orqali uzatish — har doim **oddiy Go**, hech qanday kutubxonasiz ishlaydi, va aynan shu darsda shuni mashq qilamiz.

## EXAMPLE

```go
package main

import "fmt"

type Bildirishnoma interface {
	Yubor(xabar string) string
}

type EmailBildirishnoma struct{}

func (e *EmailBildirishnoma) Yubor(xabar string) string {
	return "Email: " + xabar
}

type SMSBildirishnoma struct{}

func (s *SMSBildirishnoma) Yubor(xabar string) string {
	return "SMS: " + xabar
}

// Servis — Bildirishnoma bog'liqligini konstruktor orqali qabul qiladi,
// qaysi aniq turdan foydalanishni o'zi hal qilmaydi.
type Servis struct {
	bildirishnoma Bildirishnoma
}

func YangiServis(b Bildirishnoma) *Servis {
	return &Servis{bildirishnoma: b}
}

func (s *Servis) XabarYubor(matn string) string {
	return s.bildirishnoma.Yubor(matn)
}

func main() {
	emailServis := YangiServis(&EmailBildirishnoma{})
	fmt.Println(emailServis.XabarYubor("Salom"))

	smsServis := YangiServis(&SMSBildirishnoma{})
	fmt.Println(smsServis.XabarYubor("Salom"))
}
```

Natija:

```
Email: Salom
SMS: Salom
```

Bitta `Servis` kodi — **hech qanday o'zgarishsiz**, ikkita turli bog'liqlik bilan ishladi, chunki bog'liqlik konstruktor orqali **tashqaridan kiritildi**.

## TASK

`YangiServis(b Bildirishnoma) *Servis` konstruktorini va `Servis.XabarYubor(matn string) string` methodini to'ldiring:

1. `YangiServis` — `b` argumentini `Servis`ning `bildirishnoma` maydoniga qo'yib, `*Servis` qaytarsin.
2. `XabarYubor` — `s.bildirishnoma.Yubor(matn)` ni chaqirib, natijasini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `func YangiServis(b Bildirishnoma) *Servis { return &Servis{bildirishnoma: b} }`.
2. `func (s *Servis) XabarYubor(matn string) string { return s.bildirishnoma.Yubor(matn) }`.
