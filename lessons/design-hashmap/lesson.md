# Design HashMap

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

Test tizimi quyidagicha ishlaydi:

1. `put(1, 1)` ni bajar
2. `put(2, 2)` ni bajar
3. `get(1)` — 1 qaytarishi kerak
4. `get(3)` — -1 qaytarishi kerak

## TASK

Tayyor hash-map (Dictionary/HashMap) kutubxonalaridan foydalanmasdan quyidagi funksiyalarni amalga oshiring:

- `put(key, value)`: kalit-qiymat juftligini kiritadi. Agar kalit allaqachon mavjud bo'lsa, qiymatini yangilaydi.
- `get(key)`: berilgan kalitga mos qiymatni qaytaradi. Agar kalit mavjud bo'lmasa, `-1` qaytaradi.
- `remove(key)`: kalitni va unga mos qiymatni o'chiradi.

_Eslatma: Biz bu yerda testni soddalashtirish maqsadida funksiyalar yig'indisi sifatida operatsiyalarni ketma-ket bajaramiz va jarayonning natijasini tekshiramiz._

Quyidagi funksiyani to'ldiring:

```go
func solve(keys []int, values []int, queryKey int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
