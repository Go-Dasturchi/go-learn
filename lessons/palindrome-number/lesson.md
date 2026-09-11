# Palindrome number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: x = 121
Output: true
```

## TASK

Sizga `x` butun soni berilgan. Agar `x` palindrom bo'lsa `true`, aks holda `false` qaytaring.

Palindrom son deb chapdan o'ngga ham, o'ngdan chapga ham bir xil o'qiladigan songa aytiladi. Masalan, 121 palindrom, lekin 123 emas.

Quyidagi funksiyani to'ldiring:

```go
func solve(x int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Sonni satrga aylantirmasdan, uning yarmini matematik yo'l bilan "teskari" qurishni o'ylang; manfiy sonlar va oxiri 0 (lekin o'zi 0 emas) bo'lgan sonlar darhol palindrom emasligini unutmang.
2. Sonning oxirgi raqamlarini ajratib olib, ulardan yangi (teskari) son quring, buni teskari son asl sondan katta yoki teng bo'lguncha davom ettiring; oxirida asl son (endi qisqargan) teskari songa yoki uning oxirgi raqamisiz qismiga teng bo'lsa, bu palindrom.
