# Zigzag tartibda chiqarish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **11 daqiqa**

## EXAMPLE

```
Input: M=3, N=3, matrix=[[1,2,3],[4,5,6],[7,8,9]]
Output: 1,2,3,6,5,4,7,8,9
```

## TASK

M×N o'lchamli matritsa berilgan. Qatorlarni zigzag tartibda (1-chi qator chapdan o'ngga, 2-chi o'ngdan chapga) chiqaring.

Quyidagi funksiyani to'ldiring:

```go
func zigzagOrder(M int, N int, matrix [][]int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
