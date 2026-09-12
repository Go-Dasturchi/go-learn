# 09 — HTTP Status Codes

## THEORY

"Web Development" bo'limidagi "HTTP Status Codes" darsida asosiy kodlarni (`200`, `400`, `404`) qo'lda tanlashni ko'rgan edingiz. Katta dasturlarda esa, ko'pincha status kodni **qo'lda emas**, balki funksiyadan qaytgan **xatolik turiga qarab, avtomatik** tanlash kerak bo'ladi — "Error Handling" darsida ko'rgan sentinel xatoliklarni eslang.

**Xatolik turiga qarab status tanlash:**

```go
var ErrTopilmadi = errors.New("topilmadi")
var ErrNotogri = errors.New("notogri so'rov")

func xatoStatusi(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrTopilmadi):
		return http.StatusNotFound
	case errors.Is(err, ErrNotogri):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
```

Bu yerda "Switch" darsida ko'rgan **shartsiz `switch`** naqshi, "Error Handling" darsida ko'rgan `errors.Is` bilan birlashtirilgan — har bir `case` o'zining shartini tekshiradi, birinchi mos kelgani ishlaydi.

**Nega bu naqsh foydali.** Handler kodining o'zida har safar `if errors.Is(err, ErrTopilmadi) { w.WriteHeader(404) } else if ...` deb yozish o'rniga, bu mantiqni **bitta joyga** jamlab, handler'lar ichida shunchaki chaqirish mumkin:

```go
func kitobHandler(w http.ResponseWriter, r *http.Request) {
	_, err := kitobniTop(id)
	w.WriteHeader(xatoStatusi(err))
}
```

Bu — kodni ancha qisqartiradi, va **barcha** handler'lar bir xil xatoliklarga bir xil status kod bilan javob berishini kafolatlaydi (bitta handler `404`, boshqasi tasodifan `400` qaytarib qo'yish xavfini yo'q qiladi).

**`default: return http.StatusInternalServerError` — "bilmagan xatolik" uchun xavfsiz standart.** Agar xatolik **tanish sentinel'lardan hech biriga mos kelmasa**, bu — odatda **kutilmagan, dastur ichidagi** muammoni bildiradi, va bunday holatlarda `500 Internal Server Error` qaytarish — mijozga "bu sizning xatoingiz emas, bizning tomondan muammo" deb aniq signal beradi.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrTopilmadi = errors.New("topilmadi")
var ErrNotogri = errors.New("notogri so'rov")

func xatoStatusi(err error) int {
	switch {
	case err == nil:
		return http.StatusOK
	case errors.Is(err, ErrTopilmadi):
		return http.StatusNotFound
	case errors.Is(err, ErrNotogri):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	fmt.Println(xatoStatusi(nil))
	fmt.Println(xatoStatusi(ErrTopilmadi))
	fmt.Println(xatoStatusi(errors.New("kutilmagan")))
}
```

Natija:

```
200
404
500
```

## TASK

`ErrTopilmadi`, `ErrNotogri` sentinel xatoliklari va `xatoStatusi(err error) int` funksiyasi berilgan. Uni shunday to'ldiringki:

- `err == nil` bo'lsa → `http.StatusOK`
- `errors.Is(err, ErrTopilmadi)` bo'lsa → `http.StatusNotFound`
- `errors.Is(err, ErrNotogri)` bo'lsa → `http.StatusBadRequest`
- boshqa har qanday xatolik → `http.StatusInternalServerError`

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `switch { case err == nil: return http.StatusOK; case errors.Is(err, ErrTopilmadi): return http.StatusNotFound ... }`.
2. `default: return http.StatusInternalServerError` — barcha boshqa holatlar uchun.
