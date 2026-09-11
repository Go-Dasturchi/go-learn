# Valid Perfect Square

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: num = 16
Output: true
Tushuntirish: 4 * 4 = 16 bo'lgani uchun 16 mukammal kvadrat.
```

## TASK

Sizga musbat butun son `num` berilgan. Agar `num` mukammal kvadrat bo'lsa `true` qaytaring, aks holda `false`.

Mukammal kvadrat bu biror bir butun sonning o'ziga ko'paytmasiga teng bo'lgan sondir. O'rnatilgan tayyor kutubxonalardan, masalan, `sqrt` dan foydalanmang.

Quyidagi funksiyani to'ldiring:

```go
func solve(num int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
