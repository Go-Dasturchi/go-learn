# 05 — HTTP Status Codes

## THEORY

Restoranga qo'ng'iroq qilib joy band qilmoqchi bo'lsangiz, sizga bir necha xil javob berilishi mumkin: "Albatta, band qildik" (hammasi joyida), "Kechirasiz, bugun band" (siz xato so'radingiz — bo'sh joy yo'q), yoki "Hozir liniya band, keyinroq qo'ng'iroq qiling" (bizning tomondan muammo). HTTP'da bunday "javob turkumlari" — **status kodlar** orqali ifodalanadi: bu — javobning **3 xonali son** ko'rinishidagi "natijasi".

**Eng ko'p uchraydigan status kodlar guruhlari:**

- **2xx — muvaffaqiyat**: `200 OK` (hammasi joyida), `201 Created` (yangi narsa yaratildi).
- **4xx — mijoz xatosi**: `400 Bad Request` (so'rov noto'g'ri tuzilgan), `404 Not Found` (so'ralgan narsa topilmadi).
- **5xx — server xatosi**: `500 Internal Server Error` (serverning o'zida kutilmagan xato yuz berdi).

Go'da bu kodlar `http` paketida tayyor konstantalar sifatida mavjud: `http.StatusOK`, `http.StatusBadRequest`, `http.StatusNotFound`, `http.StatusInternalServerError` va h.k. — raqamlarni yodlashning hojati yo'q.

**Status kodni o'rnatish — `w.WriteHeader(kod)`:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest) // 400
	fmt.Fprintln(w, "So'rov noto'g'ri")
}
```

**Muhim tartib qoidasi (yana): `WriteHeader` — birinchi bo'lib chaqiriladigan yozish amali bo'lishi kerak.** "JSON Responses" darsida sarlavhalar haqida ko'rgan qoidani eslang — xuddi shunday, `WriteHeader`ni javob tanasiga birinchi marta yozishdan **oldin** chaqirish kerak. Agar uni umuman chaqirmasangiz, Go avtomatik ravishda `200 OK` deb hisoblaydi (siz allaqachon shu narsani kuzatgan bo'lishingiz mumkin — oldingi darslarda `WriteHeader`ni hech qachon chaqirmagan edik, va status har doim `200` bo'lgan edi).

**Shartga qarab turli status qaytarish — amaliy naqsh:**

```go
func yoshTekshirHandler(w http.ResponseWriter, r *http.Request) {
	yosh := r.URL.Query().Get("yosh")
	if yosh == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "yosh parametri kerak")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Qabul qilindi")
}
```

**`http.Error` — xatolik uchun qulay yordamchi.** "Reading JSON Request Bodies" darsida ko'rganingizdek, `http.Error(w, xabar, kod)` — bitta chaqiruvda **status kodni ham, xabar matnini ham** birga yozib beradi, `WriteHeader` va `Fprintln`ni alohida yozishning qisqartmasi sifatida.

## EXAMPLE

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func yoshTekshirHandler(w http.ResponseWriter, r *http.Request) {
	yosh := r.URL.Query().Get("yosh")
	if yosh == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "yosh parametri kerak")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Qabul qilindi")
}

func main() {
	req1 := httptest.NewRequest("GET", "/tekshir?yosh=25", nil)
	rec1 := httptest.NewRecorder()
	yoshTekshirHandler(rec1, req1)
	fmt.Println(rec1.Code, rec1.Body.String())

	req2 := httptest.NewRequest("GET", "/tekshir", nil)
	rec2 := httptest.NewRecorder()
	yoshTekshirHandler(rec2, req2)
	fmt.Println(rec2.Code, rec2.Body.String())
}
```

Natija:

```
200 Qabul qilindi

400 yosh parametri kerak

```

## TASK

`yoshTekshirHandler(w http.ResponseWriter, r *http.Request)` funksiyasi berilgan. Uni shunday to'ldiringki:

1. `yosh` query parametrini o'qing. Agar u bo'sh bo'lsa — `http.StatusBadRequest` (400) status bilan `"yosh parametri kerak"` deb javob bering, va funksiyani shu yerda tugating.
2. Aks holda — `http.StatusOK` (200) status bilan `"Qabul qilindi"` deb javob bering.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `yosh := r.URL.Query().Get("yosh")`, keyin `if yosh == "" { w.WriteHeader(http.StatusBadRequest); fmt.Fprintln(w, "yosh parametri kerak"); return }`.
2. Shartdan keyin: `w.WriteHeader(http.StatusOK)`, `fmt.Fprintln(w, "Qabul qilindi")` — `WriteHeader` har doim yozishdan oldin chaqiriladi.
