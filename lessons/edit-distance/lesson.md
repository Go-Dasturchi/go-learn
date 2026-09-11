# Edit distance

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: word1 = "horse", word2 = "ros"
Output: 3
Tushuntirish: horse -> rorse (h ni r ga) -> rose (r ni o'chirish) -> ros (e ni o'chirish)
```

## TASK

`word1` va `word2` satrlari berilgan. `word1` ni `word2` ga aylantirish uchun zarur minimal operatsiyalar sonini (Edit Distance) qaytaring.

Ruxsat etilgan amallar: bitta belgini qo'shish, o'chirish yoki boshqasiga almashtirish.

Quyidagi funksiyani to'ldiring:

```go
func solve(word1 string, word2 string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu ikki satrni solishtiruvchi klassik 2D dinamik dasturlash — dp[i][j] ni "word1 ning birinchi i ta va word2 ning birinchi j ta belgisini bir-biriga aylantirish uchun kerak bo'ladigan minimal amallar soni" sifatida belgilang.
2. Agar so'nggi belgilar teng bo'lsa (word1[i-1]==word2[j-1]), dp[i][j] = dp[i-1][j-1]. Aks holda dp[i][j] = 1 + min(dp[i-1][j-1] — almashtirish, dp[i-1][j] — o'chirish, dp[i][j-1] — qo'shish). Bazaviy qator/ustunlarni dp[i][0]=i va dp[0][j]=j deb sozlang.
