# 02 — PostgreSQL

## THEORY

**Diqqat: bu dars boshqalardan farq qiladi.** Shu paytgacha barcha mashqlar `go-learn`ning o'zi tomonidan avtomatik tekshirilgan. Bu darsda esa siz **haqiqiy PostgreSQL dasturini** o'zingizning kompyuteringizga o'rnatib, uni ishga tushirib, **o'z terminalingizda** SQL buyruqlarini yozasiz — buni `go-learn` avtomatik tekshira olmaydi (chunki bu haqiqiy, tashqi dastur, izolyatsiyalangan test muhitida yo'q). Quyidagi qadamlarni **o'zingizning terminalingizda**, bosqichma-bosqich bajaring.

**PostgreSQL nima.** Bu — **relatsion ma'lumotlar bazasi boshqaruv tizimi (RDBMS)**: "SQL Basics" darsida ko'rgan jadvallar, so'rovlar va h.k.ni haqiqatan saqlab, boshqarib beruvchi dastur. U alohida **server** sifatida ishlaydi (fon rejimida, doimo ishlab turadi) va unga turli dasturlar (jumladan, bizning Go dasturimiz) **ulanib**, so'rov yuboradi.

### 1-qadam — O'rnatish

**macOS (Homebrew orqali):**

```bash
brew install postgresql@16
brew services start postgresql@16
```

**Ubuntu/Debian Linux:**

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

**Windows.** README'da tavsiya etilganidek, eng silliq yo'l — WSL (Windows Subsystem for Linux) o'rnatib, ichida yuqoridagi Ubuntu ko'rsatmalarini bajarish. (Yoki postgresql.org saytidan rasmiy Windows o'rnatuvchisini yuklab olish ham mumkin.)

**O'rnatilganini tekshirish:**

```bash
psql --version
```

Bu buyruq PostgreSQL versiyasini ko'rsatishi kerak (masalan, `psql (PostgreSQL) 16.x`).

### 2-qadam — Serverga ulanish

`psql` — PostgreSQL bilan ishlashning asosiy terminal vositasi (interaktiv "SQL konsoli"):

```bash
# macOS'da (brew o'rnatgandan keyin, odatda joriy foydalanuvchi nomi bilan):
psql postgres

# Linux'da (postgres — standart administrator foydalanuvchisi):
sudo -u postgres psql
```

Muvaffaqiyatli ulansangiz, terminal `postgres=#` kabi maxsus so'rov belgisiga o'zgaradi — endi siz `psql` ichidasiz.

### 3-qadam — Baza va foydalanuvchi yaratish

`psql` ichida (yuqoridagi qadamdan keyin):

```sql
CREATE DATABASE golearn_db;
CREATE USER golearn_user WITH PASSWORD 'parol123';
GRANT ALL PRIVILEGES ON DATABASE golearn_db TO golearn_user;
```

Endi joriy bazadan chiqib, yangisiga ulaning:

```sql
\c golearn_db
```

(`\c` — psql'ning maxsus "buyrug'i", oddiy SQL emas — shuning uchun oxirida nuqta-vergul `;` kerak emas.)

### 4-qadam — Jadval yaratish va ma'lumot bilan ishlash

Endi "SQL Basics" darsida ko'rgan buyruqlarni **haqiqatan** ishlatamiz:

```sql
CREATE TABLE foydalanuvchilar (
	id SERIAL PRIMARY KEY,
	ism TEXT NOT NULL,
	yosh INTEGER
);

INSERT INTO foydalanuvchilar (ism, yosh) VALUES ('Ali', 25);
INSERT INTO foydalanuvchilar (ism, yosh) VALUES ('Vali', 30);

SELECT * FROM foydalanuvchilar;
```

Oxirgi buyruq natijasida, terminalda haqiqiy jadval ko'rinishida ma'lumot chiqishi kerak — tabriklaymiz, siz endi haqiqiy ma'lumotlar bazasi bilan ishlayapsiz!

**Foydali `psql` maxsus buyruqlari (barchasi `\` bilan boshlanadi):**

```
\l          -- barcha bazalar ro'yxati
\dt         -- joriy bazadagi barcha jadvallar ro'yxati
\d foydalanuvchilar  -- jadvalning tuzilishini (ustunlar, turlar) ko'rsatish
\q          -- psql'dan chiqish
```

### 5-qadam — Ulanish satri (connection string)

Go dasturi PostgreSQL'ga ulanishi uchun, unga bitta **manzil satri (DSN — Data Source Name)** kerak bo'ladi — bu, barcha ulanish ma'lumotlarini (host, port, foydalanuvchi, parol, baza nomi) bitta joyga jamlaydi:

```
host=localhost port=5432 user=golearn_user password=parol123 dbname=golearn_db sslmode=disable
```

Keyingi darsda ("Database Connection") aynan shu satrni Go dasturi ichida qanday ishlatishni ko'ramiz.

## EXAMPLE

Ushbu darsning "amaliy" qismi — yuqoridagi qadamlarning o'zi, sizning terminalingizda. Go tomonida esa, keyingi darsga tayyorgarlik sifatida, DSN satrini to'g'ri qurishni mashq qilamiz:

```go
package main

import "fmt"

func dsnQur(host string, port int, user, parol, dbnomi string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, parol, dbnomi)
}

func main() {
	fmt.Println(dsnQur("localhost", 5432, "golearn_user", "parol123", "golearn_db"))
}
```

Natija:

```
host=localhost port=5432 user=golearn_user password=parol123 dbname=golearn_db sslmode=disable
```

## TASK

`dsnQur(host string, port int, user, parol, dbnomi string) string` funksiyasi berilgan. Uni shunday to'ldiringki, u `"host=<host> port=<port> user=<user> password=<parol> dbname=<dbnomi> sslmode=disable"` ko'rinishidagi PostgreSQL ulanish satrini (DSN) qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, parol, dbnomi)`.
2. Argumentlar tartibiga e'tibor bering — `host, port, user, parol, dbnomi`, aynan shu ketma-ketlikda.
