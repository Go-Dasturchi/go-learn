# First Unique Character in a String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "leetcode"
Output: 0
Tushuntirish: "l" harfi satrda faqat bir marta uchraydi va u birinchi bo'lib kelgani uchun uning indeksi 0.
```

## TASK

Sizga satr ko'rinishidagi `s` berilgan. Ushbu satr ichida bir marta qatnashgan (takrorlanmagan) eng birinchi belgini toping va uning indeksini qaytaring. Agar bunday belgi mavjud bo'lmasa, `-1` ni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
