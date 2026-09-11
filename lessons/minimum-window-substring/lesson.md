# Minimum Window Substring

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Tushuntirish: Barcha 'A', 'B' va 'C' harflarini o'z ichiga olgan eng qisqa qism satr bu "BANC".
```

## TASK

Sizga ikkita `s` va `t` satrlari berilgan. `s` satrining ichidan shunday eng qisqa qism satrni topingki, u `t` satridagi barcha belgilarni (jumladan takrorlanganlarini ham) o'z ichiga olsin. Agar bunday qism satr mavjud bo'lmasa, bo'sh satr `""` ni qaytaring.

Test keyslar shunday tuzilganki, har doim faqat bitta yagona to'g'ri javob mavjud bo'ladi.

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
