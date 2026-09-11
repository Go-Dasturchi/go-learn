# Har bir ustun yig'indisi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** · Ball: **5** · Taxminiy vaqt: **6 daqiqa**

## EXAMPLE

```
Input: M=3, N=3, matrix=[[1,2,3],[4,5,6],[7,8,9]]
Output: 12, 15, 18
```

## TASK

M×N o'lchamli matritsa berilgan. Har bir ustun elementlari yig'indisini chiqaring.

Quyidagi funksiyani to'ldiring:

```go
func columnSums(M int, N int, matrix [][]int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Index bilan ishlaganda ehtiyot bo'ling
