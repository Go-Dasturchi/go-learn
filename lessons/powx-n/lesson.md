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

1. x ni n marta ketma-ket ko'paytirish o'rniga, "tez darajaga ko'tarish" (fast exponentiation / binary exponentiation) usulini o'ylab ko'ring — bu yerda daraja har safar ikkiga bo'linadi.
2. Agar n manfiy bo'lsa, x ni 1/x ga almashtirib n ni musbatga aylantiring; keyin n ni ikkiga bo'lib boruvchi siklda: agar n toq bo'lsa natijaga joriy x ni ko'paytiring, so'ng x ni o'ziga ko'paytiring (kvadratga ko'taring) va n ni ikkiga bo'ling (butun bo'lish) — n nolga tushguncha davom eting.
