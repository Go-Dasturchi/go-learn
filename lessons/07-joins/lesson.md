# 07 — Joins

## THEORY

Ma'lumotlar bazasida ma'lumotni **bitta katta jadvalda emas**, balki bir necha **bog'liq jadvallarda** saqlash — yaxshi amaliyot hisoblanadi (bunga "normalizatsiya" deyiladi). Masalan, `foydalanuvchilar` jadvalida foydalanuvchilar, `buyurtmalar` jadvalida esa buyurtmalar saqlanadi — har bir buyurtma, qaysi foydalanuvchiga tegishli ekanligini ko'rsatuvchi `foydalanuvchi_id` ustuniga ega:

```sql
CREATE TABLE foydalanuvchilar (
    id SERIAL PRIMARY KEY,
    ism TEXT NOT NULL
);

CREATE TABLE buyurtmalar (
    id SERIAL PRIMARY KEY,
    foydalanuvchi_id INTEGER REFERENCES foydalanuvchilar(id),
    mahsulot TEXT NOT NULL,
    narx NUMERIC NOT NULL
);
```

`REFERENCES foydalanuvchilar(id)` — bu, **foreign key** (tashqi kalit): u `buyurtmalar.foydalanuvchi_id` ustunidagi har bir qiymat, `foydalanuvchilar.id`da **haqiqatan mavjud** bo'lishini kafolatlaydi (masalan, mavjud bo'lmagan foydalanuvchiga buyurtma yaratib bo'lmaydi).

**Ikkita jadvalni birlashtirib o'qish — `JOIN`:**

```sql
SELECT foydalanuvchilar.ism, buyurtmalar.mahsulot, buyurtmalar.narx
FROM buyurtmalar
JOIN foydalanuvchilar ON buyurtmalar.foydalanuvchi_id = foydalanuvchilar.id;
```

Bu so'rov, ikkala jadvaldagi ma'lumotni, `foydalanuvchi_id = id` shartiga mos qatorlarni **bog'lab**, bitta natija jadvaliga birlashtiradi. Natijada, har bir buyurtma qatorida, endi foydalanuvchining ismi ham ko'rinadi — go'yo ikkala jadval **birlashtirilgandek**.

**`JOIN` turlari:**

- `INNER JOIN` (yoki oddiy `JOIN`) — faqat **ikkala** jadvalda ham mos keladigan qatorlarni qaytaradi (buyurtmasi yo'q foydalanuvchilar ko'rinmaydi).
- `LEFT JOIN` — chap jadvaldagi **barcha** qatorlarni qaytaradi, mos kelmasa ham (buyurtmasi yo'q foydalanuvchilar ham ko'rinadi, lekin `buyurtmalar` ustunlari `NULL` bo'ladi):

```sql
SELECT foydalanuvchilar.ism, buyurtmalar.mahsulot
FROM foydalanuvchilar
LEFT JOIN buyurtmalar ON buyurtmalar.foydalanuvchi_id = foydalanuvchilar.id;
```

**Nega bu real hayotga o'xshaydi.** Bu — kutubxonadagi ikkita katalogni solishtirish kabi: bittasida "kitob raqami → kitob nomi", ikkinchisida "kitob raqami → qaysi javonda" saqlangan. Ikkalasini **kitob raqami** orqali bog'lab, "qaysi kitob qaysi javonda" degan savolga javob topasiz — aynan shu `JOIN`ning vazifasi.

**Go'da bu g'oyani sinash — ikkita xaritani umumiy kalit orqali birlashtirish:**

```go
func birlashtir(foydalanuvchilar map[int]string, buyurtmalar map[int]int) map[string]int {
	natija := make(map[string]int)
	for foydalanuvchiID, mahsulotSoni := range buyurtmalar {
		ism, bor := foydalanuvchilar[foydalanuvchiID]
		if bor {
			natija[ism] = mahsulotSoni
		}
	}
	return natija
}
```

Bu funksiya — `foydalanuvchilar` (id → ism) va `buyurtmalar` (id → mahsulot soni) xaritalarini, umumiy `id` kaliti orqali bog'lab, `ism → mahsulot soni` natijasini quradi. Bu — SQL `JOIN`ning, Go xaritalari yordamida qurilgan soddalashtirilgan ko'rinishi (aynan `INNER JOIN` kabi: faqat ikkalasida ham mavjud bo'lgan `id`lar natijaga tushadi).

## EXAMPLE

```go
package main

import "fmt"

func birlashtir(foydalanuvchilar map[int]string, buyurtmalar map[int]int) map[string]int {
	natija := make(map[string]int)
	for foydalanuvchiID, mahsulotSoni := range buyurtmalar {
		ism, bor := foydalanuvchilar[foydalanuvchiID]
		if bor {
			natija[ism] = mahsulotSoni
		}
	}
	return natija
}

func main() {
	foydalanuvchilar := map[int]string{1: "Ali", 2: "Vali"}
	buyurtmalar := map[int]int{1: 3, 2: 5, 99: 1} // 99 — mavjud bo'lmagan foydalanuvchi

	natija := birlashtir(foydalanuvchilar, buyurtmalar)
	fmt.Println(natija["Ali"], natija["Vali"], len(natija))
}
```

Natija:

```
3 5 2
```

`99` uchun buyurtma natijaga qo'shilmadi — chunki `foydalanuvchilar` xaritasida `99` mavjud emas, aynan `INNER JOIN`dagidek.

## TASK

`birlashtir(foydalanuvchilar map[int]string, buyurtmalar map[int]int) map[string]int` funksiyasini to'ldiring:

1. `buyurtmalar` xaritasi bo'ylab yuring (`foydalanuvchiID`, `mahsulotSoni`).
2. Har bir `foydalanuvchiID` uchun, `foydalanuvchilar` xaritasidan mos ismni **comma-ok** naqshi bilan qidiring.
3. Agar ism topilsa, natija xaritasiga `natija[ism] = mahsulotSoni` qo'shing. Topilmasa — o'tkazib yuboring (`INNER JOIN` kabi).

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `for foydalanuvchiID, mahsulotSoni := range buyurtmalar { ism, bor := foydalanuvchilar[foydalanuvchiID]; ... }`.
2. `if bor { natija[ism] = mahsulotSoni }` — `bor` yolg'on bo'lsa, hech narsa qo'shmang.
