# Number of 1 Bits

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: n = 11
Output: 3
Tushuntirish: 11 ning ikkilik ko'rinishi 00000000000000000000000000001011 bo'lib, unda 3 ta '1' bit bor.
```

## TASK

Berilgan `n` butun sonining ikkilik (binary) ko'rinishidagi `'1'` bitlarning umumiy sonini qaytaring (bu **Hamming Vazni** deb ham ataladi).

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Sonni ikkilik ko'rinishida "1" bitlarini sanash uchun uni qadam-baqadam 2 ga bo'lish yoki bitli operatsiyalar bilan tekshirishni o'ylang.
2. Son 0 bo'lmaguncha uni 2 ga bo'lib (yoki 1 bit o'ngga surib) boring, har qadamda eng past bit (`n % 2` yoki `n & 1`) 1 ga teng bo'lsa hisoblagichni oshiring.
