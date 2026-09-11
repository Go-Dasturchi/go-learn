# Is Subsequence

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "abc", t = "ahbgdc"
Output: true
```

## TASK

Sizga ikkita `s` va `t` satrlar berilgan. Agar `s` satri `t` satrining qism-satri bo'lsa `true` qaytaring, aks holda `false`.

Satrning qism-satri (subsequence) bu boshqa satrdan ba'zi belgilarni o'chirish (yoki o'chirmaslik) orqali hosil bo'ladigan yangi satrdir, bunda qolgan belgilarni o'zaro joylashuv tartibi o'zgarmasligi kerak (masalan, `"ace"` satri `"abcde"` ning qism-satri, lekin `"aec"` emas).

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
