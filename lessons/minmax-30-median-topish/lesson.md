# Median topish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, MinMax bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **12 daqiqa**

## EXAMPLE

```
Input: N=5, nums=[3, 7, 2, 9, 1]
Output: 3
```

## TASK

N ta son berilgan. Medianni (o'rtasidagi qiymat) toping. Agar N juft bo'lsa, o'rtadagi ikkitasining o'rtacha.

Quyidagi funksiyani to'ldiring:

```go
func findMedian(N int, nums []int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha elementlarni ko'rib chiqing
2. Minimal/maksimal qiymatni saqlang
