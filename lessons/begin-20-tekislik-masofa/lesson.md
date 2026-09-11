# Tekislikda ikki nuqta orasidagi masofa

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Begin bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **9 daqiqa**

## EXAMPLE

```
Input: x1=0, y1=0, x2=3, y2=4
Output: 5.0
```

## TASK

Tekislikda ikki nuqta koordinatalari berilgan. Ular orasidagi masofani toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(x1 int, y1 int, x2 int, y2 int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. dx = x2 - x1, dy = y2 - y1
2. d = sqrt(dx*dx + dy*dy)
