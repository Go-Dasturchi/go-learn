# Bulb Switcher

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 3
Output: 1
Tushuntirish:
Boshlash: [o'chiq, o'chiq, o'chiq]
1-tur: [yoniq, yoniq, yoniq]
2-tur: [yoniq, o'chiq, yoniq]
3-tur: [yoniq, o'chiq, o'chiq]
1 ta yoniq.
```

## TASK

Sizda `n` ta lampochka bor, barchasi dastlab o'chiq. Siz `n` ta tur (round) bajaryapsiz:

- 1-turda siz har bitta lampochkaning holatini o'zgartirasiz.
- 2-turda har ikkinchi lampochkaning holatini o'zgartirasiz.
- 3-turda har uchinchi lampochkaning holatini o'zgartirasiz.
- ...va `n`-turda faqat `n`-chi lampochkaning holatini o'zgartirasiz.

`n` ta turdan so'ng nechta lampochka yoniq qolganligini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
