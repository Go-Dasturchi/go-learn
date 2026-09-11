# Power of Two

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 1
Output: true
Tushuntirish: 2^0 = 1.
```

## TASK

Butun son `n` berilgan. Agar u 2 ning darajasi bo'lsa `true`, aks holda `false` qaytaring.

Yani, agar `x` butun son topilib `n == 2^x` tenglik bajarilsa, u holda `n` soni 2 ning darajasi hisoblanadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. 2 ning darajalari ikkilik ko'rinishda faqat bitta `1` bitiga ega bo'lishini eslang — bu holatni bitli amallar bilan tezda tekshirish mumkin.
2. `n` musbat bo'lishi va `n & (n-1)` ifodasi 0 ga teng bo'lishi kerak — chunki `n-1` operatsiyasi eng oxirgi 1 bitni va undan keyingi barcha nollarni "1" ga aylantiradi, shu sabab AND natijasi faqat yagona bit bo'lganda 0 bo'ladi.
