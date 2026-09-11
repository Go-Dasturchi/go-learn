# Ikki matritsa ko'paytmasi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **9 daqiqa**

## EXAMPLE

```
Input: M=2, N=2, K=2, A=[[1,2],[3,4]], B=[[5,6],[7,8]]
Output: [[19,22],[43,50]]
```

## TASK

M×N va N×K o'lchamli ikkita matritsa berilgan. Ularning ko'paytmasini (M×K) toping.

Quyidagi funksiyani to'ldiring:

```go
func multiplyMatrices(M int, N int, K int, A [][]int, B [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
