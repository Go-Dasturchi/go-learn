# Maximum Number of Vowels in a Substring of Given Length

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: s = "abciiidef", k = 3
Output: 3
Tushuntirish: "iii" qism-satri 3 ta unli harfdan iborat.
```

## TASK

Sizga kichik ingliz harflaridan iborat `s` satri va butun son `k` berilgan. `k` uzunlikdagi har qanday qism-satrida mavjud bo'lishi mumkin bo'lgan maksimal unli harflar sonini qaytaring.

Unli harflar: `'a'`, `'e'`, `'i'`, `'o'`, `'u'`.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, k int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Uzunligi qat'iy `k` bo'lgan har bir qism-satrni qaytadan sanash shart emas — belgilangan o'lchamli oynani (fixed-size sliding window) bir pozitsiyaga siljitganda unli harflar sonining qanday o'zgarishini o'ylab ko'ring.
2. Dastlabki `k` uzunlikdagi qism-satrdagi unli harflar sonini hisoblang; keyin oynani birma-bir o'ngga siljiting — har safar yangi qo'shilgan belgi unli bo'lsa hisoblagichni oshiring, oynadan chiqib ketgan belgi unli bo'lsa kamaytiring, va shu jarayonda eng katta hisoblagich qiymatini saqlab boring.
