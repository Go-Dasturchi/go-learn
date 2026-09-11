# Ketma-ketlikning yaqinlashuvi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, While bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **11 daqiqa**

## EXAMPLE

```
Input: eps=0.1
Output: K va Ak-1, Ak
```

## TASK

ε > 0 berilgan. A1=2, Ak=2+1/Ak-1 ketma-ketlikda |Ak - Ak-1| < ε bo'ladigan birinchi K ni toping.

Quyidagi funksiyani to'ldiring:

```go
func sequenceConverge(eps float64) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. while siklidan foydalaning
2. Shart to'g'ri bo'lguncha takrorlang
