# Sqrt(x)

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: x = 4
Output: 2
Tushuntirish: 4 ning kvadrat ildizi 2 ga teng.
```

## TASK

Sizga manfiy bo'lmagan butun son `x` kiritiladi. `x` ning kvadrat ildizini butun songa yaxlitlab qaytaring (kasr qismi tashlab yuboriladi).

**Izoh:** `sqrt` kabi tayyor darajaga ko'taruvchi yoki ildiz chiqaruvchi funksiyalardan foydalanish tavsiya etilmaydi.

Quyidagi funksiyani to'ldiring:

```go
func solve(x int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `sqrt` funksiyasidan foydalanish mumkin bo'lmagani uchun javobni `1` dan `x` gacha bo'lgan oraliqda binary search bilan qidirishni o'ylang.
2. `lo=1`, `hi=x/2` oralig'ida binary search yurgizing: o'rtadagi qiymat `mid <= x/mid` shartini qanoatlantirsa, uni nomzod javob sifatida saqlab qidiruvni yuqoriga davom ettiring, aks holda pastga qarab torayting — oxirida saqlangan eng katta mos nomzod javob bo'ladi.
