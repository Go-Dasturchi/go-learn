# Excel Sheet Column Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: columnTitle = "AB"
Output: 28
```

## TASK

Sizga Excel varag'ida ko'rinadigan ustun nomi `columnTitle` deb nomlangan satr berilgan. Unga mos keluvchi ustun raqamini qaytaring.

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
func solve(columnTitle string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu 26 lik sanoq sistemasiga o'xshaydi (lekin A=1 dan boshlanadi, 0 emas) — o'nlik sondagi kabi chapdan o'ngga har bir harfni 'siljitib' qo'shishni o'ylang.
2. Natijani 0 dan boshlang, har bir harf uchun natija = natija*26 + (harf - 'A' + 1) formulasini qo'llang.
