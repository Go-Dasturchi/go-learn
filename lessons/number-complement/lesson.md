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

1. Sonning barcha bitlarini teskari qilish uchun sonning necha xonali (bitli) ekanini bilishga to'g'ri keladi — shu uzunlikda "hammasi 1" bo'lgan maska yaratishni o'ylang.
2. `num` dan katta yoki teng bo'lguncha maskani chapga surib 1 bilan to'ldirib boring (masalan `1, 11, 111, ...` ikkilikda), so'ng `num` ni shu maska bilan XOR qiling — natija to'ldiruvchi son bo'ladi.
