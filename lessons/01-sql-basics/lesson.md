# 01 — SQL Basics

## THEORY

Katta, juda tartibli kutubxonani tasavvur qiling: har bir kitob aniq javon, aniq rafda, aniq tartib raqami bilan turadi. Kutubxonachiga "1990-yildan keyin yozilgan, muallifi Cho'lpon bo'lgan barcha she'riy kitoblarni ber" desangiz, u sizga aynan shu shartlarga mos kitoblarni topib beradi. **SQL (Structured Query Language)** — ma'lumotlar bazasiga aynan shunday "so'rov" berish uchun ishlatiladigan til. Bu darsda SQL'ning barcha asosiy buyruqlarini ko'rib chiqamiz — keyingi darsda esa ularni **haqiqiy PostgreSQL bazasida, o'z terminalingizda** ishlatasiz.

**Jadval (table) — Excel jadvaliga o'xshaydi.** Ma'lumotlar bazasidagi asosiy tuzilma — jadval: har bir **ustun (column)** bitta turdagi ma'lumotni ("Structs" darsida ko'rgan struct maydonlariga o'xshaydi), har bir **qator (row)** esa bitta yozuvni ifodalaydi:

```
foydalanuvchilar jadvali:
+----+----------+------+
| id | ism      | yosh |
+----+----------+------+
| 1  | Ali      | 25   |
| 2  | Vali     | 30   |
| 3  | Guli     | 22   |
+----+----------+------+
```

**Jadval yaratish — `CREATE TABLE`:**

```sql
CREATE TABLE foydalanuvchilar (
	id SERIAL PRIMARY KEY,
	ism TEXT NOT NULL,
	yosh INTEGER
);
```

- **`SERIAL PRIMARY KEY`** — bu ustun har bir qator uchun **avtomatik, noyob** son (1, 2, 3, ...) bo'lib, uni "asosiy kalit" (primary key) sifatida belgilaydi — har bir qatorni **noyob** aniqlash uchun ishlatiladi.
- **`TEXT`, `INTEGER`** — "Types" darsida ko'rgan `string`, `int` kabi, lekin SQL'ning o'z turlari.
- **`NOT NULL`** — bu ustun **hech qachon bo'sh (nolga teng)** bo'lmasligi kerakligini bildiradi — "Validation" darsida ko'rgan g'oyaning, ma'lumotlar bazasi darajasidagi ko'rinishi.

**Ma'lumot qo'shish — `INSERT INTO`:**

```sql
INSERT INTO foydalanuvchilar (ism, yosh) VALUES ('Ali', 25);
INSERT INTO foydalanuvchilar (ism, yosh) VALUES ('Vali', 30);
```

**Ma'lumot o'qish — `SELECT`, SQL'ning eng ko'p ishlatiladigan buyrug'i:**

```sql
SELECT * FROM foydalanuvchilar;                    -- barcha ustunlar, barcha qatorlar
SELECT ism, yosh FROM foydalanuvchilar;             -- faqat kerakli ustunlar
SELECT * FROM foydalanuvchilar WHERE yosh > 24;     -- shartga mos qatorlar ("If/Else" darsidagi shartga o'xshaydi)
SELECT * FROM foydalanuvchilar ORDER BY yosh DESC;  -- kamayish tartibida saralash ("Sorting" darsini eslang)
SELECT * FROM foydalanuvchilar LIMIT 2;             -- faqat birinchi 2 ta qator
SELECT * FROM foydalanuvchilar WHERE ism = 'Ali' OR ism = 'Vali'; -- "Boolean" darsidagi OR
```

**Ma'lumotni yangilash — `UPDATE`:**

```sql
UPDATE foydalanuvchilar SET yosh = 26 WHERE ism = 'Ali';
```

**Diqqat: `WHERE`siz `UPDATE`/`DELETE` — juda xavfli!** Agar `WHERE` shartini yozmasangiz, buyruq **jadvaldagi barcha qatorlarga** ta'sir qiladi — masalan, `UPDATE foydalanuvchilar SET yosh = 26;` (shartsiz) barcha foydalanuvchilarning yoshini 26 qilib qo'yadi. Bu — SQL o'rganishdagi eng muhim, eng ko'p uchraydigan xato manbai.

**Ma'lumotni o'chirish — `DELETE`:**

```sql
DELETE FROM foydalanuvchilar WHERE ism = 'Vali';
```

**Guruhlash va agregatsiya — `GROUP BY`, `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`:**

```sql
SELECT COUNT(*) FROM foydalanuvchilar;              -- jami nechta qator bor
SELECT AVG(yosh) FROM foydalanuvchilar;              -- o'rtacha yosh
SELECT yosh, COUNT(*) FROM foydalanuvchilar GROUP BY yosh; -- har bir yosh bo'yicha nechtadan bor
```

`GROUP BY` — "Maps" darsida ko'rgan g'oyaga o'xshaydi: qatorlarni ma'lum ustun qiymati bo'yicha "guruhlarga" bo'lib, har bir guruh uchun alohida hisoblash (son, yig'indi, o'rtacha) qiladi.

**Jadvalni o'zgartirish/o'chirish:**

```sql
ALTER TABLE foydalanuvchilar ADD COLUMN email TEXT; -- yangi ustun qo'shish
DROP TABLE foydalanuvchilar;                         -- butun jadvalni o'chirish (juda ehtiyot bo'ling!)
```

Keyingi darsda ("PostgreSQL") aynan shu buyruqlarning barchasini **haqiqiy, o'zingiz o'rnatgan bazada** ishlatasiz.

## EXAMPLE

Bu dars — SQL tilining o'zi haqida, shuning uchun "kodga" o'xshash misol yo'q. Buning o'rniga, keyingi Go mashqi — SQL so'rovini **matn (string) sifatida to'g'ri qurish**ga qaratilgan, chunki Go dasturi ma'lumotlar bazasiga aynan shunday, matn ko'rinishidagi so'rov yuboradi:

```go
package main

import "fmt"

func selectSorovQur(jadval, ustun string, qiymat int) string {
	return fmt.Sprintf("SELECT * FROM %s WHERE %s > %d", jadval, ustun, qiymat)
}

func main() {
	fmt.Println(selectSorovQur("foydalanuvchilar", "yosh", 24))
}
```

Natija:

```
SELECT * FROM foydalanuvchilar WHERE yosh > 24
```

## TASK

`selectSorovQur(jadval, ustun string, qiymat int) string` funksiyasi berilgan. Uni shunday to'ldiringki, u `"SELECT * FROM <jadval> WHERE <ustun> > <qiymat>"` ko'rinishidagi SQL so'rov matnini qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Sprintf("SELECT * FROM %s WHERE %s > %d", jadval, ustun, qiymat)` — "String Formatting" darsida ko'rgan joy egallovchilar.
2. `%s` — matn uchun, `%d` — son uchun, aynan shu tartibda.
