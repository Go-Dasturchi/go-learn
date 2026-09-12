# 07 — Docker Compose

## THEORY

Haqiqiy dastur kamdan-kam hollarda **yakka o'zi** ishlaydi — odatda unga PostgreSQL baza, Redis kesh, va boshqa xizmatlar ham kerak bo'ladi. Har birini **qo'lda**, alohida `docker run` buyruqlari bilan ishga tushirish — noqulay va xato qilish oson (portlarni, tarmoqni, tartibni eslab qolish kerak).

**Yechim — Docker Compose.** Bitta `docker-compose.yml` faylida, **barcha xizmatlarni** va ular orasidagi bog'liqlikni tasvirlash, keyin **bitta buyruq** bilan hammasini ishga tushirish mumkin.

### 1-qadam: docker-compose.yml yozish

```yaml
version: "3.9"

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=5432
    depends_on:
      - db

  db:
    image: postgres:16
    environment:
      - POSTGRES_USER=admin
      - POSTGRES_PASSWORD=parol123
      - POSTGRES_DB=mening_bazam
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

- `services` — har bir xizmat (bu yerda `app` — bizning Go dasturimiz, `db` — PostgreSQL) alohida bo'lim sifatida tasvirlangan.
- `depends_on` — `app` xizmati, `db` xizmatidan **keyin** ishga tushirilishini bildiradi (garchi bu, `db`ning to'liq **tayyor** bo'lishini emas, faqat **ishga tushirilganini** kafolatlaydi — buni keyinroq tushuntiramiz).
- `environment` — konteyner ichiga muhit o'zgaruvchilarini uzatadi ("Configuration" darsidagi `Konfiguratsiya` structi, aynan shu o'zgaruvchilarni o'qiydi).
- `volumes` — `pgdata` nomli **doimiy xotira**: konteyner o'chirilsa ham, baza ma'lumotlari **yo'qolmaydi** (chunki ular konteyner ichida emas, alohida saqlanadi).
- `DB_HOST=db` — muhim: konteynerlar ichida, boshqa xizmatga ularning **nomi orqali** murojaat qilinadi (`db`), `localhost` orqali emas — Docker Compose, xizmatlar orasida avtomatik ichki tarmoq (DNS) yaratadi.

### 2-qadam: hammasini ishga tushirish

```bash
docker compose up
```

Bu buyruq — `app` va `db` xizmatlarining ikkalasini ham quradi (agar kerak bo'lsa) va ishga tushiradi, konsolda **ikkalasining ham** logini birga ko'rsatadi.

Fonda ishga tushirish uchun:

```bash
docker compose up -d
```

Foydali qo'shimcha buyruqlar:

```bash
docker compose ps        # ishlab turgan xizmatlar holati
docker compose logs db   # faqat db xizmatining logi
docker compose down      # hammasini to'xtatish va o'chirish
docker compose down -v   # + volume'larni ham o'chirish (baza ma'lumoti yo'qoladi!)
```

### Nega `depends_on` yetarli emas — "tayyorlik" muammosi

`depends_on` faqat **ishga tushirish tartibini** kafolatlaydi, lekin PostgreSQL konteyneri ishga tushgandan keyin ham, **so'rov qabul qilishga tayyor bo'lishi uchun** bir necha soniya kerak bo'lishi mumkin. Agar `app` xizmati shu payt bazaga ulanishga urinsa — xatolik oladi.

**Bu — "Database Connection" darsida ko'rgan `kutibUlanish` (retry) naqshining aynan qo'llanilish joyi!** Go dasturi, ishga tushganda, bazaga **darhol** emas, balki **qayta-qayta urinib**, muvaffaqiyatli bo'lguncha yoki belgilangan urinishlar tugaguncha ulanishga harakat qilishi kerak — bu, docker-compose muhitida ishlashning **standart** yechimi hisoblanadi.

**Go'da xizmatlar tayyorligini tekshirish g'oyasi.** Compose faylidagi xizmatlarni ishga tushirishdan oldin, ularning kerakli bog'liqliklari (`depends_on`dagi ro'yxat) **allaqachon tayyor** ekanligini tekshirish mumkin:

```go
func bogliqXizmatlarTayyormi(tayyor map[string]bool, kerak []string) bool {
	for _, xizmat := range kerak {
		if !tayyor[xizmat] {
			return false
		}
	}
	return true
}
```

## EXAMPLE

```go
package main

import "fmt"

func bogliqXizmatlarTayyormi(tayyor map[string]bool, kerak []string) bool {
	for _, xizmat := range kerak {
		if !tayyor[xizmat] {
			return false
		}
	}
	return true
}

func main() {
	tayyor := map[string]bool{"db": true, "redis": false}

	fmt.Println(bogliqXizmatlarTayyormi(tayyor, []string{"db"}))          // true
	fmt.Println(bogliqXizmatlarTayyormi(tayyor, []string{"db", "redis"})) // false — redis tayyor emas
	fmt.Println(bogliqXizmatlarTayyormi(tayyor, []string{}))              // true — hech narsa kerak emas
}
```

Natija:

```
true
false
true
```

## TASK

`bogliqXizmatlarTayyormi(tayyor map[string]bool, kerak []string) bool` funksiyasini to'ldiring:

1. `kerak` ro'yxatidagi har bir xizmat nomi bo'ylab yuring.
2. Agar `tayyor` xaritasida shu xizmat `true` bo'lmasa (ya'ni tayyor emas), `false` qaytaring.
3. Hammasi tayyor bo'lsa (yoki `kerak` bo'sh bo'lsa), `true` qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

**Terminalda sinash uchun** (checker tomonidan tekshirilmaydi): yuqoridagi `docker-compose.yml`ni yozib, `docker compose up` bilan ishga tushiring, `docker compose ps` orqali holatini ko'ring, `docker compose down` bilan to'xtating.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `for _, xizmat := range kerak { if !tayyor[xizmat] { return false } }`.
2. Tsikldan keyin `return true` — agar hech qaysi xizmat "tayyor emas" bo'lib chiqmasa.
