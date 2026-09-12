# 05 — Routing

## THEORY

"Web Development" bo'limidagi "Routing and Path Parameters" darsida `http.ServeMux` bilan asosiy yo'naltirishni ko'rgan edingiz. Katta API'larda esa yana bir muhim naqsh bor: **versiyalash (versioning)**. Kutubxonadagi kitoblarning turli nashrlari bo'lgani kabi (1-nashr, 2-nashr), API'lar ham vaqt o'tishi bilan o'zgaradi — va eski mijozlarni "sindirmasdan" yangilash uchun, manzillarga **versiya** qo'shiladi: `/api/v1/kitoblar`, `/api/v2/kitoblar`.

**Nega versiyalash kerak.** Agar API'ning ikkinchi versiyasida `Kitob` struct'ining shakli o'zgargan bo'lsa (masalan, yangi maydon qo'shilgan yoki eskisi o'chirilgan), lekin ba'zi mijozlar hali eski shaklga "o'rganib qolgan" bo'lsa — ikkalasini **bir vaqtda** ishlab turish kerak bo'ladi. Manzilga versiya qo'shish, buni oddiy va aniq qilib qo'yadi: `/api/v1/...` va `/api/v2/...` — ikkalasi ham alohida, mustaqil ishlaydi.

**Manzildan versiyani ajratib olish.** "Web Development" bo'limidagi Routing darsida ko'rgan `strings.TrimPrefix` naqshini davom ettiramiz, endi versiya qismini **tekshirish va ajratish** uchun:

```go
func apiVersiyaniAjrat(path string) (versiya, qolganPath string, ok bool) {
	if !strings.HasPrefix(path, "/api/") {
		return "", "", false
	}
	qism := strings.TrimPrefix(path, "/api/")
	boluklar := strings.SplitN(qism, "/", 2)
	if len(boluklar) < 1 || boluklar[0] == "" {
		return "", "", false
	}
	versiya = boluklar[0]
	if len(boluklar) == 2 {
		qolganPath = "/" + boluklar[1]
	}
	return versiya, qolganPath, true
}
```

Bu funksiya `/api/v1/kitoblar` manzilidan `"v1"` (versiya) va `"/kitoblar"` (qolgan yo'l)ni ajratib beradi — bu, "Multiple Return Values" darsida ko'rgan uchta qiymatni bir vaqtda qaytarish naqshi ("versiya, qolganPath, ok" — oxirgisi "Maps"dagi comma-ok naqshiga o'xshaydi).

**Nega bu, `http.ServeMux`ning o'ziga qo'shimcha qatlam.** `http.ServeMux` manzilni handler'ga **yo'naltiradi**, lekin manzil ichidan **ma'noli qismlarni** (bu holda, versiya) ajratib olish — alohida vazifa. Katta API'larda, bu ikkalasi birga ishlatiladi: avval `ServeMux` so'rovni tegishli handler guruhiga yo'naltiradi, keyin handler o'zi manzildan kerakli qismlarni (versiya, ID va h.k.) ajratib oladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"strings"
)

func apiVersiyaniAjrat(path string) (versiya, qolganPath string, ok bool) {
	if !strings.HasPrefix(path, "/api/") {
		return "", "", false
	}
	qism := strings.TrimPrefix(path, "/api/")
	boluklar := strings.SplitN(qism, "/", 2)
	if len(boluklar) < 1 || boluklar[0] == "" {
		return "", "", false
	}
	versiya = boluklar[0]
	if len(boluklar) == 2 {
		qolganPath = "/" + boluklar[1]
	}
	return versiya, qolganPath, true
}

func main() {
	v, p, ok := apiVersiyaniAjrat("/api/v1/kitoblar")
	fmt.Println(v, p, ok)
}
```

Natija:

```
v1 /kitoblar true
```

## TASK

`apiVersiyaniAjrat(path string) (versiya, qolganPath string, ok bool)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. Agar `path` `/api/` bilan boshlanmasa, `"", "", false` qaytarsin.
2. Aks holda, `/api/` dan keyingi birinchi bo'lakni `versiya` sifatida, qolganini (agar bo'lsa, `/` bilan boshlanib) `qolganPath` sifatida ajratib, `true` bilan qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `if !strings.HasPrefix(path, "/api/") { return "", "", false }`, keyin `qism := strings.TrimPrefix(path, "/api/")` va `boluklar := strings.SplitN(qism, "/", 2)`.
2. `versiya = boluklar[0]`, agar `len(boluklar) == 2` bo'lsa `qolganPath = "/" + boluklar[1]`, so'ng `return versiya, qolganPath, true`.
