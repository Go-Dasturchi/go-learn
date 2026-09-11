# Saddle point topish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Matrix bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **12 daqiqa**

## EXAMPLE

```
Input: matrix=[[1,2,3],[4,5,6],[7,8,9]]
Output: value=7, pos=(2,0)
```

## TASK

M×N o'lchamli matritsa berilgan. Saddle point (qatorida min, ustunida max) toping. Saddle point - bu qatorida eng kichik, ustunida eng katta bo'lgan element.

Quyidagi funksiyani to'ldiring:

```go
func findSaddlePoint(matrix [][]int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki sikl kerak: qator va ustun uchun
2. Avval qatordagi minimumni toping, keyin u ustundagi maksimum ekanini tekshiring
