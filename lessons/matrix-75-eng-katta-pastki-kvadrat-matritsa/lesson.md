# Eng katta pastki kvadrat matritsa

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **13 daqiqa**

## EXAMPLE

```
Input: M=4, N=5, K=2, matrix=[[1,2,3,4,5],[2,3,4,5,6],[3,4,5,6,7],[4,5,6,7,8]]
Output: max=26, pos=(2,3)
```

## TASK

M×N o'lchamli matritsa berilgan. Yig'indi eng katta bo'lgan K×K pastki matritsani toping.

Quyidagi funksiyani to'ldiring:

```go
func maxSquareSubmatrix(M int, N int, K int, matrix [][]int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
