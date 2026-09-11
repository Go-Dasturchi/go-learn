# Interleaving String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Tushuntirish: "aadbbcbcac" ni "aabcc" va "dbbca" dan uzilmagan holda terib chiqish mumkin.
```

## TASK

Sizga uchta `s1`, `s2` va `s3` satrlari berilgan. `s3` satri `s1` va `s2` satrlarining elementlarini qandaydir ketma-ketlikda qorishib yuborish (interleave) orqali hosil bo'lganligini yoki yo'qligini aniqlang.

Ikki satrning qorishib yuborilishi deganda ularni bir qancha qismlarga bo'lib, keyin shu qismlarni navbatma-navbat birlashtirib chiqish tushuniladi. Bunda har bir satr ichidagi harflarning o'zaro ketma-ketlik tartibi o'zgarmaydi.

Quyidagi funksiyani to'ldiring:

```go
func solve(s1 string, s2 string, s3 string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu 2D dinamik dasturlash masalasi — dp[i][j] ni "s1 ning birinchi i ta va s2 ning birinchi j ta belgisidan s3 ning birinchi i+j ta belgisini hosil qilish mumkinmi" degan mantiqiy qiymat sifatida ta'riflang.
2. dp[i][j] true bo'ladi, agar (dp[i-1][j] true va s1[i-1]==s3[i+j-1]) yoki (dp[i][j-1] true va s2[j-1]==s3[i+j-1]) shartlaridan kamida bittasi bajarilsa. Bazaviy qator va ustunni (dp[i][0] va dp[0][j]) alohida to'ldirishni unutmang, chunki ular faqat bitta satrdan kelib chiqadi.
