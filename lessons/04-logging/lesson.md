# 04 — Logging

## THEORY

`fmt.Println("xatolik yuz berdi")` — kichik dasturda yetarli, lekin production serverida, minglab log qatorlari orasida, **qaysi vaqtda**, **qaysi darajada** (info/warning/error), va **qanday qo'shimcha ma'lumot bilan** (foydalanuvchi ID, so'rov ID) xatolik yuz berganini bilish kerak. Oddiy matn logi buni qiyinlashtiradi.

**Yechim — structured logging (tuzilgan log).** Har bir log yozuvi, erkin matn o'rniga, **kalit-qiymat juftliklari** ko'rinishida yoziladi — bu, keyinchalik logni dastur yordamida (masalan, qidiruv tizimlarida) qidirish va filtrlashni osonlashtiradi.

**Go standart kutubxonasida — `log/slog` paketi** (Go 1.21+):

```go
import "log/slog"

slog.Info("foydalanuvchi yaratildi", "id", 42, "ism", "Ali")
slog.Error("bazaga ulanib bo'lmadi", "xato", err, "host", "localhost")
```

Standart holatda, bu matn ko'rinishida chiqadi:

```
time=2024-01-15T10:30:00.000Z level=INFO msg="foydalanuvchi yaratildi" id=42 ism=Ali
```

**JSON formatida chiqarish** — production'da, log yig'uvchi tizimlar (masalan, Elasticsearch, Loki) uchun qulay:

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("foydalanuvchi yaratildi", "id", 42, "ism", "Ali")
```

```json
{"time":"2024-01-15T10:30:00Z","level":"INFO","msg":"foydalanuvchi yaratildi","id":42,"ism":"Ali"}
```

**Log darajalari (log levels)** — har bir xabar, o'z **muhimlik darajasi** bilan belgilanadi:

- `Debug` — faqat rivojlantirish paytida kerak bo'lgan batafsil ma'lumot.
- `Info` — oddiy, kutilgan hodisalar ("server ishga tushdi", "foydalanuvchi kirdi").
- `Warn` — muammo emas, lekin e'tibor talab qiladigan holat ("keshga ulanish sekin").
- `Error` — haqiqiy xatolik ("bazaga ulanib bo'lmadi").

Production serverida, odatda faqat `Info` va undan yuqori darajadagi xabarlar chiqariladi — `Debug` xabarlar **o'chirilgan** bo'ladi, chunki ular juda ko'p va kerak bo'lmaydi.

**Terminalda ko'rish.** Kichik dastur yozib, uni ishga tushirsangiz:

```bash
go run main.go
```

konsolda darhol strukturalangan log qatorlarini ko'rasiz — har birida vaqt, daraja, xabar va qo'shimcha maydonlar.

**Nega bu — "Custom Errors" va "Error Handling" darslari bilan bog'liq.** Yaxshi logging, ko'pincha xatolik bilan birga qo'shimcha **kontekst** (masalan, qaysi foydalanuvchi, qaysi so'rov) qo'shadi — bu, keyinchalik muammoni tez topishni osonlashtiradi.

**Go'da darajalarni sinash uchun oddiy naqsh:**

```go
func logDarajaRaqami(daraja string) int {
	darajalar := map[string]int{"DEBUG": 0, "INFO": 1, "WARN": 2, "ERROR": 3}
	return darajalar[daraja]
}

func chiqarilsinmi(xabarDarajasi, minDaraja string) bool {
	return logDarajaRaqami(xabarDarajasi) >= logDarajaRaqami(minDaraja)
}
```

Bu — real logging kutubxonalari ichida sodir bo'ladigan **filtrlash mantig'ining** soddalashtirilgan ko'rinishi: agar xabar darajasi, o'rnatilgan minimal darajadan **past** bo'lsa, u chiqarilmaydi.

## EXAMPLE

```go
package main

import "fmt"

func logDarajaRaqami(daraja string) int {
	darajalar := map[string]int{"DEBUG": 0, "INFO": 1, "WARN": 2, "ERROR": 3}
	return darajalar[daraja]
}

func chiqarilsinmi(xabarDarajasi, minDaraja string) bool {
	return logDarajaRaqami(xabarDarajasi) >= logDarajaRaqami(minDaraja)
}

func main() {
	fmt.Println(chiqarilsinmi("DEBUG", "INFO")) // false — DEBUG, INFO'dan past
	fmt.Println(chiqarilsinmi("ERROR", "INFO")) // true — ERROR, INFO'dan yuqori
	fmt.Println(chiqarilsinmi("INFO", "INFO"))  // true — teng daraja chiqariladi
}
```

Natija:

```
false
true
true
```

## TASK

`chiqarilsinmi(xabarDarajasi, minDaraja string) bool` funksiyasini to'ldiring — u, `xabarDarajasi`ning `logDarajaRaqami` bo'yicha raqami, `minDaraja`ning raqamidan **katta yoki teng** bo'lsa, `true` qaytarsin.

`logDarajaRaqami` funksiyasi allaqachon yozilgan — undan foydalaning.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `return logDarajaRaqami(xabarDarajasi) >= logDarajaRaqami(minDaraja)`.
2. Xarita mavjud bo'lmagan kalit uchun `0` qaytaradi — noto'g'ri daraja nomi uchun ham funksiya xato bermaydi, shunchaki `DEBUG` kabi eng past daraja deb hisoblanadi.
