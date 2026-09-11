# Happy Number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: n = 19
Output: true
Tushuntirish:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1 (baxtli!)
```

## TASK

**Baxtli son (Happy Number)** quyidagi jarayon orqali aniqlanadi:

1. Sonni uning har bir raqamining kvadratlari yig'indisi bilan almashtiring.
2. Bu jarayonni natija 1 ga teng bo'lgunga yoki sonsiz siklga kirgunga qadar takrorlang.
3. Agar jarayon 1 da tugasa — bu baxtli son.
4. Agar sonsiz siklga kirsa — bu baxtli son emas.

`n` baxtli son bo'lsa `true`, aks holda `false` qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
