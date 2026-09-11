# Regular expression matching

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "aa", p = "a*"
Output: true
Tushuntirish: '*' belgisi 'a' ni qayta takrorlash imkonini beradi.
```

## TASK

Kirish satri `s` va namuna `p` berilgan bo'lib, `.` va `*` belgilarini hisobga oluvchi regex tekshirgichni amalga oshiring:

- `.` (nuqta) har qanday bitta belgiga mos keladi.
- `*` (yulduzcha) undan oldingi belgining nol yoki undan ortiq marta uchrashiga mos keladi.
  Mos kelish satrning bir qismini emas, balki to'liq satrni qoplashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, p string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu 2D dinamik dasturlash masalasi — dp[i][j] ni "s ning birinchi i ta belgisi p ning birinchi j ta belgisiga mos keladimi" deb ta'riflang, `*` belgisini alohida holat sifatida ko'rib chiqing.
2. Agar p[j-1] == '*' bo'lsa, ikki variant mavjud: yulduzcha oldidagi belgini "nol marta" ishlatish (dp[i][j] = dp[i][j-2]) yoki agar p[j-2] joriy s[i-1] ga mos kelsa "yana bir marta" ishlatish (dp[i][j] = dp[i][j] || dp[i-1][j]). Aks holda, agar p[j-1] == '.' yoki p[j-1] == s[i-1] bo'lsa, dp[i][j] = dp[i-1][j-1]; bazaviy holatlarda dp[0][0]=true va dp[0][j] uchun p ning boshida "x*y*..." kabi nolga moslashadigan naqshlarni to'g'ri hisoblang.
