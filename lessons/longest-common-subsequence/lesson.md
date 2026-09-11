# Longest Common Subsequence

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: text1 = "abcde", text2 = "ace"
Output: 3
Tushuntirish: Eng uzun umumiy qism-ketma-ketlik "ace" bo'lib, uzunligi 3.
```

## TASK

Sizga ikkita `text1` va `text2` satrlari berilgan. Ularning eng uzun umumiy qism-ketma-ketligi (Longest Common Subsequence — LCS) uzunligini toping. Agar bunday qism-ketma-ketlik mavjud bo'lmasa, `0` qaytaring.

Qism-ketma-ketlik — bu satrdan ba'zi belgilarni o'chirish (yoki o'chirmaslik) orqali hosil bo'ladigan yangi satr bo'lib, qolgan belgilarning tartibi o'zgarmaydi.

Quyidagi funksiyani to'ldiring:

```go
func solve(text1 string, text2 string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ikki satrning har bir prefiks juftligi uchun "eng uzun umumiy qism-ketma-ketlik uzunligi"ni saqlaydigan ikki o'lchamli dinamik dasturlash (2D DP) jadvalini o'ylab ko'ring.
2. dp[i][j] — text1 ning birinchi i belgisi bilan text2 ning birinchi j belgisi orasidagi LCS uzunligi bo'lsin: agar text1[i-1]==text2[j-1] bo'lsa dp[i][j]=dp[i-1][j-1]+1, aks holda dp[i][j]=max(dp[i-1][j], dp[i][j-1]); javob dp[m][n] bo'ladi.
