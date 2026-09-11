# Run-length encoding

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, String bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **13 daqiqa**

## EXAMPLE

```
Kiritish: S='aaabbc'
Chiqish: a3b2c1
```

## TASK

S satr berilgan. Run-length encoding qiling. Masalan: 'aaabbc' -> 'a3b2c1'

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. String metodlaridan foydalaning
2. Looplar yordamida belgilarni aylanib o'ting
