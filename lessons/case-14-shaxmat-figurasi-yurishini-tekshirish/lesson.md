# Shaxmat figurasi yurishini tekshirish

## THEORY

Bu — amaliy mashq masalasi (Abramyan to'plamidan, Case bo'limi). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** · Ball: **15** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: fig=1, x1=1, y1=1, x2=1, y2=5
Output: true
```

## TASK

Shaxmat taxtasi (8×8), figura turi (1-Rook, 2-Bishop, 3-Queen, 4-Knight) va ikkita katakcha berilgan. Figura bir yurishda borishi mumkinmi?

Quyidagi funksiyani to'ldiring:

```go
func chessMove(fig int, x1 int, y1 int, x2 int, y2 int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. switch operatoridan foydalaning
2. Har bir case uchun break qo'yishni unutmang
