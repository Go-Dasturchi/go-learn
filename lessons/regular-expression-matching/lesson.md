# Regular expression matching

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "aa", p = "a*"
Output: true
Tushuntirish: '*' belgisi 'a' ni qayta takrorlash imkonini beradi.
```

## TASK

Kirish satri `s` va namuna `p` berilgan bo'lib, `.` va `*` belgilarini hisobga oluvchi regex tekshirgichni amalga oshiring:

- `.` (nuqta) har qanday bitta belgiga mos keladi.
- `*` (yulduzcha) undan oldingi belgining nol yoki undan ortiq marta uchrashiga mos keladi.
  Mos kelish satrning bir qismini emas, balki to'liq satrni qoplashi kerak.

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
