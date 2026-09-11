# K dan kichik elementlar ichida maksimum

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, MinMax bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **8 daqiqa**

## EXAMPLE

```
Input: N=6, K=10, nums=[3, 15, 7, 20, 5, 12]
Output: 7
```

## TASK

N ta son va K berilgan. K dan kichik elementlar ichidan eng kattasini toping. Bo'lmasa 0.

Quyidagi funksiyani to'ldiring:

```go
func maxLessThanK(N int, K int, nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha elementlarni ko'rib chiqing
2. Minimal/maksimal qiymatni saqlang
