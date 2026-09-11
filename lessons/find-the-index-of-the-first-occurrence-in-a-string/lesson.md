# Find the Index of the First Occurrence in a String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Tushuntirish: "sad" so'zi haystack ichida 0 va 6-indekslarda uchraydi.
Birinchi marta uchragan joyi 0 bo'lgani uchun 0 ni qaytaramiz.
```

## TASK

Ikkita `haystack` va `needle` satrlari berilgan. `needle` satrining `haystack` satrida birinchi marta paydo bo'lgan joyining indeksini qaytaring. Agar `needle` satri `haystack` tarkibida bo'lmasa, `-1` ni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(haystack string, needle string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
