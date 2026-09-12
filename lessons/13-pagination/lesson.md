# 13 — Pagination

## THEORY

Katta kitobni bir vaqtning o'zida to'liq o'qib chiqib bo'lmaydi — u sahifalarga bo'lingan, va siz har safar bitta sahifani ochib o'qiysiz. **Pagination (sahifalash)** — API'da ham xuddi shu g'oya: agar ma'lumotlar to'plami katta bo'lsa (masalan, minglab foydalanuvchi), ularning **hammasini birdaniga** qaytarish o'rniga, kichik "sahifalar" bo'lib qaytariladi.

**Eng ko'p ishlatiladigan yondashuv — sahifa raqami va hajmi (page/pageSize):**

```go
func sahifalash(royxat []string, sahifa, hajmi int) []string {
	boshlanish := (sahifa - 1) * hajmi
	tugash := boshlanish + hajmi

	if boshlanish >= len(royxat) {
		return []string{} // so'ralgan sahifa mavjud emas
	}
	if tugash > len(royxat) {
		tugash = len(royxat) // oxirgi sahifa to'liq bo'lmasligi mumkin
	}
	return royxat[boshlanish:tugash]
}
```

Masalan, `hajmi = 10` bo'lsa: 1-sahifa — `[0:10]`, 2-sahifa — `[10:20]`, va hokazo. `(sahifa - 1) * hajmi` formulasi — "Operators" darsida ko'rgan oddiy arifmetikadan foydalanib, har bir sahifaning slice'dagi boshlang'ich indeksini hisoblaydi.

**Ikkita chegara holatini tekshirish — muhim.** "Slices" va "Slice Indexing" darslarida ko'rgan chegaradan chiqib ketish xatolarini eslang: agar `boshlanish` butun ro'yxat uzunligidan katta yoki teng bo'lsa (foydalanuvchi mavjud bo'lmagan sahifani so'ragan), yoki `tugash` ro'yxat uzunligidan katta chiqib ketsa (oxirgi, to'liq bo'lmagan sahifa) — bularni oldindan tekshirib, mos ravishda tuzatish kerak, aks holda dastur **panic** qilib qulab tushishi mumkin.

**API'da qanday ishlatiladi.** "Query Parameters" darsida ko'rgan naqshni eslang — haqiqiy API'da `sahifa` va `hajmi` odatda URL query parametrlari sifatida keladi:

```
GET /foydalanuvchilar?sahifa=2&hajmi=10
```

Handler ichida bu qiymatlar o'qilib (`strconv.Atoi` bilan songa aylantirilib), keyin aynan shu `sahifalash` funksiyasiga uzatiladi.

**Nega pagination muhim.** Agar API minglab yozuvni **bir so'rovda** qaytarsa, bu ham serverga (katta hajmdagi ma'lumotni tayyorlash), ham tarmoqqa (uzatish), ham mijozga (qayta ishlash) og'ir yuk beradi. Sahifalash, har bir so'rovni **kichik va boshqariladigan** qilib, umumiy tizimni ancha samaraliroq va tezroq qiladi.

## EXAMPLE

```go
package main

import "fmt"

func sahifalash(royxat []string, sahifa, hajmi int) []string {
	boshlanish := (sahifa - 1) * hajmi
	tugash := boshlanish + hajmi

	if boshlanish >= len(royxat) {
		return []string{}
	}
	if tugash > len(royxat) {
		tugash = len(royxat)
	}
	return royxat[boshlanish:tugash]
}

func main() {
	royxat := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(sahifalash(royxat, 1, 3)) // [a b c]
	fmt.Println(sahifalash(royxat, 2, 3)) // [d e f]
	fmt.Println(sahifalash(royxat, 3, 3)) // [g] — oxirgi, to'liq bo'lmagan sahifa
}
```

Natija:

```
[a b c]
[d e f]
[g]
```

## TASK

`sahifalash(royxat []string, sahifa, hajmi int) []string` funksiyasi berilgan. Uni shunday to'ldiringki, u `royxat`ning `sahifa`-raqamli (1dan boshlab), har biri `hajmi` ta elementdan iborat sahifasini qaytarsin. Agar so'ralgan sahifa mavjud bo'lmasa, bo'sh slice qaytaring; agar oxirgi sahifa to'liq bo'lmasa, mavjud qismini qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `boshlanish := (sahifa - 1) * hajmi`, `tugash := boshlanish + hajmi`, keyin `if boshlanish >= len(royxat) { return []string{} }`.
2. `if tugash > len(royxat) { tugash = len(royxat) }`, so'ng `return royxat[boshlanish:tugash]`.
