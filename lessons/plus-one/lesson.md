# Plus one

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: digits = [1,2,3]
Output: [1, 2, 4]
Tushuntirish: Massiv 123 sonini ifodalaydi. 123 + 1 = 124. Shuning uchun [1,2,4] qaytariladi.
```

## TASK

Katta butun son raqamlari massiv (`digits`) ko'rinishida berilgan bo'lib, eng katta darajali raqam massivning boshida joylashgan. Har bir element faqat bitta raqamdan iborat (0 dan 9 gacha).

Shu katta butun songa birni qo'shing va natijani xuddi shunday massiv ko'rinishida qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(digits []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Qo'lda qo'shishda bo'lgani kabi, eng oxirgi (kichik xonali) raqamdan boshlab chapga qarab carry (o'tkazma) tarqalishini o'ylang.
2. Massivning oxiridan boshlab yuring: agar joriy raqam 9 dan kichik bo'lsa, uni birga oshirib darhol natijani qaytaring; 9 bo'lsa uni 0 ga aylantirib chapdagi raqamga o'ting — agar hamma raqam 9 bo'lib chiqsa, oldiga qo'shimcha `1` qo'shing.
