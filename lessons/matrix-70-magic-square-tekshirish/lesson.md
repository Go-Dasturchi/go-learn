# Magic Square tekshirish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **12 daqiqa**

## EXAMPLE

```
Input: N=3, matrix=[[2,7,6],[9,5,1],[4,3,8]]
Output: true
```

## TASK

N×N kvadrat matritsa berilgan. Magic Square (barcha qator, ustun va diagonallar yig'indisi teng) ekanligini tekshiring.

Quyidagi funksiyani to'ldiring:

```go
func isMagicSquare(N int, matrix [][]int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
