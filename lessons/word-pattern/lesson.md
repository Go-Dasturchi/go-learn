# Word Pattern

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Tushuntirish:
'a' -> "dog"
'b' -> "cat"
```

## TASK

Sizga bitta qonuniyatni bildiruvchi `pattern` satri va kichik ingliz harflaridan tuzilgan so'zlardan iborat bo'lgan `s` satri berilgan.

Agar `s` dagi so'zlar ketma-ketligi `pattern` dagi qonuniyatga (bijection) to'liq mos tushsa, `true` qaytaring, aks holda `false`.

Bijection deganda `pattern` dagi har bir belgi `s` dagi alohida va yagona bir so'zga mos tushishi va aksincha bo'lishi tushuniladi.

Quyidagi funksiyani to'ldiring:

```go
func solve(pattern string, s string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
