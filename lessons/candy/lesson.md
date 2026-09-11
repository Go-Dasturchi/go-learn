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

1. Har bir bolani bir vaqtning o'zida ikkala qo'shnisi bilan solishtirishga urinmang — masalani ikkita alohida o'tish (chapdan o'ngga va o'ngdan chapga) orqali yeching, har bir o'tishda faqat bitta tomondagi qoidani ta'minlang.
2. Avval barcha bolaga 1 tadan konfet bering. Chapdan o'ngga yurib, ratings[i] > ratings[i-1] bo'lsa candies[i] = candies[i-1]+1 qiling. Keyin o'ngdan chapga yurib, ratings[i] > ratings[i+1] bo'lsa candies[i] ni max(candies[i], candies[i+1]+1) ga tenglashtiring — bu ikkinchi o'tishda birinchi o'tishda qo'yilgan qiymatni kamaytirib yubormaslik uchun kerak.
