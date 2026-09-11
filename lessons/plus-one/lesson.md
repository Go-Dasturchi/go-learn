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

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
