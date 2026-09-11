# Search a 2D Matrix

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

## TASK

Sizga `m x n` o'lchamdagi butun sonlardan iborat `matrix` matritsasi va bitta butun `target` soni berilgan. Matritsa quyidagi xususiyatlarga ega:

- Har bir qatordagi butun sonlar chapdan o'ngga qarab o'sib borish tartibida joylashgan.
- Har bir qatorning birinchi butun soni undan oldingi qatorning oxirgi butun sonidan katta.

Agar `target` matritsa ichida mavjud bo'lsa `true`, aks holda `false` qaytaring.

Yechimingiz `O(log(m * n))` vaqt murakkabligida ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(matrix [][]int, target int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Matritsani qator-ustun sifatida emas, balki bitta uzun saralangan massiv sifatida tasavvur qilib ko'ring — bu binary search qo'llashga imkon beradi.
2. lo=0, hi=rows*cols-1 oralig'ida odatdagi binary search yurgizing; har bir "tekis" mid indeksni `mid/cols` va `mid%cols` orqali qator va ustun indeksiga aylantirib, o'sha katakdagi qiymatni target bilan solishtiring.
