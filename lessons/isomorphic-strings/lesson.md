# Isomorphic strings

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "egg", t = "add"
Output: true
```

## TASK

`s` va `t` satrlari berilgan, ular izomorf ekanligini aniqlang.

Ikki satr `s` va `t` bir biriga izomorf deyiladi agar ulardagi harflar boshqa harflar bilan almashtirilganda ikkinchi satr kelib chiqsa. Bunda bir xil belgi o'zini joyida bitta belgiga o'zgarishi kerak va aksincha, xarita o'rnatilishi zarur.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, t string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
