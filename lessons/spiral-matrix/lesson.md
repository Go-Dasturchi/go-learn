# Spiral Matrix

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

## TASK

Sizga `m x n` o'lchamdagi ikki o'lchamli `matrix` berilgan. Uning barcha elementlarini spiral tartibda (tashqaridan ichkariga qarab soat mili yo'nalishida) qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(matrix [][]int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Matritsaning tashqi chegaralarini (top, bottom, left, right) kuzatib, har safar bir tomonni to'liq aylanib chiqqach chegarani ichkariga siljitib borishni o'ylab ko'ring.
2. top<=bottom va left<=right bo'lguncha davom eting: avval yuqori qatorni chapdan o'ngga o'qib top++ qiling, keyin o'ng ustunni yuqoridan pastga o'qib right-- qiling, so'ng (agar hali qator qolgan bo'lsa) pastki qatorni o'ngdan chapga o'qib bottom-- qiling, va (agar hali ustun qolgan bo'lsa) chap ustunni pastdan yuqoriga o'qib left++ qiling.
