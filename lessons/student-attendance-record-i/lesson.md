# Student Attendance Record I

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: s = "PPALLP"
Output: true
Tushuntirish: 1 ta A va ketma-ket LL bor. Shartlar bajarilgan.
```

## TASK

Sizga talabaning davomat yozuvini bildiruvchi `s` satri berilgan:

- `'A'` — Absent (Kelmagan)
- `'L'` — Late (Kechikkan)
- `'P'` — Present (Kelgan)

Agar talaba quyidagi shartlarning ikkalasi ham bajarilsa mukofotga loyiq:

1. Jami `'A'` (kelmagan) lar soni 2 dan **kam** bo'lsin.
2. Ketma-ket `'L'` (kechikkan) lar soni 3 ta yoki undan **ortiq** bo'lmasin.

Agar talaba mukofotga loyiq bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
