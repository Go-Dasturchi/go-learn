# Interleaving String

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Tushuntirish: "aadbbcbcac" ni "aabcc" va "dbbca" dan uzilmagan holda terib chiqish mumkin.
```

## TASK

Sizga uchta `s1`, `s2` va `s3` satrlari berilgan. `s3` satri `s1` va `s2` satrlarining elementlarini qandaydir ketma-ketlikda qorishib yuborish (interleave) orqali hosil bo'lganligini yoki yo'qligini aniqlang.

Ikki satrning qorishib yuborilishi deganda ularni bir qancha qismlarga bo'lib, keyin shu qismlarni navbatma-navbat birlashtirib chiqish tushuniladi. Bunda har bir satr ichidagi harflarning o'zaro ketma-ketlik tartibi o'zgarmaydi.

Quyidagi funksiyani to'ldiring:

```go
func solve(s1 string, s2 string, s3 string) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
