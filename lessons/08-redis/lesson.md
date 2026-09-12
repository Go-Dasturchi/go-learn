# 08 — Redis

## THEORY

PostgreSQL — ma'lumotni **diskda**, doimiy saqlash uchun ajoyib, lekin har bir so'rov, diskdan o'qish tufayli, nisbatan **sekin** (millisekundlar). Ba'zi ma'lumotlar (masalan, foydalanuvchi sessiyasi, tez-tez so'raladigan hisob-kitob natijasi) uchun, bu tezlik **yetarli emas**.

**Redis** — ma'lumotni **operativ xotirada (RAM)** saqlaydigan, juda tezkor ma'lumotlar bazasi. RAM'dan o'qish, diskdan o'qishdan **minglab marta tezroq** — shuning uchun Redis, ko'pincha **kesh (cache)** sifatida ishlatiladi: tez-tez so'raladigan ma'lumotni, PostgreSQL'dan har safar qidirish o'rniga, Redis'da saqlab qo'yish.

### 1-qadam: Redis'ni o'rnatish va ishga tushirish

**macOS:**

```bash
brew install redis
brew services start redis
```

**Ubuntu/Debian:**

```bash
sudo apt update
sudo apt install redis-server
sudo systemctl start redis-server
sudo systemctl enable redis-server
```

O'rnatilganini tekshirish:

```bash
redis-cli ping
```

Agar `PONG` javobi kelsa — Redis ishlayapti.

### 2-qadam: `redis-cli` orqali asosiy buyruqlar

```bash
redis-cli
```

Bu, interaktiv Redis konsolini ochadi:

```
127.0.0.1:6379> SET foydalanuvchi:1 "Ali"
OK
127.0.0.1:6379> GET foydalanuvchi:1
"Ali"
127.0.0.1:6379> EXPIRE foydalanuvchi:1 60
(integer) 1
127.0.0.1:6379> TTL foydalanuvchi:1
(integer) 58
127.0.0.1:6379> DEL foydalanuvchi:1
(integer) 1
127.0.0.1:6379> EXISTS foydalanuvchi:1
(integer) 0
```

- `SET kalit qiymat` — ma'lumotni saqlaydi.
- `GET kalit` — ma'lumotni o'qiydi.
- `EXPIRE kalit soniya` — ma'lumotga **muddat** (TTL — Time To Live) belgilaydi; shu vaqtdan keyin, ma'lumot **avtomatik o'chadi**.
- `TTL kalit` — qolgan vaqtni ko'rsatadi.
- `DEL kalit` — ma'lumotni qo'lda o'chiradi.

**Nega TTL — Redis'ning eng muhim xususiyati.** Kesh, "eskirgan" ma'lumotni saqlab qolmasligi kerak — masalan, foydalanuvchi ma'lumoti PostgreSQL'da yangilangan bo'lsa, lekin Redis'dagi eski nusxa hali ham qaytarilsa, bu **noto'g'ri natija** beradi. TTL, bu muammoni, ma'lumotni **ma'lum vaqtdan keyin avtomatik eskirtirish** orqali yumshatadi.

### 3-qadam: Go'dan Redis'ga ulanish

Buning uchun, standart kutubxonada Redis drayveri yo'q — eng mashhur tashqi kutubxona `github.com/redis/go-redis/v9`:

```bash
go get github.com/redis/go-redis/v9
```

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

ctx := context.Background()
err := rdb.Set(ctx, "foydalanuvchi:1", "Ali", 60*time.Second).Err()
qiymat, err := rdb.Get(ctx, "foydalanuvchi:1").Result()
```

`60*time.Second` — aynan `EXPIRE` buyrug'idagi TTL'ning Go'dagi ko'rinishi.

**Nega bu darsda haqiqiy Redis kutubxonasi ishlatilmaydi.** Bu kurs, imkon qadar Go standart kutubxonasi bilan cheklanadi, va har bir mashq, tarmoqsiz, izolyatsiyalangan muhitda tekshiriladi. Shuning uchun, Redis'ning **TTL g'oyasini**, xotiradagi oddiy struct orqali, "soat" (clock) ni **tashqaridan in'ektsiya qilib** (Dependency Injection darsini eslang!) sinaymiz — bu, real vaqtga bog'liq bo'lmagan, **deterministik** test yozish imkonini beradi:

```go
type Kesh struct {
	malumotlar map[string]string
	tugashVaqti map[string]int64
}

func (k *Kesh) Saqla(kalit, qiymat string, hozir int64, ttlSoniya int64) {
	k.malumotlar[kalit] = qiymat
	k.tugashVaqti[kalit] = hozir + ttlSoniya
}

func (k *Kesh) Ol(kalit string, hozir int64) (string, bool) {
	tugash, bor := k.tugashVaqti[kalit]
	if !bor || hozir >= tugash {
		return "", false // muddati o'tgan yoki mavjud emas
	}
	return k.malumotlar[kalit], true
}
```

`hozir` (joriy vaqt) — funksiyaga **argument sifatida** uzatiladi (`time.Now()` chaqirilmaydi) — bu, "Concurrency" bo'limidagi kabi, testda **ixtiyoriy vaqtni** simulyatsiya qilish imkonini beradi.

## EXAMPLE

```go
package main

import "fmt"

type Kesh struct {
	malumotlar  map[string]string
	tugashVaqti map[string]int64
}

func YangiKesh() *Kesh {
	return &Kesh{malumotlar: map[string]string{}, tugashVaqti: map[string]int64{}}
}

func (k *Kesh) Saqla(kalit, qiymat string, hozir, ttlSoniya int64) {
	k.malumotlar[kalit] = qiymat
	k.tugashVaqti[kalit] = hozir + ttlSoniya
}

func (k *Kesh) Ol(kalit string, hozir int64) (string, bool) {
	tugash, bor := k.tugashVaqti[kalit]
	if !bor || hozir >= tugash {
		return "", false
	}
	return k.malumotlar[kalit], true
}

func main() {
	kesh := YangiKesh()
	kesh.Saqla("ism", "Ali", 100, 60) // 100-vaqtda saqlandi, 60 soniyaga

	qiymat, bor := kesh.Ol("ism", 120) // 20 soniya o'tdi — hali amal qiladi
	fmt.Println(qiymat, bor)

	qiymat, bor = kesh.Ol("ism", 200) // 100 soniya o'tdi — muddati tugagan
	fmt.Println(qiymat, bor)
}
```

Natija:

```
Ali true
 false
```

## TASK

`Kesh.Ol(kalit string, hozir int64) (string, bool)` methodini to'ldiring:

1. `k.tugashVaqti[kalit]` ni comma-ok bilan tekshiring.
2. Agar kalit mavjud bo'lmasa, **yoki** `hozir >= tugash` bo'lsa (muddati tugagan), `"", false` qaytaring.
3. Aks holda, `k.malumotlar[kalit]`ni `true` bilan qaytaring.

`Saqla` methodi va `main()` funksiyasini o'zgartirish shart emas.

**Terminalda sinash uchun** (checker tekshirmaydi): Redis'ni o'rnatib, `redis-cli` orqali `SET`/`GET`/`EXPIRE`/`TTL` buyruqlarini bajarib ko'ring.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `tugash, bor := k.tugashVaqti[kalit]`, keyin `if !bor || hozir >= tugash { return "", false }`.
2. Aks holda `return k.malumotlar[kalit], true`.
