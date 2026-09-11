# Uchburchak perimetri va yuzi

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Begin bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: x1=0,y1=0,x2=4,y2=0,x3=0,y3=3
Output: P=12.0, S=6.0
```

## TASK

Uchburchakning uch uchining koordinatalari berilgan. Perimetri va yuzini toping.

Quyidagi funksiyani to'ldiring:

```go
func yuza(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Avval tomonlarni hisoblang
2. p = (a + b + c) / 2
3. S = sqrt(p * (p-a) * (p-b) * (p-c))
