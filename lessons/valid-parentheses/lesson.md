# Valid parentheses

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "()"
Output: true
```

## TASK

Faqatgina `(`, `)`, `{`, `}`, `[` va `]` belgilaridan iborat `s` satri berilgan. Satr yaroqli ekanligini aniqlang.

Yaroqli satr shartlari:

1. Ochiq qavslar aynan o'sha turdagi yopiq qavslar bilan yopilishi kerak.
2. Ochiq qavslar to'g'ri tartibda yopilishi kerak.
3. Har bir yopiq qavsga mos keladigan ochiq qavs bo'lishi kerak.

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
