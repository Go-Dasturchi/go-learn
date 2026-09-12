# 11 — JWT

## THEORY

"Authentication" darsida oddiy tokenni **saqlab qo'yilgan qiymat bilan solishtirish** orqali tekshirdik. Bu usulning bir kamchiligi bor: server har safar tokenni tekshirish uchun, uni qayerdadir (masalan, ma'lumotlar bazasida) "eslab qolishi" kerak. **JWT (JSON Web Token)** — bu muammoni hal qiladigan, **o'z-o'zini tasdiqlaydigan** token formati: token o'zi bilan o'zining "imzosi"ni olib yuradi, va server hech qayerga qaramasdan, faqat shu imzoni tekshirib, tokenning **haqiqiy va o'zgartirilmagan** ekanligiga ishonch hosil qila oladi.

**JWT'ning uch qismi:** `header.payload.signature` — nuqta bilan ajratilgan, har biri Base64 (aniqrog'i, URL-xavfsiz Base64) bilan kodlangan:

1. **Header** — odatda o'zgarmas, algoritm turini bildiradi (masalan, `{"alg":"HS256","typ":"JWT"}`).
2. **Payload** — haqiqiy ma'lumot ("claims"): foydalanuvchi ID'si, amal qilish muddati va h.k. ("JSON" darsida ko'rgan struct'larni JSON'ga kodlash shu yerda ishlatiladi).
3. **Signature (imzo)** — `header.payload` matnining, **maxfiy kalit** yordamida hisoblangan HMAC xeshi.

**Nega imzo muhim — soxtalashtirishning oldini oladi.** Agar kimdir payload'ni (masalan, "men — administrator" deb) o'zgartirsa, lekin maxfiy kalitni bilmasa, u **to'g'ri imzoni** hisoblab chiqara olmaydi. Server tokenni qabul qilganda, `header.payload`ning imzosini **o'zi qayta hisoblab**, kelgan imzo bilan solishtiradi — agar ular mos kelmasa, token o'zgartirilgan (yoki soxta) deb rad etiladi.

**Imzoni hisoblash — `crypto/hmac` va `crypto/sha256`:**

```go
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func imzoHisobla(malumot, maxfiyKalit string) string {
	h := hmac.New(sha256.New, []byte(maxfiyKalit))
	h.Write([]byte(malumot))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}
```

`hmac.New(sha256.New, kalit)` — HMAC-SHA256 algoritmi bilan ishlaydigan "xesh hisoblagich" yaratadi, `kalit` bilan "tuzlangan" holda. `h.Write(...)` — imzolanadigan matnni beradi, `h.Sum(nil)` — yakuniy xesh baytlarini qaytaradi. `base64.RawURLEncoding` — "Encoding/Decoding" darsida ko'rgan oddiy Base64'dan farqli, JWT standarti talab qiladigan, URL'da xavfsiz ishlatiladigan (`+`/`/` o'rniga `-`/`_`, va `=` to'ldiruvchisiz) shakl.

**Tekshirish — imzoni qayta hisoblab, solishtirish:**

```go
func tokenTogrimi(malumot, imzo, maxfiyKalit string) bool {
	kutilganImzo := imzoHisobla(malumot, maxfiyKalit)
	return kutilganImzo == imzo
}
```

**Nega JWT foydali.** Server foydalanuvchi haqida hech qanday holatni "eslab qolmasdan" (stateless — "HTTP Basics" darsida ko'rgan tushunchani eslang), faqat maxfiy kalitni bilgan holda, kelgan istalgan tokenning haqiqiyligini darhol tekshira oladi. Bu — ko'plab zamonaviy API'larning autentifikatsiya tizimining asosini tashkil etadi.

## EXAMPLE

```go
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func imzoHisobla(malumot, maxfiyKalit string) string {
	h := hmac.New(sha256.New, []byte(maxfiyKalit))
	h.Write([]byte(malumot))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func tokenTogrimi(malumot, imzo, maxfiyKalit string) bool {
	return imzoHisobla(malumot, maxfiyKalit) == imzo
}

func main() {
	imzo := imzoHisobla("salom", "sirli-kalit")
	fmt.Println(imzo)
	fmt.Println(tokenTogrimi("salom", imzo, "sirli-kalit"))
	fmt.Println(tokenTogrimi("salom", imzo, "boshqa-kalit"))
}
```

Natija:

```
51POYU7Lbt2byuqaeXtJXU--lY0bOU92CpChneE4y3g
true
false
```

## TASK

`imzoHisobla(malumot, maxfiyKalit string) string` funksiyasi berilgan. Uni `crypto/hmac`, `crypto/sha256` va `encoding/base64` (aniqrog'i, `base64.RawURLEncoding`) yordamida shunday to'ldiringki, u `malumot`ning `maxfiyKalit` bilan hisoblangan HMAC-SHA256 imzosini, URL-xavfsiz Base64 matn sifatida qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `h := hmac.New(sha256.New, []byte(maxfiyKalit))`, keyin `h.Write([]byte(malumot))`.
2. `return base64.RawURLEncoding.EncodeToString(h.Sum(nil))`.
