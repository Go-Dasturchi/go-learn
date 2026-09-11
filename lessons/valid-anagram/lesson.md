# Valid anagram

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "anagram", t = "nagaram"
Output: true
```

## TASK

Ikkita satr `s` va `t` berilgan. Agar `t` satr `s` satrning anagrammasi bo'lsa `true`, aks holda `false` qaytaring.

Anagram deb bironta so'z yoki gapdagi harflarning o'rinlarini almashtirish orqali yasalgan so'zga aytiladi.

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
