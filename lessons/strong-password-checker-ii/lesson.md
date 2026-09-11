# Strong Password Checker II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: password = "IloveLe3tcode!"
Output: true
Tushuntirish: Parol barcha xavfsizlik talablariga javob beradi.
```

## TASK

Sizga `password` satri berilgan. Agar u quyidagi barcha shartlarga mos kelsa `true` aks holda `false` qaytaring:

1. Kamida 8 ta belgidan iborat bo'lishi kerak.
2. Kamida bitta kichik ingliz harfi bo'lishi kerak.
3. Kamida bitta katta ingliz harfi bo'lishi kerak.
4. Kamida bitta raqam bo'lishi kerak.
5. Kamida bitta maxsus belgi bo'lishi kerak: `!@#$%^&*()-+`
6. Hech qanday ikkita ketma-ket belgilar bir xil bo'lmasligi kerak (masalan, `"aab"` yaroqsiz).

Quyidagi funksiyani to'ldiring:

```go
func solve(password string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir shartni alohida bayroq (flag) sifatida kuzatib, satrni faqat bir marta aylanib chiqish orqali barchasini birdan tekshirish mumkin.
2. Satr bo'ylab yurib har bir belgi turini (kichik, katta, raqam, maxsus belgi) aniqlab mos bayroqni true qiling, shu bilan birga joriy belgi oldingisi bilan bir xil emasligini tekshiring; oxirida uzunlik shartini va barcha bayroqlarni birga tekshiring.
