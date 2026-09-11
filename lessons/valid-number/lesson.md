# Valid number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "0"
Output: true
```

## TASK

`s` satri berilgan. Undagi yozuv matematikkada yaroqli bo'lgan raqamli qiymat ekanligini aniqlang.
Unda kasr, ishora (`+`, `-`) hamda eksponensial yozuvlar (`e`, `E`) qatnashishi mumkin, lekin ularning sintaksisi qat'iy qoidalarga mos kelishi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
