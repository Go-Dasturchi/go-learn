# 08 — Connection Pool

## THEORY

Ma'lumotlar bazasiga **har bir so'rov uchun** yangi ulanish (connection) ochish — juda qimmat amal: TCP handshake, autentifikatsiya, va boshqa tayyorgarlik amallari **millisekundlar** talab qiladi. Agar sizning veb-serveringiz soniyasiga yuzlab so'rov qabul qilsa, har birida yangi ulanish ochish — serverni **sekinlashtiradi** va ma'lumotlar bazasini **ortiqcha yuklaydi**.

**Yechim — Connection Pool (ulanishlar havzasi).** Bu — oldindan tayyorlab qo'yilgan, **qayta ishlatiladigan** ulanishlar to'plami: dastur ulanish kerak bo'lganda, havzadan **bo'sh** ulanishni oladi, ishlatib bo'lgach, uni **yopmasdan**, havzaga qaytaradi — keyingi so'rov uni qayta ishlatadi.

**Go'da `database/sql` paketi — buni avtomatik boshqaradi.** `sql.Open()` chaqirilganda, u darhol ulanish ochmaydi — balki, keyinchalik so'rovlar kelganda ulanishlarni ochadigan va ularni ichki havzada saqlaydigan **obyekt** yaratadi. Havzani sozlash mumkin:

```go
db, err := sql.Open("postgres", dsn)
if err != nil {
	return err
}

db.SetMaxOpenConns(25)                  // bir vaqtda maksimal ochiq ulanishlar soni
db.SetMaxIdleConns(5)                   // "bo'sh turgan" holatda saqlanadigan ulanishlar soni
db.SetConnMaxLifetime(5 * time.Minute)  // har bir ulanishning maksimal "umri"
```

- `SetMaxOpenConns` — bazani **ortiqcha yuklanishdan** himoya qiladi (masalan, PostgreSQL odatda cheklangan sonli parallel ulanishni qo'llab-quvvatlaydi).
- `SetMaxIdleConns` — ishlatilmayotgan ulanishlarni **darhol yopib yubormaslik** uchun, keyingi so'rov tez kelsa, uni qayta ochishga hojat qolmasligi uchun.
- `SetConnMaxLifetime` — uzoq vaqt ochiq turgan ulanishlarni **davriy yangilash** uchun (masalan, tarmoq muammolari yoki bazaning o'zi ulanishni yopib qo'yishining oldini olish).

**Nega bu — restoran metaforasi bilan tushuniladi.** Har bir mijoz kelganda yangi ofitsiant yollash o'rniga, restoran **belgilangan sonli** ofitsiant yollaydi (`MaxOpenConns`) — ular mijozdan mijozga o'tib xizmat qiladi (qayta ishlatiladi), va agar hammasi band bo'lsa, yangi mijoz **navbatda kutadi**, yangi ofitsiant yollanmaydi.

**Havza to'lib qolsa nima bo'ladi?** Agar barcha `MaxOpenConns` ulanish band bo'lsa, yangi so'rov **navbatda kutadi**, toki biror ulanish bo'shamaguncha (yoki `context` orqali berilgan vaqt tugamaguncha xatolik qaytaradi). Shuning uchun `MaxOpenConns`ni juda kichik qilib qo'yish — dasturni **sekinlashtiradi** (so'rovlar navbatda tiqilib qoladi), juda katta qilib qo'yish esa — bazani **ortiqcha yuklaydi**. To'g'ri qiymat — bazangiz va serveringiz quvvatiga bog'liq.

**Go'da havza g'oyasini sinash — cheklangan sondagi "ishchi"larni qayta ishlatish:**

```go
type Havza struct {
	boshlar chan int
}

func YangiHavza(hajm int) *Havza {
	h := &Havza{boshlar: make(chan int, hajm)}
	for i := 0; i < hajm; i++ {
		h.boshlar <- i // hajm dona "bo'sh o'rin" bilan to'ldiramiz
	}
	return h
}

func (h *Havza) Ol() int {
	return <-h.boshlar // bo'sh o'rin bo'lmasa, shu yerda kutadi
}

func (h *Havza) Qaytar(id int) {
	h.boshlar <- id
}
```

Bu — "Buffered Channels" darsida ko'rgan g'oyaning to'g'ridan-to'g'ri qo'llanilishi: `chan int` — cheklangan sig'imli (`hajm` dona) navbat bo'lib xizmat qiladi, `Ol()` bo'sh o'rin bo'lmaguncha **bloklanadi** (xuddi haqiqiy connection pool'dagi kabi), `Qaytar()` esa ulanishni havzaga qaytaradi.

## EXAMPLE

```go
package main

import "fmt"

type Havza struct {
	boshlar chan int
}

func YangiHavza(hajm int) *Havza {
	h := &Havza{boshlar: make(chan int, hajm)}
	for i := 0; i < hajm; i++ {
		h.boshlar <- i
	}
	return h
}

func (h *Havza) Ol() int {
	return <-h.boshlar
}

func (h *Havza) Qaytar(id int) {
	h.boshlar <- id
}

func main() {
	h := YangiHavza(2)

	a := h.Ol()
	b := h.Ol()
	fmt.Println("olindi:", a, b)

	h.Qaytar(a)
	c := h.Ol() // a qaytarilgani uchun darhol oladi
	fmt.Println("qayta olindi:", c)
}
```

Natija (aniq raqamlar tartibi farq qilishi mumkin, lekin mantiq bir xil):

```
olindi: 0 1
qayta olindi: 0
```

## TASK

`Havza` strukturasi va uning `Ol()` methodi berilgan, lekin `Ol()` hali to'liq emas. Quyidagilarni to'ldiring:

`func (h *Havza) Ol() int` — `h.boshlar` kanalidan bitta qiymatni **oling va qaytaring** (`<-h.boshlar`). Agar kanal bo'sh bo'lsa, bu chaqiruv boshqa birov `Qaytar` chaqirmaguncha **kutadi** — buni o'zgartirish shart emas, kanal buni o'zi ta'minlaydi.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `return <-h.boshlar` — kanaldan qiymat olish uchun `<-` operatoridan foydalaning.
2. Bufferlangan kanal bo'sh bo'lsa, undan o'qish avtomatik ravishda **bloklanadi** — qo'shimcha tekshiruv shart emas.
