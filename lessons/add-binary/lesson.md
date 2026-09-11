# Add Binary

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: a = "11", b = "1"
Output: "100"
Tushuntirish: 11 + 1 = 100 binar sanoq sistemasida.
```

## TASK

Sizga binar (faqat 0 va 1 dan iborat) satr ko'rinishidagi `a` va `b` berilgan. Ularning yig'indisini ham binar satr ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(a string, b string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Buni ikkala satrni songa aylantirib qo'shish emas (juda katta bo'lishi mumkin), balki qo'lda binar qo'shish kabi o'ylang — o'ngdan chapga xonalab qo'shish.
2. Ikkala satrning oxiridan boshlab ikkita ko'rsatkich yuriting, har qadamda mos xonalar va carry'ni qo'shib natijadagi bitni va yangi carry'ni hisoblang, so'ng hosil bo'lgan natijani teskari aylantiring.
