# Dungeon Game

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
Output: 7
Tushuntirish: Qahramon (O'ngga -> O'ngga -> Pastga -> Pastga) marshruti orqali yursa eng kam zarar ko'radi. Boshlash uchun unga kamida 7 sog'lik kerak bo'ladi.
```

## TASK

Qahramon (Knight) zindonda (dungeon) qamab qo'yilgan malikani qutqarishi kerak. Zindon $m \times n$ o'lchamli kataklardan iborat. Qahramon dastlab eng yuqori chap tomondagi $(0, 0)$ katakda bo'ladi va har safar faqat o'ngga yoki pastga harakatlanib malika joylashgan eng pastki o'ng $(m-1, n-1)$ katakka yetib borishi kerak.

Zindonning har bir katagida `dungeon[i][j]` qiymati mavjud.

- Agar qiymat manfiy bo'lsa, qahramonning sog'ligi (health) shu miqdorga kamayadi (maxluqqa duch keladi).
- Agar qiymat musbat bo'lsa, sog'ligi shu miqdorga ortadi (sehrli dori topadi).
- Agar qahramonning sog'ligi 0 yoki undan kam bo'lib qolsa, u darhol halok bo'ladi.

Qahramon malikani muvaffaqiyatli qutqarib qolishi uchun eng kamida qanday dastlabki sog'likka (initial health) ega bo'lishi kerakligini toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(dungeon [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
