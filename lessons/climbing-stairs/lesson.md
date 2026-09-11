# Climbing stairs

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 2
Output: 2
Tushuntirish: 2 ta usul bor:
1. 1 qadam + 1 qadam
2. 2 qadam
```

## TASK

Siz zinadan yuqoriga chiqyapsiz. Yuqoriga yetish uchun `n` qadam (pog'ona) chiqishingiz kerak bo'ladi.

Har bir qadamda siz 1 ta yoki 2 ta pog'ona yuqoriga ko'tarilishingiz mumkin. Eng tepaga yetib borish uchun sizda jami nechta turli xil yo'l bor?

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu masala Fibonacci ketma-ketligiga juda o'xshaydi — har bir pog'onaga yetish yo'llari soni oldingi ikkita pog'onaga yetish yo'llari yig'indisiga teng.
2. dp[i] = dp[i-1] + dp[i-2] formulasidan foydalanib, faqat oxirgi ikkita qiymatni saqlagan holda pastdan yuqoriga hisoblang (dp[1]=1, dp[2]=2).
