# Diagonal bo'ylab aylantirish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **11 daqiqa**

## EXAMPLE

```
Input: N=3, matrix=[[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

## TASK

N×N kvadrat matritsa berilgan. Uni asosiy diagonal bo'ylab aylantiring (90 gradus).

Quyidagi funksiyani to'ldiring:

```go
func diagonalRotate(N int, matrix [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
