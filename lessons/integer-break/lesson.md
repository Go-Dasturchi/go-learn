# Integer Break

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: n = 10
Output: 36
Tushuntirish: 10 = 3 + 3 + 4, ko'paytma = 3 * 3 * 4 = 36.
```

## TASK

Sizga `n` butun soni berilgan. Uni kamida ikkita musbat butun sonlarga bo'ling va ularning ko'paytmasini maksimal qiling.

Maksimal ko'paytmani qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir sonni "birinchi bo'lak + qolgan qismning eng yaxshi ko'paytmasi" ko'rinishida ifodalab, kichikroq sonlar uchun javoblarni saqlab boruvchi dinamik dasturlash (DP) haqida o'ylab ko'ring.
2. dp[1..3] uchun boshlang'ich qiymatlarni to'g'ridan-to'g'ri belgilang; keyin har bir i=4..n uchun barcha j=1..i-1 bo'linishlarini sinab, j * dp[i-j] ning eng kattasini dp[i] sifatida saqlang (bu yerda dp[i-j] o'zi ham qolgan qismni yana bo'lish natijasi bo'lishi mumkin).
