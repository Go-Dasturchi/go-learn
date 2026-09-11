# N-Queens

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
```

## TASK

`n` soni berilgan bo'lib, `n` ta shaxmat malikasini `n x n` taxtasiga shunday joylashtiringki, ulardan hech biri bir-birini ura olmasin.
Har bir yechim taxtadagi tuzilishni o'z ichiga olishi kerak, bunda `'Q'` malikani va `'.'` bo'sh katakni bildiradi.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) [][]string {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu N-Queens II bilan bir xil backtracking (orqaga qaytish) mantig'iga asoslanadi, farqi shundaki bu safar taxta holatining o'zini ham saqlab, har bir to'liq yechim uchun taxta chizmasini yig'ish kerak bo'ladi.
2. Ustunlar va ikkala diagonal (row+col hamda row-col+n) band-emasligini belgilaydigan uchta bool massiv bilan qatorma-qator ferzi joylashtiring; har bir qatorda qaysi ustunga ferzi qo'yilganini alohida massivda saqlang. n-qator to'liq to'ldirilganda, shu saqlangan ustun raqamlaridan '.' va 'Q' belgilaridan iborat taxta satrlarini yasab natijalar ro'yxatiga qo'shing, so'ng backtrack qilib boshqa variantlarni qidiring.
