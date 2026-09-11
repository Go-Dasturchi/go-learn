# Reverse Vowels of a String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: s = "hello"
Output: "holle"
Tushuntirish: unli harflar e va o. Ularning o'rnini almashtirsak "holle" hosil bo'ladi.
```

## TASK

Sizga satr ko'rinishidagi `s` berilgan. Faqatgina undagi unli harflarning (vowels) o'rnini teskari tartibda o'zgartiring va satrni qaytaring.

Unli harflar `'a'`, `'e'`, `'i'`, `'o'` va `'u'` bo'lib, ular satrda katta yoki kichik harflarda uchrashi mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
