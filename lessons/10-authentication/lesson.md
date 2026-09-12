# 10 — Authentication

## THEORY

Konsertga kirishda, chiptangizni ko'rsatasiz — nazoratchi uni tekshirib, "bu haqiqiy chipta, kiring" deydi. **Authentication (autentifikatsiya)** — "siz kimsiz, buni isbotlang" degan savolga javob berish jarayoni. "Middleware" darsida ko'rgan oddiy tekshiruvni (`Authorization` sarlavhasi bo'sh yoki yo'q) endi bir qadam oldinga olib boramiz: sarlavhaning **aniq formatini** va **haqiqiy qiymatini** tekshirish.

**"Bearer" token formati — sanoat standarti.** Ko'plab API'lar `Authorization` sarlavhasini quyidagi formatda kutadi:

```
Authorization: Bearer <token>
```

`"Bearer "` (bo'shliq bilan) — prefiks, undan keyingi qism — haqiqiy token. Buni tahlil qilish uchun "Strings" darsida ko'rgan `strings.TrimPrefix`:

```go
func tokenniAjrat(sarlavha string) (string, bool) {
	if !strings.HasPrefix(sarlavha, "Bearer ") {
		return "", false
	}
	token := strings.TrimPrefix(sarlavha, "Bearer ")
	if token == "" {
		return "", false
	}
	return token, true
}
```

**Tokenni haqiqiy qiymat bilan solishtirish.** Eng sodda autentifikatsiya — tokenni, oldindan ma'lum bo'lgan "to'g'ri" qiymat bilan solishtirish (haqiqiy tizimlarda, bu odatda ma'lumotlar bazasida saqlangan, har bir foydalanuvchi uchun **individual** token bo'ladi, lekin g'oyaning o'zi bir xil):

```go
func authTekshir(sarlavha, togriToken string) bool {
	token, ok := tokenniAjrat(sarlavha)
	if !ok {
		return false
	}
	return token == togriToken
}
```

**Muhim xavfsizlik eslatmasi — vaqt bo'yicha hujum (timing attack).** Haqiqiy ishlab chiqarish tizimlarida, tokenlarni oddiy `==` bilan solishtirish, nazariy jihatdan xavfli hisoblanadi — chunki `==` odatda birinchi mos kelmagan belgida **darhol** to'xtaydi, va bu "qancha belgi to'g'ri kelgani"ni javob vaqtidan bilib olish imkonini (juda kam, lekin nolga teng bo'lmagan) beradi. Buning uchun `crypto/subtle` paketidagi `subtle.ConstantTimeCompare` ishlatiladi — bu, "Password Hashing" darsida yana uchraydigan mavzu.

**Nega bu — "JWT" darsining zaruriy poydevori.** Keyingi darsda ko'radigan JWT (JSON Web Token) — aslida, aynan shu `Authorization: Bearer <token>` naqshining, **o'z-o'zini tasdiqlaydigan** (o'zida imzo olib yuradigan) tokenlar bilan boyitilgan versiyasi. Shu darsda ko'rgan asosiy tuzilma — sarlavhani tahlil qilish, tokenni ajratish — JWT bilan ishlashda ham deyarli o'zgarmasdan qo'llaniladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
)

func tokenniAjrat(sarlavha string) (string, bool) {
	if !strings.HasPrefix(sarlavha, "Bearer ") {
		return "", false
	}
	token := strings.TrimPrefix(sarlavha, "Bearer ")
	if token == "" {
		return "", false
	}
	return token, true
}

func authTekshir(sarlavha, togriToken string) bool {
	token, ok := tokenniAjrat(sarlavha)
	if !ok {
		return false
	}
	return token == togriToken
}

func main() {
	fmt.Println(authTekshir("Bearer maxfiy-token-123", "maxfiy-token-123"))
	fmt.Println(authTekshir("Bearer notogri", "maxfiy-token-123"))
	fmt.Println(authTekshir("", "maxfiy-token-123"))
}
```

Natija:

```
true
false
false
```

## TASK

`tokenniAjrat` funksiyasi berilgan (allaqachon yozilgan). `authTekshir(sarlavha, togriToken string) bool` funksiyasini shunday to'ldiringki, u `tokenniAjrat` orqali tokenni ajratib, agar muvaffaqiyatli ajratilgan bo'lsa, uni `togriToken` bilan solishtirib, natijani qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `token, ok := tokenniAjrat(sarlavha)`, keyin `if !ok { return false }`.
2. `return token == togriToken`.
