# Candy

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **35 daqiqa**

## EXAMPLE

```
Input: ratings = [1,0,2]
Output: 5
Tushuntirish: Siz birinchi, ikkinchi va uchinchi bolaga mos ravishda 2, 1, 2 ta konfet berasiz.
```

## TASK

Bir qatorda turgan `n` ta bola bor. Sizga har bir bolaning bahosini ko'rsatadigan `ratings` nomli butun sonlar massivi berilgan.

Siz bolalarga quyidagi qoidalarga ko'ra konfet tarqatishingiz kerak:

1. Har bir bola kamida bitta konfet olishi shart.
2. Yonidagi qo'shnisidan yuqori bahoga ega bo'lgan bola undan ko'proq konfet olishi shart.

Bolalarga tarqatilishi kerak bo'lgan eng kam konfetlar miqdorini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(ratings []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
