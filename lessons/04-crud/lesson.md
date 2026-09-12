# 04 — CRUD

## THEORY

**CRUD** — to'rtta harfning qisqartmasi: **C**reate (yaratish), **R**ead (o'qish), **U**pdate (yangilash), **D**elete (o'chirish). Deyarli har qanday dastur — kim bilandir "narsalarni yaratish, ko'rish, o'zgartirish, o'chirish" bilan shug'ullanadi, va "SQL Basics" darsida ko'rgan `INSERT`/`SELECT`/`UPDATE`/`DELETE` — aynan shu to'rtta amalning SQL'dagi ko'rinishi.

**Go'da CRUD — `db.Exec` va `db.Query`/`db.QueryRow`:**

```go
// Create — natija qaytarmaydigan amal, Exec ishlatiladi
_, err := db.Exec("INSERT INTO foydalanuvchilar (ism, yosh) VALUES ($1, $2)", "Ali", 25)

// Read (bitta qator) — QueryRow
var ism string
var yosh int
err := db.QueryRow("SELECT ism, yosh FROM foydalanuvchilar WHERE id = $1", 1).Scan(&ism, &yosh)

// Read (bir nechta qator) — Query
rows, err := db.Query("SELECT ism, yosh FROM foydalanuvchilar")
defer rows.Close()
for rows.Next() {
	var ism string
	var yosh int
	rows.Scan(&ism, &yosh)
	fmt.Println(ism, yosh)
}

// Update / Delete — bularda ham natija yo'q, Exec ishlatiladi
_, err = db.Exec("UPDATE foydalanuvchilar SET yosh = $1 WHERE id = $2", 26, 1)
_, err = db.Exec("DELETE FROM foydalanuvchilar WHERE id = $1", 1)
```

**`$1`, `$2` — parametrlashtirilgan so'rovlar, xavfsizlik uchun MUHIM.** E'tibor bering: qiymatlar SQL matni ichiga to'g'ridan-to'g'ri "yopishtirilmagan" (masalan, `fmt.Sprintf` bilan) — ular `$1`, `$2` kabi **joy egallovchilar** orqali, alohida argument sifatida beriladi. Bu — **SQL in'ektsiyasi (SQL injection)** deb ataladigan, jiddiy xavfsizlik zaifligining oldini oladi: agar foydalanuvchi ismi o'rniga `'; DROP TABLE foydalanuvchilar; --` kabi zararli matn kiritsa, va siz uni to'g'ridan-to'g'ri SQL matniga "yopishtirsangiz", bu butun jadvalni o'chirib yuborishi mumkin edi. Parametrlashtirilgan so'rovlarda esa, drayver qiymatni **har doim oddiy ma'lumot sifatida**, hech qachon "buyruq" sifatida talqin qilmaydi.

**`Scan` — natijani Go o'zgaruvchilariga o'qish.** "JSON Request Body" darsida ko'rgan `Decode(&struct)`ga o'xshaydi — `Scan(&ism, &yosh)` ham natijani, siz bergan o'zgaruvchilarning **manzillariga** ("Pointers" darsini eslang) yozib beradi.

**Repository naqshi — CRUD'ni interfeys ortiga yashirish.** Haqiqiy loyihalarda, CRUD kodini har bir handler'ga sochib tashlash o'rniga, uni **alohida struct/interfeys**ga jamlash odat tusiga kiradi ("Interfaces" va "Mocking Concepts" darslarini eslang):

```go
type Foydalanuvchilar interface {
	Yarat(ism string, yosh int) error
	Ol(id int) (string, int, error)
}
```

Bu — "Repository Pattern" darsida to'liq ko'rib chiqiladigan, muhim naqsh. Hozircha, buni **haqiqiy bazasiz**, xotirada (`map` orqali) sinash uchun, sodda "soxta" (fake) implementatsiya yozamiz — bu, aynan haqiqiy, Postgres-ga ulangan implementatsiya keyinchalik qanday ishlashini ko'rsatadi, lekin uni avtomatik testda ishlatish imkonini beradi.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

type XotiraFoydalanuvchilar struct {
	malumot map[int]string
}

func (x *XotiraFoydalanuvchilar) Yarat(id int, ism string) {
	if x.malumot == nil {
		x.malumot = make(map[int]string)
	}
	x.malumot[id] = ism
}

func (x *XotiraFoydalanuvchilar) Ol(id int) (string, error) {
	ism, bor := x.malumot[id]
	if !bor {
		return "", errors.New("topilmadi")
	}
	return ism, nil
}

func main() {
	repo := &XotiraFoydalanuvchilar{}
	repo.Yarat(1, "Ali")

	ism, err := repo.Ol(1)
	fmt.Println(ism, err)

	_, err = repo.Ol(99)
	fmt.Println(err)
}
```

Natija:

```
Ali <nil>
topilmadi
```

## TASK

`XotiraFoydalanuvchilar` struct'i va `Yarat` methodi berilgan (allaqachon yozilgan). `Ol(id int) (string, error)` methodini shunday to'ldiringki, u `id` bo'yicha ismni xaritadan topib qaytarsin; agar topilmasa, `"", errors.New("topilmadi")` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ism, bor := x.malumot[id]` — "Maps" darsida ko'rgan comma-ok naqshi.
2. `if !bor { return "", errors.New("topilmadi") }`, aks holda `return ism, nil`.
