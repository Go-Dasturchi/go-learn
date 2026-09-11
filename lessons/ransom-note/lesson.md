# Ransom Note

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: ransomNote = "aa", magazine = "aab"
Output: true
```

## TASK

Sizga ikkita `ransomNote` va `magazine` nomli satrlar berilgan.

Agar `ransomNote` ni `magazine` da mavjud bo'lgan harflar yordamida yozib bo'lsa `true` qaytaring, aks holda `false`.

`magazine` dagi har bir harfdan `ransomNote` da faqat bir marta foydalanish mumkin. Harflar faqat ingliz alifbosining kichik harflaridan iborat.

Quyidagi funksiyani to'ldiring:

```go
func solve(ransomNote string, magazine string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Harflar faqat kichik lotin alifbosidan iborat bo'lgani uchun har bir harfning sonini 26 ta katakli massivda saqlashni o'ylang.
2. Avval `magazine` dagi har bir harfning nechta borligini sanang, so'ng `ransomNote` dagi har bir harf uchun shu sonlarni kamaytirib boring — agar biror harf uchun son manfiyga tushib qolsa, demak yetarli emas va `false` qaytarish kerak.
