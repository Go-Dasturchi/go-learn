# Matritsani transponirlash

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **8 daqiqa**

## EXAMPLE

```
Input: M=2, N=3, matrix=[[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
```

## TASK

M×N o'lchamli matritsa berilgan. Transponirlangan (N×M) matritsani chiqaring.

Quyidagi funksiyani to'ldiring:

```go
func transpose(M int, N int, matrix [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
