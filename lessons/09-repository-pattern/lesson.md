# 09 — Repository Pattern

## THEORY

Bu darsgacha, biz `database/sql` bilan ishlash uchun kerakli SQL bilimlarni o'rgandik: ulanish, CRUD, tranzaksiya, indeks, join. Lekin katta dasturda, SQL so'rovlarini **har bir joyda** yozaverish — muammo tug'diradi: agar bazani PostgreSQL'dan boshqasiga o'zgartirish kerak bo'lsa, yoki so'rovlarni sinash (test) kerak bo'lsa, SQL kodi **butun dastur bo'ylab tarqalgan** bo'ladi.

**Yechim — Repository Pattern.** Bu — ma'lumotlar bazasi bilan ishlashning **barcha tafsilotlarini** (SQL so'rovlar, `database/sql` chaqiruvlari) bitta joyga — **repository**ga jamlash, va qolgan dastur qismlari bilan faqat **interfeys** orqali gaplashish g'oyasi.

**"Interfaces" darsini eslang** — u yerda interfeys, "qanday qilib" emas, balki "nima qila oladi"ni belgilaydi edi. Repository pattern — aynan shu g'oyani ma'lumotlar bazasiga qo'llaydi:

```go
type FoydalanuvchiRepository interface {
	Ol(id int) (Foydalanuvchi, error)
	Saqla(f Foydalanuvchi) error
	Ochir(id int) error
}
```

Bu interfeysning **haqiqiy** (PostgreSQL bilan ishlaydigan) implementatsiyasi:

```go
type PostgresRepository struct {
	db *sql.DB
}

func (r *PostgresRepository) Ol(id int) (Foydalanuvchi, error) {
	var f Foydalanuvchi
	qator := r.db.QueryRow("SELECT id, ism FROM foydalanuvchilar WHERE id = $1", id)
	err := qator.Scan(&f.ID, &f.Ism)
	return f, err
}

// Saqla va Ochir ham xuddi shunday, SQL orqali
```

**Nega bu foydali — "Mocking Concepts" darsi bilan bog'liq.** Dasturning qolgan qismi (masalan, HTTP handler) `FoydalanuvchiRepository` interfeysi bilan ishlaydi, `PostgresRepository`ning o'zi bilan emas. Bu shuni anglatadiki, **testlarda**, haqiqiy bazaga ulanmasdan, xuddi shu interfeysni amalga oshiruvchi **soxta (fake)** repository ishlatish mumkin:

```go
type XotiraRepository struct {
	malumotlar map[int]Foydalanuvchi
}

func (r *XotiraRepository) Ol(id int) (Foydalanuvchi, error) {
	f, bor := r.malumotlar[id]
	if !bor {
		return Foydalanuvchi{}, errors.New("topilmadi")
	}
	return f, nil
}
```

Ikkalasi ham bir xil interfeysni amalga oshiradi — dastur ularni **farqlamaydi**. Bu, aynan "04-crud" darsida ko'rgan `XotiraFoydalanuvchilar` g'oyasining davomi, lekin endi **rasmiy interfeys** orqali rasmiylashtirilgan.

**Amaliy foyda:**

1. **Sinov qulayligi** — haqiqiy PostgreSQL kerak bo'lmasdan, xotiradagi soxta repository bilan tez va ishonchli test yozish mumkin.
2. **Almashtirish qulayligi** — agar ertaga PostgreSQL o'rniga MySQL yoki boshqa baza ishlatish kerak bo'lsa, faqat **yangi implementatsiya** yozish kifoya, dasturning qolgan qismini o'zgartirish shart emas.
3. **Tozalik** — SQL so'rovlari, bitta faylda (yoki paketda) jamlangan bo'ladi, HTTP handler'lar va biznes-mantiq esa SQL haqida **umuman bilmaydi**.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

type Foydalanuvchi struct {
	ID  int
	Ism string
}

type FoydalanuvchiRepository interface {
	Ol(id int) (Foydalanuvchi, error)
}

type XotiraRepository struct {
	malumotlar map[int]Foydalanuvchi
}

func (r *XotiraRepository) Ol(id int) (Foydalanuvchi, error) {
	f, bor := r.malumotlar[id]
	if !bor {
		return Foydalanuvchi{}, errors.New("topilmadi")
	}
	return f, nil
}

func ismChiqar(repo FoydalanuvchiRepository, id int) string {
	f, err := repo.Ol(id)
	if err != nil {
		return "noma'lum"
	}
	return f.Ism
}

func main() {
	repo := &XotiraRepository{malumotlar: map[int]Foydalanuvchi{
		1: {ID: 1, Ism: "Ali"},
	}}

	fmt.Println(ismChiqar(repo, 1))
	fmt.Println(ismChiqar(repo, 99))
}
```

Natija:

```
Ali
noma'lum
```

`ismChiqar` funksiyasi `FoydalanuvchiRepository` interfeysi bilan ishlaydi — u `XotiraRepository` yoki haqiqiy `PostgresRepository` ekanligini **bilmaydi ham, bilishga hojati ham yo'q**.

## TASK

`XotiraRepository` strukturasi va `FoydalanuvchiRepository` interfeysi berilgan. `ismChiqar(repo FoydalanuvchiRepository, id int) string` funksiyasini to'ldiring:

1. `repo.Ol(id)` ni chaqiring.
2. Agar xatolik qaytsa, `"noma'lum"` qaytaring.
3. Aks holda, topilgan foydalanuvchining `Ism` maydonini qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `f, err := repo.Ol(id)` — natijada ikkita qiymat qaytadi, ikkalasini ham oling.
2. `if err != nil { return "noma'lum" }`, aks holda `return f.Ism`.
