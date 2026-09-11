# Sum of Two Integers

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: a = 1, b = 2
Output: 3
```

## TASK

Sizga ikkita butun son `a` va `b` berilgan. Ularning yig'indisini `+` yoki `-` operatorlaridan foydalanmasdan qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(a int, b int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Qo'shish amalini bitli operatsiyalar orqali ifodalash mumkinligini o'ylab ko'ring — XOR "ko'chirishsiz qo'shish"ni, AND esa qayerda ko'chirish (carry) yuzaga kelishini ko'rsatadi.
2. b nolga aylanguncha davom eting: har safar carry ni `(a & b) << 1` sifatida hisoblang, a ni `a XOR b` bilan yangilang, so'ng b ni carry ga tenglashtiring — b nolga aylanganda a natijaviy yig'indini beradi.
