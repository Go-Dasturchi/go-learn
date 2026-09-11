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

1. Masalani boshidan (yuqori-chapdan) emas, oxiridan (pastki-o'ng burchakdan) boshlab teskari yo'nalishda yeching — bu "shu katakdan boshlab, malika turgan katakkacha omon yetib borish uchun kerakli minimal boshlang'ich HP" haqidagi dinamik dasturlash.
2. dp[i][j] ni (i,j) katakka kirishdan oldin zarur bo'lgan minimal sog'liq deb belgilang. Chegara sifatida dp[m][n-1] = dp[m-1][n] = 1 qo'ying, so'ngra orqaga qarab dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]) formulasi bilan to'ldiring; yakuniy javob dp[0][0] bo'ladi.
