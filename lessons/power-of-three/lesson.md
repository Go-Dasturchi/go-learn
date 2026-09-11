# Power of Three

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 27
Output: true
Tushuntirish: 27 = 3^3
```

## TASK

Sizga butun son `n` berilgan. Agar u `3` ning darajasi bo'lsa `true` qaytaring, aks holda `false`.

Boshqacha aytganda, agar shunday bir butun `x` soni mavjud bo'lsaki, bunda $n == 3^x$ o'rinli bo'lsa, u holda `n` soni 3 ning darajasi hisoblanadi.

Sikllardan foydalanmasdan (rekursiya ham yo'q) buni hal qila olasizmi?

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Sikl yoki rekursiyasiz yechish uchun 3 ning eng katta darajasini (int chegarasida sig'adigan) oldindan bilishdan foydalanishni o'ylang.
2. Agar `n` musbat bo'lsa va int32 chegarasidagi eng katta `3^x` qiymati (`1162261467`) `n` ga qoldiqsiz bo'linsa, demak `n` ham 3 ning darajasi — chunki 3 tub son bo'lgani uchun uning darajalari faqat bir-birini bo'ladi.
