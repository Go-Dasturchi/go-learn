# Eng kichik va eng katta orasidagi sonlar

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, MinMax bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** · Ball: **5** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: N=5, nums=[3, 7, 2, 9, 1]
Output: 3
```

## TASK

N ta son berilgan. Eng kichik va eng katta orasidagi elementlar sonini toping.

Quyidagi funksiyani to'ldiring:

```go
func elementsBetweenMinMax(N int, nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Barcha elementlarni ko'rib chiqing
2. Minimal/maksimal qiymatni saqlang
