# Pow(x, n)

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: x = 2.00000, n = 10
Output: 1024.00000
```

## TASK

Sizga o'nli kasr ko'rinishidagi `x` va butun son `n` berilgan. `x` sonini `n` chi darajaga ko'tarib qaytaring, ya'ni $x^n$ ni hisoblang.

Sizning yechimingiz tayyor `pow` funksiyasidan foydalanmasdan ishlashi tavsiya etiladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(x float64, n int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
