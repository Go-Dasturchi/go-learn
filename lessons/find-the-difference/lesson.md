# Find the Difference

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: s = "abcd", t = "abcde"
Output: "e"
Tushuntirish: "e" harfi qo'shilgan.
```

## TASK

Sizga `s` va `t` nomli ikkita satr berilgan.
`t` satri `s` satrining harflari aralashtirilgan va bitta tasodifiy harf qo'shilgan holatda hosil qilingan.

`t` satriga qo'shilgan o'sha bitta harfni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string, t string) string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
