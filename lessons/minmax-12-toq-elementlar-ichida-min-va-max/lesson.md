# Toq elementlar ichida min va max

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, MinMax bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** · Ball: **10** · Taxminiy vaqt: **7 daqiqa**

## EXAMPLE

```
Input: N=6, nums=[3, 8, 2, 9, 4, 11]
Output: min=3, max=11
```

## TASK

N ta son berilgan. Faqat toq sonlar ichidan eng kichik va eng kattasini toping.

Quyidagi funksiyani to'ldiring:

```go
func minMaxOdd(N int, nums []int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha elementlarni ko'rib chiqing
2. Minimal/maksimal qiymatni saqlang
