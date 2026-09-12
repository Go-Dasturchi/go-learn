# 08 — Form Data

## THEORY

Veb-saytdagi oddiy HTML formani (masalan, "Ism" va "Familiya" maydonlari bo'lgan ro'yxatdan o'tish formasi) to'ldirib, "Yuborish" tugmasini bosganingizda, brauzer bu ma'lumotni ko'pincha **JSON emas**, balki maxsus "form" formatida yuboradi. Go'da bunday ma'lumotni o'qish uchun alohida, qulay usul bor.

**`r.FormValue("kalit")` — forma maydonini o'qish:**

```go
func royxatdanOtishHandler(w http.ResponseWriter, r *http.Request) {
	ism := r.FormValue("ism")
	fmt.Fprintf(w, "Salom, %s!", ism)
}
```

`FormValue` — "Query Parameters" darsida ko'rgan `r.URL.Query().Get(...)`ga juda o'xshaydi, lekin muhim farqi bor: `FormValue` ma'lumotni **ikkala** manbadan qidiradi — URL query parametrlaridan **va** so'rov tanasidagi form ma'lumotidan (agar `POST` so'rovi bo'lsa). Bu, ma'lumot qayerdan kelganidan qat'iy nazar, bitta qulay usulda o'qish imkonini beradi.

**Форма ma'lumotini sinash uchun to'g'ri sarlavha kerak.** `POST` so'rovi orqali form ma'lumotini yuborganda, so'rovga maxsus `Content-Type` sarlavhasi qo'yiladi — bu Go'ga tanani qanday "parslashni" bildiradi:

```go
tana := "ism=Ali&yosh=25" // "kalit=qiymat&kalit2=qiymat2" ko'rinishidagi format
req := httptest.NewRequest("POST", "/royxat", strings.NewReader(tana))
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
```

`application/x-www-form-urlencoded` — HTML formalar uchun standart format nomi; `&` — maydonlarni ajratuvchi belgi, `=` — kalit va qiymatni ajratuvchi belgi (URL query qismidagi formatning aynan o'zi, shuning uchun `FormValue` ikkalasini ham bir xilda tushuna oladi).

**Agar maydon berilmagan bo'lsa.** Xuddi `Query().Get(...)` kabi, `FormValue` ham mavjud bo'lmagan kalit uchun xato bermaydi — shunchaki bo'sh string qaytaradi, va "Query Parameters" darsida ko'rgan standart qiymat naqshini bu yerda ham qo'llash mumkin.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

func royxatdanOtishHandler(w http.ResponseWriter, r *http.Request) {
	ism := r.FormValue("ism")
	if ism == "" {
		ism = "Noma'lum"
	}
	fmt.Fprintf(w, "Ro'yxatga olindi: %s", ism)
}

func main() {
	tana := "ism=Ali"
	req := httptest.NewRequest("POST", "/royxat", strings.NewReader(tana))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()
	royxatdanOtishHandler(rec, req)

	fmt.Println(rec.Body.String())
}
```

Natija:

```
Ro'yxatga olindi: Ali
```

## TASK

`royxatdanOtishHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `r.FormValue("ism")` orqali `ism` forma maydonini o'qisin.
2. Agar bo'sh bo'lsa, `"Noma'lum"` deb qabul qilsin.
3. Javob sifatida `"Ro'yxatga olindi: <ism>"` matnini yozsin.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `ism := r.FormValue("ism")`, keyin `if ism == "" { ism = "Noma'lum" }`.
2. `fmt.Fprintf(w, "Ro'yxatga olindi: %s", ism)`.
