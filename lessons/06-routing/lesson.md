# 06 — Routing and Path Parameters

## THEORY

Katta do'kon bo'limlarga bo'lingan bo'ladi — "oziq-ovqat", "kiyim-kechak", "elektronika" — va har bir bo'limda o'ziga xos xodimlar ishlaydi. Veb-server ham xuddi shunday: turli **manzillar (path)** turli handler'lar tomonidan boshqariladi. Bu — **routing (yo'naltirish)** deb ataladi.

**`http.ServeMux` — manzillarni handler'larga taqsimlovchi "dispetcher".**

```go
mux := http.NewServeMux()
mux.HandleFunc("/salom", salomHandler)
mux.HandleFunc("/xayr", xayrHandler)
```

`mux` (multiplexer) — kelgan so'rovning manzilini (`r.URL.Path`) tekshirib, mos handler'ga "yo'naltiradi". Uni sinash uchun, "HTTP Handlers" darsida ko'rgan `httptest.NewRecorder()`ni ishlatib, so'rovni to'g'ridan-to'g'ri handler'ga emas, balki **`mux.ServeHTTP`**ga beramiz:

```go
req := httptest.NewRequest("GET", "/salom", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req) // mux o'zi to'g'ri handler'ni tanlaydi
```

**"Subtree" (pastki daraxt) manzillar — oxiri `/` bilan tugaydigan naqsh.** Agar manzil naqshi `/` bilan tugasa, u shu manzil va uning **barcha "pastki" manzillari** uchun ishlaydi:

```go
mux.HandleFunc("/foydalanuvchi/", foydalanuvchiHandler)
// /foydalanuvchi/42, /foydalanuvchi/ali, /foydalanuvchi/hokazo — hammasi shu handler'ga tushadi
```

**Manzildan parametrni qo'lda ajratib olish.** Yuqoridagi holatda, `foydalanuvchiHandler` qaysi aniq foydalanuvchi so'ralganini (masalan, `42`ni) bilishi kerak. Buning uchun manzilning "doimiy" qismini kesib tashlab, qolganini olish mumkin — "Strings" darsida ko'rgan `strings.TrimPrefix` bilan:

```go
func foydalanuvchiHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/foydalanuvchi/")
	fmt.Fprintf(w, "Foydalanuvchi ID: %s", id)
}
```

`strings.TrimPrefix(matn, prefiks)` — agar `matn` aynan `prefiks` bilan boshlansa, o'sha qismni olib tashlab, qolganini qaytaradi. `r.URL.Path` — `"/foydalanuvchi/42"` bo'lsa, `TrimPrefix` natijasi — `"42"`.

**Nega bu usul — barcha Go versiyalarida ishlaydigan, "portativ" yondashuv.** Go'ning ba'zi yangiroq versiyalarida `mux.HandleFunc("/foydalanuvchi/{id}", ...)` kabi maxsus sintaksis ham qo'shilgan, lekin u loyihaning `go.mod` faylida ko'rsatilgan Go versiyasiga qattiq bog'liq. `strings.TrimPrefix` orqali qo'lda ajratib olish esa — **har qanday** Go versiyasida, har qanday muhitda bir xilda ishlaydigan, universal usul, va aynan shu sababdan ko'plab mavjud (production) loyihalarda hali ham keng qo'llaniladi.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func salomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Salom!")
}

func foydalanuvchiHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/foydalanuvchi/")
	fmt.Fprintf(w, "Foydalanuvchi ID: %s", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/salom", salomHandler)
	mux.HandleFunc("/foydalanuvchi/", foydalanuvchiHandler)

	req1 := httptest.NewRequest("GET", "/salom", nil)
	rec1 := httptest.NewRecorder()
	mux.ServeHTTP(rec1, req1)
	fmt.Println(rec1.Body.String())

	req2 := httptest.NewRequest("GET", "/foydalanuvchi/42", nil)
	rec2 := httptest.NewRecorder()
	mux.ServeHTTP(rec2, req2)
	fmt.Println(rec2.Body.String())
}
```

Natija:

```
Salom!

Foydalanuvchi ID: 42
```

## TASK

`salomHandler` va `foydalanuvchiHandler` funksiyalari, va `yaratMux() *http.ServeMux` funksiyasi berilgan. `foydalanuvchiHandler`ni shunday to'ldiringki, u `r.URL.Path`dan `/foydalanuvchi/` prefiksini kesib tashlab, qolgan qismni (ID'ni) `"Foydalanuvchi ID: <id>"` ko'rinishida javob qilib yozsin. `yaratMux()` funksiyasini ham shunday to'ldiringki, u ikkala handler'ni ham (`"/salom"` va `"/foydalanuvchi/"` manzillariga) ro'yxatdan o'tkazilgan `*http.ServeMux` qaytarsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `foydalanuvchiHandler` ichida: `id := strings.TrimPrefix(r.URL.Path, "/foydalanuvchi/")`, keyin `fmt.Fprintf(w, "Foydalanuvchi ID: %s", id)`.
2. `yaratMux()` ichida: `mux := http.NewServeMux()`, `mux.HandleFunc("/salom", salomHandler)`, `mux.HandleFunc("/foydalanuvchi/", foydalanuvchiHandler)`, so'ng `return mux`.
