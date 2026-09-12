# 03 — Database Connection

## THEORY

Go standart kutubxonasida `database/sql` degan paket bor — lekin bu paketning o'zi, hech qanday **aniq** ma'lumotlar bazasi bilan (PostgreSQL, MySQL va h.k.) qanday "gaplashishni" bilmaydi. `database/sql` — bu faqat **umumiy interfeys** ("Interfaces" darsini eslang): u "so'rov yubor", "natijani o'qi" kabi umumiy amallarni belgilaydi, lekin haqiqiy "tarjimon" vazifasini **drayver (driver)** bajaradi.

**Drayver — alohida paket, va u "bo'sh import" orqali ro'yxatdan o'tkaziladi:**

```go
import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL drayveri — faqat ro'yxatdan o'tkazish uchun import qilinadi
)
```

**Nega import nomi `_` (pastki chiziqcha).** "Scope" darsida ko'rgan qoidani eslang — Go'da ishlatilmagan import xato beradi. Lekin bu holatda, biz `pq` paketining hech qanday funksiyasini **to'g'ridan-to'g'ri** chaqirmaymiz — bizga faqat uning **"men ro'yxatdan o'taman" degan yon ta'siri (side effect)** kerak: har bir paket, import qilinganda, o'zining `init()` funksiyasini avtomatik ishga tushiradi, va `pq` paketi aynan shu `init()` ichida o'zini `database/sql`ga "men PostgreSQL bilan gaplasha olaman" deb ro'yxatdan o'tkazadi. `_` prefiksi Go'ga "bu paketni ataylab, faqat yon ta'siri uchun import qilyapman, funksiyalarini ishlatmayman" deb bildiradi.

**Diqqat: bu — kursning yagona, standart kutubxonadan tashqari kerak bo'ladigan qismi.** Go standart kutubxonasida hech qanday tayyor ma'lumotlar bazasi drayveri yo'q — shuning uchun PostgreSQL bilan ishlash uchun, `go get github.com/lib/pq` (yoki zamonaviyroq `github.com/jackc/pgx`) orqali, tashqi kutubxonani loyihangizga qo'shishga to'g'ri keladi.

**Ulanish — `sql.Open` va `Ping`:**

```go
db, err := sql.Open("postgres", dsn) // "PostgreSQL Basics" darsida ko'rgan DSN
if err != nil {
	return err
}
defer db.Close() // "Defer" va "HTTP Client" darslaridagi resurs tozalash naqshi

if err := db.Ping(); err != nil {
	return err // baza chindan ham javob berayotganini tekshiradi
}
```

**Muhim, ko'p adashiladigan nozik joy: `sql.Open` hali haqiqiy ulanish o'rnatmaydi!** `sql.Open` faqat DSN va drayver nomini **tekshiradi**, real tarmoq ulanishini darhol yaratmaydi — u "dangasa" (lazy): haqiqiy ulanish faqat birinchi haqiqiy so'rov (yoki `Ping()`) chaqirilganda amalga oshadi. Shuning uchun, `sql.Open` xato qaytarmagani, "baza ishlab turibdi" degani emas — buni bilish uchun `Ping()` kerak.

**Amaliy naqsh — "kutib, qayta urinish" (retry).** Ba'zan dastur ishga tushganda, baza hali to'liq tayyor bo'lmagan bo'lishi mumkin (masalan, Docker konteynerlar bir vaqtda ishga tushganda). Shuning uchun, `Ping()`ni **darhol muvaffaqiyatsiz bo'lsa taslim bo'lish** o'rniga, bir necha marta, orada kutib, qayta urinib ko'rish odat tusiga kiradi:

```go
func kutibUlanish(ping func() error, urinishlar int, kutish time.Duration) error {
	var oxirgiXato error
	for i := 0; i < urinishlar; i++ {
		if err := ping(); err == nil {
			return nil // muvaffaqiyatli
		} else {
			oxirgiXato = err
		}
		time.Sleep(kutish)
	}
	return oxirgiXato
}
```

Bu yerda `ping func() error` — "Higher-Order Functions" darsida ko'rgan naqsh: haqiqiy `db.Ping`ni to'g'ridan-to'g'ri chaqirish o'rniga, uni **parametr sifatida** qabul qilamiz — bu, funksiyani **haqiqiy bazasiz, "soxta" ping funksiyasi bilan** sinash imkonini beradi (xuddi "Mocking Concepts" darsida ko'rgan naqsh).

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
	"time"
)

func kutibUlanish(ping func() error, urinishlar int, kutish time.Duration) error {
	var oxirgiXato error
	for i := 0; i < urinishlar; i++ {
		if err := ping(); err == nil {
			return nil
		} else {
			oxirgiXato = err
		}
		time.Sleep(kutish)
	}
	return oxirgiXato
}

func main() {
	urinishSoni := 0
	soxtaPing := func() error {
		urinishSoni++
		if urinishSoni < 3 {
			return errors.New("baza hali tayyor emas")
		}
		return nil
	}

	err := kutibUlanish(soxtaPing, 5, time.Millisecond)
	fmt.Println(err)
	fmt.Println(urinishSoni)
}
```

Natija:

```
<nil>
3
```

## TASK

`kutibUlanish(ping func() error, urinishlar int, kutish time.Duration) error` funksiyasi berilgan. Uni shunday to'ldiringki, u `ping()`ni ko'pi bilan `urinishlar` marta chaqirsin, har muvaffaqiyatsiz urinishdan keyin `kutish` vaqtga kutib tursin. Agar biror urinish muvaffaqiyatli (`nil` xatolik) bo'lsa, darhol `nil` qaytarsin. Agar barcha urinishlar muvaffaqiyatsiz bo'lsa, **eng oxirgi** xatolikni qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `for i := 0; i < urinishlar; i++ { if err := ping(); err == nil { return nil } else { oxirgiXato = err }; time.Sleep(kutish) }`.
2. Tsikldan tashqarida `var oxirgiXato error` e'lon qiling (tsikldan oldin), tsikldan keyin `return oxirgiXato`.
