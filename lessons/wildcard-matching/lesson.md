# Wildcard matching

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "aa", p = "*"
Output: true
Tushuntirish: '*' har qanday satrga to'g'ri keladi.
```

## TASK

`s` satri va `p` namunasi (pattern) berilgan. Quyidagi maxsus belgilarni inobatga olib solishtiruvchini (matcher) amalga oshiring:

- `?` har qanday yakka belgiga mos tushadi.
- `*` har qanday belgilar ketma-ketligiga (jumladan bo'shga) mos tushadi.
  Bu ham xuddi Regex kabi butun satrni qoplashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, p string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
