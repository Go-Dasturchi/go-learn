# First Letter to Appear Twice

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: s = "abccbaacz"
Output: "c"
Tushuntirish:
'a' ning 2-paydo bo'lish indeksi 6.
'b' ning 2-paydo bo'lish indeksi 4.
'c' ning 2-paydo bo'lish indeksi 3.
Demak birinchi bo'lib 'c' ikkinchi marta uchradi.
```

## TASK

Sizga kichik ingliz harflaridan iborat `s` satri berilgan.

Satrda takrorlangan harflar orasida, **ikkinchi marta** eng qisqa vaqt (eng kichik indeks) ichida paydo bo'lgan birinchi harfni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
