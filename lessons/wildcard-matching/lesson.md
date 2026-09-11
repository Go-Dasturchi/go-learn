# Wildcard matching

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "aa", p = "*"
Output: true
Tushuntirish: '*' har qanday satrga to'g'ri keladi.
```

## TASK

`s` satri va `p` namunasi (pattern) berilgan. Quyidagi maxsus belgilarni inobatga olib solishtiruvchini (matcher) amalga oshiring:

- `?` har qanday yakka belgiga mos tushadi.
- `*` har qanday belgilar ketma-ketligiga (jumladan bo'shga) mos tushadi.
  Bu ham xuddi Regex kabi butun satrni qoplashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, p string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu ham 2D dinamik dasturlash masalasi (Regular Expression Matching ga o'xshash), lekin bu yerda `*` ancha sodda ma'noga ega — u nol yoki istalgan sondagi ISTALGAN belgilarga mos keladi, oldingi belgini takrorlamaydi.
2. dp[i][j] ni "s ning birinchi i ta belgisi p ning birinchi j ta belgisiga mos keladimi" deb belgilang. Agar p[j-1]=='*' bo'lsa, dp[i][j] = dp[i-1][j] (yulduzcha joriy s[i-1] ni ham "yutib yuboradi") YOKI dp[i][j-1] (yulduzcha bo'sh satrga mos keladi). Aks holda, agar p[j-1]=='?' yoki p[j-1]==s[i-1] bo'lsa dp[i][j]=dp[i-1][j-1]. Bazaviy holatda dp[0][j] faqat p ning boshi ketma-ket `*` lardan iborat bo'lsagina true bo'ladi.
