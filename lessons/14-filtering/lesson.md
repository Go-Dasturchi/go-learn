# 14 — Filtering

## THEORY

Onlayn do'konda "faqat 100,000 so'mdan arzon mahsulotlarni ko'rsat" yoki "faqat qizil ranglilarni ko'rsat" deb qidiruv qilganingizni tasavvur qiling — bu **filtering (filtrlash)**: to'liq ro'yxatdan, ma'lum **shartga mos** elementlarnigina tanlab olish. "Higher-Order Functions" darsida ko'rgan `filter` funksiyasini eslaysizmi? Bu dars, aynan shu g'oyani, **HTTP API** kontekstida — query parametrlar orqali boshqariladigan filtrlash sifatida qo'llaydi.

**Oddiy filtrlash funksiyasi:**

```go
func elementlarniFiltrlash(royxat []string, prefiks string) []string {
	natija := []string{}
	for _, element := range royxat {
		if strings.HasPrefix(element, prefiks) {
			natija = append(natija, element)
		}
	}
	return natija
}
```

Bu — "Higher-Order Functions" darsidagi `filter`ga juda o'xshaydi, faqat bu safar shart **tashqi funksiya sifatida emas**, balki to'g'ridan-to'g'ri **ichkarida** yozilgan — chunki filtrlash mezoni (`prefiks` bilan boshlanishi) aniq va o'zgarmas.

**API'da qanday ishlatiladi — query parametr orqali boshqariladigan filtr:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	prefiks := r.URL.Query().Get("prefiks") // "Query Parameters" darsini eslang
	natija := elementlarniFiltrlash(royxat, prefiks)
	json.NewEncoder(w).Encode(natija) // "JSON API" darsini eslang
}
```

Mijoz `/mahsulotlar?prefiks=kitob` deb so'rov yuborsa, faqat `"kitob"` bilan boshlanadigan elementlar qaytariladi — filtrlash mezoni **so'rovning o'zida**, dinamik tarzda keladi.

**Bo'sh filtr — "hech narsani filtrlamaslik" ma'nosini anglatishi kerak.** Muhim amaliy qoida: agar `prefiks` bo'sh bo'lsa (foydalanuvchi filtr bermagan bo'lsa), odatda **barcha** elementlar qaytarilishi kutiladi, filtrlanmagan holda emas, bo'sh natija emas. `strings.HasPrefix(element, "")` — har doim `true` qaytaradi (har qanday satr bo'sh prefiks bilan "boshlanadi" deb hisoblanadi), shuning uchun bu funksiya **allaqachon** to'g'ri ishlaydi — alohida "agar bo'sh bo'lsa" shartini yozishning hojati yo'q.

**Nega filtrlash va pagination birga ishlatiladi.** Real API'larda, odatda avval **filtrlanadi** (kerakli elementlar tanlanadi), so'ng natija **sahifalanadi** ("Pagination" darsini eslang) — bu ikkalasi birgalikda, katta ma'lumotlar to'plamidan aniq va boshqariladigan qismni olish imkonini beradi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
)

func elementlarniFiltrlash(royxat []string, prefiks string) []string {
	natija := []string{}
	for _, element := range royxat {
		if strings.HasPrefix(element, prefiks) {
			natija = append(natija, element)
		}
	}
	return natija
}

func main() {
	royxat := []string{"kitob-1", "kitob-2", "qalam-1", "daftar-1"}
	fmt.Println(elementlarniFiltrlash(royxat, "kitob"))
	fmt.Println(elementlarniFiltrlash(royxat, ""))
}
```

Natija:

```
[kitob-1 kitob-2]
[kitob-1 kitob-2 qalam-1 daftar-1]
```

## TASK

`elementlarniFiltrlash(royxat []string, prefiks string) []string` funksiyasi berilgan. Uni shunday to'ldiringki, u `royxat`dan faqat `prefiks` bilan boshlanadigan elementlarni o'z ichiga olgan yangi slice qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bo'sh slice bilan boshlang: `natija := []string{}`, keyin `for _, element := range royxat { ... }` bilan aylaning.
2. `if strings.HasPrefix(element, prefiks) { natija = append(natija, element) }`, tsikldan keyin `return natija`.
