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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
