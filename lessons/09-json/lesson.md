# 09 — JSON

## THEORY

Turli davlatlarga pochta jo'natish uchun bir xil, universal shakldagi konteynerlar ishlatiladi — konteynerning ichida nima bo'lishidan qat'iy nazar, uni istalgan port qabul qilib, ocha oladi. **JSON** (JavaScript Object Notation) — dasturlar orasida ma'lumot almashishning aynan shunday universal formati: Go, Python, JavaScript, va boshqa deyarli har qanday til JSON'ni tushunadi.

**Struct'ni JSON'ga aylantirish — `json.Marshal`.** "Structs" darsida ko'rgan struct'ingizni JSON matniga aylantirish uchun `encoding/json` paketidagi `Marshal` funksiyasi ishlatiladi:

```go
type Mahsulot struct {
	Nomi string
	Narx int
}

m := Mahsulot{Nomi: "Non", Narx: 5000}
baytlar, err := json.Marshal(m)
// baytlar — []byte turida, string(baytlar) == `{"Nomi":"Non","Narx":5000}`
```

`json.Marshal` ikkita qiymat qaytaradi — natija (`[]byte` turida, ya'ni baytlar to'plami) va xatolik ("Errors" darsida ko'rgan naqsh). Odatda struct'larni JSON'ga aylantirish deyarli hech qachon xato bermaydi, lekin Go qoidasiga ko'ra, xato qaytarishi mumkin bo'lgan har qanday funksiya baribir xatolikni ham qaytaradi.

**Struct teglari (tags) — JSON'dagi maydon nomini boshqarish.** Standart holatda, JSON'dagi maydon nomlari struct maydon nomlari bilan bir xil bo'ladi (`Nomi`, `Narx` — katta harf bilan, chunki Go struct maydonlari odatda shunday). Ko'pincha esa JSON'da kichik harfli nomlar (`nomi`, `narx`) kerak bo'ladi — buning uchun struct **tegi** ishlatiladi:

```go
type Mahsulot struct {
	Nomi string `json:"nomi"`
	Narx int    `json:"narx"`
}
```

Endi `json.Marshal(m)` natijasi `{"nomi":"Non","narx":5000}` bo'ladi — orqa qatorli qo'shtirnoq (backtick) ichidagi `` `json:"nomi"` `` — bu Go'ning **struct tegi** deb ataladigan maxsus, metama'lumot yozish usuli.

**Teskari yo'nalish — `json.Unmarshal`.** JSON matnini qayta struct'ga aylantirish uchun `Unmarshal` ishlatiladi — bu funksiya natijani **qaytarish** o'rniga, siz bergan struct'ning **manzili** (pointer) orqali uni to'ldiradi ("Pointers" darsida ko'rgan naqshni eslang):

```go
matn := `{"nomi":"Sut","narx":8000}`
var m2 Mahsulot
err := json.Unmarshal([]byte(matn), &m2)
fmt.Println(m2.Nomi, m2.Narx) // Sut 8000
```

`&m2` — bu yerda majburiy, chunki `Unmarshal` `m2`ning o'zini emas, balki uning xotiradagi manzilini olib, o'sha manzildagi qiymatni to'g'ridan-to'g'ri o'zgartirishi kerak.

## EXAMPLE

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Mahsulot struct {
	Nomi string `json:"nomi"`
	Narx int    `json:"narx"`
}

func main() {
	m := Mahsulot{Nomi: "Non", Narx: 5000}
	baytlar, _ := json.Marshal(m)
	fmt.Println(string(baytlar))

	var m2 Mahsulot
	json.Unmarshal([]byte(`{"nomi":"Sut","narx":8000}`), &m2)
	fmt.Println(m2.Nomi, m2.Narx)
}
```

Natija:

```
{"nomi":"Non","narx":5000}
Sut 8000
```

## TASK

`Mahsulot` struct'i (`Nomi string` — tegi `json:"nomi"`, `Narx int` — tegi `json:"narx"`, ikkalasi ham allaqachon yozilgan) va `toJSON(m Mahsulot) (string, error)` funksiyasi berilgan. Funksiyani shunday to'ldiringki, u `m`ni JSON matniga aylantirib, string va xatolik (agar bo'lsa) qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `json.Marshal(m)` — `[]byte` va xatolik qaytaradi, buni `string(...)` orqali matnga aylantiring.
2. `baytlar, err := json.Marshal(m); return string(baytlar), err` — natijani xatolik bilan birga qaytaring.
