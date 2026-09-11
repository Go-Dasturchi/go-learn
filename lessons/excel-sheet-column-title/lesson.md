# Excel Sheet Column Title

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: columnNumber = 28
Output: "AB"
```

## TASK

Sizga butun son ko'rinishidagi `columnNumber` berilgan. Uning Excel varag'ida qanday ustun nomi sifatida paydo bo'lishini ko'rsatuvchi satrni qaytaring.

Masalan:
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Quyidagi funksiyani to'ldiring:

```go
func solve(columnNumber int) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu 26 lik sanoq sistemasiga o'tkazishga o'xshaydi, lekin oddiy qoldiq olish usuli ishlamaydi chunki bu yerda 0 emas 1 dan boshlanadi — sonni kamaytirib olishni o'ylab ko'ring.
2. columnNumber 0 dan katta ekan, avval uni 1 ga kamaytiring, so'ng (columnNumber % 26) orqali harfni toping va columnNumber'ni 26 ga bo'ling; hosil bo'lgan harflarni yig'ib, keyin ularni teskari aylantiring.
