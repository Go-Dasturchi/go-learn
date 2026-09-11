# Permutation in String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Tushuntirish: s2 ning "ba" qism-satri s1 ning permutatsiyalaridan biridir.
```

## TASK

Sizga ikkita `s1` va `s2` satrlari berilgan. Agar `s1` ning ixtiyoriy permutatsiyasi (harflar joyini almashtirib hosil qilingan shakl) `s2` satrining qism-satri (substring) bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s1 string, s2 string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. s1 ning permutatsiyasi degani — uzunligi s1 bilan bir xil bo'lgan va harflar chastotasi (har bir harf necha marta uchrashi) bir xil bo'lgan qism-satr degani; shuning uchun belgilangan o'lchamli oyna (fixed-size sliding window) va harf sanoqlarini solishtirish g'oyasini o'ylab ko'ring.
2. s1 dagi har bir harfning chastotasini hisoblang, so'ng s2 ustida uzunligi len(s1) bo'lgan oynani bir pozitsiyaga siljitib boring — har safar oynaga kirgan harfning sonini oshirib, chiqib ketgan harfning sonini kamaytiring; agar oynadagi chastotalar to'plami s1 nikiga aynan teng bo'lsa, true qaytaring.
