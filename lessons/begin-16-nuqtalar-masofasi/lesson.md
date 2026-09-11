# Ikki nuqta orasidagi masofa

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Begin bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** · Ball: **5** · Taxminiy vaqt: **3 daqiqa**

## EXAMPLE

```
Input: x1=3, x2=10
Output: 7.0
```

## TASK

Haqiqiy o'qda ikki nuqta berilgan. Ular orasidagi masofani toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(x1 int, x2 int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Mutlaq qiymat: abs(x2 - x1)
2. Natija manfiy bo'lib qolmasligi uchun ayirmani mutlaq qiymatga (`math.Abs`) o'rab oling — `x2 - x1` manfiy chiqishi mumkin.
