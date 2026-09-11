# Number Complement

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: num = 5
Output: 2
Tushuntirish: 5 ning ikkilik ko'rinishi 101.
Uni teskari qilsak 010, ya'ni 2.
```

## TASK

Sizga musbat butun son `num` berilgan. Uning to'ldiruvchi sonini (complement) toping.

To'ldiruvchi son — bu asl sonning har bir bitini teskari qilib (0 → 1, 1 → 0) hosil qilingan son. Faqat sonning ikkilik ko'rinishidagi muhim bitlari (leading zeros emas) teskari qilinadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(num int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
