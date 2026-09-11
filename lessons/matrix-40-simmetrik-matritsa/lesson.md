# Simmetrik matritsa

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **9 daqiqa**

## EXAMPLE

```
Input: N=3, matrix=[[1,2,3],[2,4,5],[3,5,6]]
Output: true
```

## TASK

N×N kvadrat matritsa berilgan. U simmetrikmi (A[i][j] = A[j][i]) tekshiring.

Quyidagi funksiyani to'ldiring:

```go
func isSymmetric(N int, matrix [][]int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
