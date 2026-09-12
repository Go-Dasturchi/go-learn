# 03 — Configuration

## THEORY

"Environment Variables" darsida `os.Getenv` yordamida muhit o'zgaruvchilarini o'qishni ko'rgan edik. Lekin haqiqiy dasturda, o'nlab sozlama (baza manzili, port, log darajasi, API kalitlari) bo'lishi mumkin — va ularni dastur bo'ylab **har xil joyda** `os.Getenv("DB_HOST")` deb chaqiraverish — xatoga moyil va tarqoq bo'ladi.

**Yechim — Configuration struct.** Barcha sozlamalarni **bitta joyda**, dastur boshida **bir marta** o'qib, struct ichiga jamlash:

```go
type Konfiguratsiya struct {
	Port        int
	DBHost      string
	DBPort      int
	LogDarajasi string
}

func KonfiguratsiyaYukla() Konfiguratsiya {
	return Konfiguratsiya{
		Port:        atoiYokiDefault(os.Getenv("PORT"), 8080),
		DBHost:      strYokiDefault(os.Getenv("DB_HOST"), "localhost"),
		DBPort:      atoiYokiDefault(os.Getenv("DB_PORT"), 5432),
		LogDarajasi: strYokiDefault(os.Getenv("LOG_LEVEL"), "info"),
	}
}
```

Endi dasturning qolgan qismi — muhit o'zgaruvchilari haqida **umuman bilmaydi**, faqat `Konfiguratsiya` structi bilan ishlaydi:

```go
cfg := KonfiguratsiyaYukla()
fmt.Println("Server port:", cfg.Port)
```

**Nega "default qiymat" (standart qiymat) muhim.** Har bir sozlama uchun muhit o'zgaruvchisi **berilmagan** bo'lishi mumkin (masalan, lokal rivojlantirishda `.env` fayl yo'q) — shu sababli, har bir sozlama uchun **oqilona standart qiymat** belgilash kerak, aks holda dastur, sozlama yo'qligi sababli, ishga tushmasdanoq qulab tushishi mumkin:

```go
func strYokiDefault(qiymat, standart string) string {
	if qiymat == "" {
		return standart
	}
	return qiymat
}
```

**Terminalda sinash.** Real muhitda, sozlamalarni ishga tushirishdan oldin belgilash mumkin:

```bash
export PORT=3000
export DB_HOST=prod-db.example.com
./mening-dasturim
```

Yoki bir martalik ishga tushirish uchun:

```bash
PORT=3000 DB_HOST=prod-db.example.com ./mening-dasturim
```

**Nega bu — "Dependency Injection" darsi bilan bog'liq.** `Konfiguratsiya` struct — o'zi ham bir turdagi bog'liqlik: u dastur boshida yaratiladi va kerakli komponentlarga **kiritiladi** (masalan, baza ulanish funksiyasiga `cfg.DBHost` uzatiladi), aynan `YangiServis(bog'liqlik)` konstruktoriga o'xshab.

**12-Factor App** — mashhur amaliyot bo'yicha, sozlamalar **kodga qattiq yozilmasligi** (hardcode qilinmasligi), balki muhit orqali berilishi kerak — bu, bir xil dastur kodini, hech qanday o'zgarishsiz, turli muhitlarda (development, staging, production) ishga tushirish imkonini beradi.

## EXAMPLE

```go
package main

import "fmt"

type Konfiguratsiya struct {
	Port   int
	DBHost string
}

func strYokiDefault(qiymat, standart string) string {
	if qiymat == "" {
		return standart
	}
	return qiymat
}

func intYokiDefault(qiymat string, standart int) int {
	if qiymat == "0" || qiymat == "" {
		return standart
	}
	var natija int
	fmt.Sscanf(qiymat, "%d", &natija)
	return natija
}

func KonfiguratsiyaYukla(muhit map[string]string) Konfiguratsiya {
	return Konfiguratsiya{
		Port:   intYokiDefault(muhit["PORT"], 8080),
		DBHost: strYokiDefault(muhit["DB_HOST"], "localhost"),
	}
}

func main() {
	// bo'sh muhit — standart qiymatlar ishlatiladi
	cfg1 := KonfiguratsiyaYukla(map[string]string{})
	fmt.Println(cfg1.Port, cfg1.DBHost)

	// berilgan qiymatlar ustunlik qiladi
	cfg2 := KonfiguratsiyaYukla(map[string]string{"PORT": "3000", "DB_HOST": "prod-db"})
	fmt.Println(cfg2.Port, cfg2.DBHost)
}
```

Natija:

```
8080 localhost
3000 prod-db
```

(Haqiqiy dasturda `muhit` — `os.Getenv` chaqiruvlaridan to'ldiriladi; bu yerda, testda ishlatish qulay bo'lishi uchun, `map[string]string` sifatida uzatilgan.)

## TASK

`KonfiguratsiyaYukla(muhit map[string]string) Konfiguratsiya` funksiyasini to'ldiring:

1. `muhit["PORT"]` mavjud va bo'sh bo'lmasa, uni `intYokiDefault` orqali songa aylantirib `Port` maydoniga qo'ying, aks holda `8080` standart qiymatini ishlating.
2. `muhit["DB_HOST"]` mavjud va bo'sh bo'lmasa, uni `DBHost` maydoniga qo'ying, aks holda `"localhost"` standart qiymatini ishlating.

Yordamchi funksiyalar (`strYokiDefault`, `intYokiDefault`) allaqachon yozilgan — ulardan foydalaning.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `Port: intYokiDefault(muhit["PORT"], 8080)` — mavjud bo'lmagan xarita kaliti, bo'sh qatorni qaytaradi, funksiyaning o'zi buni to'g'ri boshqaradi.
2. `DBHost: strYokiDefault(muhit["DB_HOST"], "localhost")`.
