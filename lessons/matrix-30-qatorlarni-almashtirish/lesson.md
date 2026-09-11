# Qatorlarni almashtirish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **8 daqiqa**

## EXAMPLE

```
Input: M=3, N=3, K1=0, K2=2, matrix=[[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,8,9],[4,5,6],[1,2,3]]
```

## TASK

M×N o'lchamli matritsa va K1, K2 berilgan. K1 va K2 qatorlarini almashtiring.

Quyidagi funksiyani to'ldiring:

```go
func swapRows(M int, N int, K1 int, K2 int, matrix [][]int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
