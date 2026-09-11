# Toq indeksli elementlar

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Array bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** · Ball: **5** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: N=6, arr=[1,2,3,4,5,6]
Output: [6,4,2]
```

## TASK

N ta elementli massiv berilgan. Toq indeksli elementlarni (1, 3, 5, ...) teskari tartibda chiqaring.

**Eslatma:** Indekslar 0 dan boshlanadi. Ya'ni toq indekslar: 1, 3, 5, ...

Quyidagi funksiyani to'ldiring:

```go
func oddIndexElements(N int, arr []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Toq indekslar: 1, 3, 5, ...
2. Oxirgi toq indeksdan boshlab aylanib chiqing
