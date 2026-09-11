# Guess Number Higher or Lower II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: n = 10
Output: 16
Tushuntirish: 7 ni taxmin qiling. Agar 7 dan katta bo'lsa: 9 ni taxmin qiling. Agar 9 dan kichik bo'lsa: 8 ni taxmin qiling.
Eng yomon holat 7 + 9 = 16 dollar.
```

## TASK

Biz `[1, n]` oralig'idagi raqamni topish o'yinini o'ynamoqdamiz. Men bir raqam tanlayman. Agar siz `x` ni taxmin qilsangiz va bu noto'g'ri bo'lsa, siz `x` dollar to'laysiz. So'ng men kattami yoki kichikmi ekanligini aytaman.

Qaysi strategiyani qo'llasangiz ham g'alaba qozonish kafolatlangan holda sarflashingiz kerak bo'lgan **minimal miqdorni** qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu oddiy binary search emas — har bir [lo, hi] oralig' uchun "eng yomon holatda kafolatlangan minimal xarajat"ni saqlaydigan intervalli dinamik dasturlash (interval DP) kerakligini o'ylab ko'ring.
2. Uzunligi oshib boruvchi har bir [lo, hi] oralig'i uchun, oraliq ichidagi har bir x ni taxmin sifatida sinab ko'ring: xarajat x + max(chap qism DP natijasi, o'ng qism DP natijasi) bo'ladi (chunki raqib eng yomon tomonni tanlaydi), va shu x lardan eng kichik xarajat beruvchisini dp[lo][hi] sifatida saqlang.
