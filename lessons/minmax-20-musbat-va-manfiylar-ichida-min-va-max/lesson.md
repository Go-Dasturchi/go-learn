# Musbat va manfiylar ichida min va max

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, MinMax bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **8 daqiqa**

## EXAMPLE

```
Input: N=6, nums=[3, -8, 2, -9, 10, -4]
Output: minPos=2, maxNeg=-4
```

## TASK

N ta son berilgan. Musbatlar ichidan min, manfiylar ichidan max ni toping.

Quyidagi funksiyani to'ldiring:

```go
func minMaxPosNeg(N int, nums []int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha elementlarni ko'rib chiqing
2. Minimal/maksimal qiymatni saqlang
