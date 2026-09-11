# Ugly Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 6
Output: true
Tushuntirish: 6 ning tub ko'paytuvchilari 2 va 3 dir (6 = 2 * 3).
```

## TASK

Sizga butun son `n` berilgan. Agar u xunuk son (ugly number) bo'lsa `true` qaytaring, aks holda `false`.

**Xunuk son** deb tub ko'paytuvchilari (prime factors) faqat `2`, `3` va `5` bo'lgan musbat butun songa aytiladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
