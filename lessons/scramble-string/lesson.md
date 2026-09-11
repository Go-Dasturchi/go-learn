# Scramble String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: s1 = "great", s2 = "rgeat"
Output: true
Tushuntirish: "great" satrini "gr" va "eat" qismlariga ajratamiz.
Ularning o'rnini almashtirsak "eatgr" hosil bo'ladi.
"gr" ni ham "g" va "r" ga ajratib o'rnini almashtirsak "rg" hosil bo'ladi.
Jami o'zgarishlar bilan "rgeat" ni hosil qilish mumkin.
```

## TASK

Biz biron bir satrni uning ikkita bo'sh bo'lmagan qismga ajratib, ularning o'rinlarini almashtirish orqali chalkashtirib (scramble qilib) yuborishimiz mumkin. Ushbu jarayonni har bir qism satrda rekursiv ravishda davom ettirish mumkin.

Sizga bir xil uzunlikdagi ikkita `s1` va `s2` satrlari berilgan. Agar `s2` satri `s1` ning chalkashtirilgan versiyasi bo'lsa, `true` qaytaring, aks holda `false`.

Quyidagi funksiyani to'ldiring:

```go
func solve(s1 string, s2 string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
