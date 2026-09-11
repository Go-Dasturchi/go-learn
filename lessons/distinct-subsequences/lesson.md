# Distinct Subsequences

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "rabbbit", t = "rabbit"
Output: 3
Tushuntirish:
Quyidagi tarzda s satridan 3 xil usulda "rabbit" so'zini hosil qilish mumkin:
rabbbit (uchinchi b ni o'chirish)
rabbbit (ikkinchi b ni o'chirish)
rabbbit (birinchi b ni o'chirish)
```

## TASK

Sizga `s` va `t` ikkita satr berilgan. `s` satrining qismlari (subsequence) ichida `t` satriga teng bo'lganlarining umumiy sonini qaytaring.

Satrning qismi (subsequence) bu asl satrdagi ba'zi belgilarni (yoki hech birini) o'chirish orqali hosil bo'ladigan yangi satrdir, qolgan belgilarni o'zaro joylashuv tartibini o'zgartirmagan holatda. Masalan, `"ACE"` bu `"ABCDE"` ning qismi hisoblanadi, ammo `"AEC"` emas.

Test keyslar javob 32-bitli butun songa (Int32) sig'ishiga moslab tuzilgan.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, t string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu ikki satr bo'yicha quriladigan 2D dinamik dasturlash masalasi — dp[i][j] ni "s ning birinchi i ta belgisi ichida t ning birinchi j ta belgisiga mos keluvchi subsekvensiyalar soni" sifatida ta'riflashni o'ylab ko'ring.
2. dp[i][j] = dp[i-1][j] (s[i-1] ni umuman ishlatmaslik varianti); agar s[i-1] == t[j-1] bo'lsa, bunga yana dp[i-1][j-1] ni ham qo'shing (s[i-1] ni t[j-1] ga moslashtirish varianti). Bazaviy holat sifatida barcha i uchun dp[i][0] = 1 qo'ying, chunki bo'sh satrga mos kelishning yagona yo'li bor.
