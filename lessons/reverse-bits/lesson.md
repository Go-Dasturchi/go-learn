# Reverse Bits

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 43261596
Output: 964176192
Tushuntirish:
n'ning ikkilik ko'rinishi:  00000010100101000001111010011100
Teskari tartibi:             00111001011110000010100101000000
Teskari tartibdagi son:       964176192
```

## TASK

Berilgan `n` 32-bitli belgisiz (unsigned) butun sonining bitlarini teskari tartibga keltirgan holda (reverse) hosil bo'lgan yangi sonni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. 32 bitni teskari qilish uchun har bir bitni birma-bir o'qib, natijaga qarama-qarshi tomondan joylashtirishni o'ylang.
2. 32 marta tsikl yurgizing: har qadamda natijani chapga bir bitga suring va `n` ning eng past bitini natijaga qo'shing, so'ngra `n` ni o'ngga bir bitga suring — shunday qilib bitlar teskari tartibda yig'iladi.
