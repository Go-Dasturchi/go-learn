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

1. Bu "sirg'anuvchi oyna" (sliding window) masalasi — ikkita ko'rsatkich (left, right) bilan oynani kengaytirib-toraytirib, oyna ichida `t` dagi barcha belgilar yetarli miqdorda mavjudligini kuzatib boring.
2. `t` dagi har bir belgi uchun kerakli sonini xesh-jadvalda saqlang. `right` ni oshirib oynani kengaytiring va oyna ichidagi hisoblagichni yangilang; qachonki oyna talab qilingan barcha belgilarni yetarlicha o'z ichiga olsa (masalan, "formed == required" turdagi hisoblagich orqali tekshiring), `left` ni imkon qadar o'ngga siljitib oynani toraytiring va har safar joriy oyna uzunligini eng yaxshi natija bilan solishtiring.
