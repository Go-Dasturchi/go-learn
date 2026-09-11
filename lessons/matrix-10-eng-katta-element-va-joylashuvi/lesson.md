# Eng katta element va joylashuvi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **7 daqiqa**

## EXAMPLE

```
Input: M=3, N=3, matrix=[[1,2,3],[4,9,6],[7,8,5]]
Output: max=9, pos=(1,1)
```

## TASK

M×N o'lchamli matritsa berilgan. Eng katta elementni va uning (qator, ustun) koordinatasini toping.

Quyidagi funksiyani to'ldiring:

```go
func maxElementPosition(M int, N int, matrix [][]int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
