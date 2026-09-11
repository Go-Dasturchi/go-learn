# N-Queens II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: n = 4
Output: 2
Tushuntirish: 4 x 4 taxtasida ferzlarni shart bilan joylashtirishning 2 xil usuli bor.
```

## TASK

**N-Ferzi** masalasida `n x n` shaxmat taxtasiga `n` ta ferzini shunday joylashtirish kerakki, ularning biortasi boshqasini shoh (tahdid) qilmasin.

Sizga `n` berilgan. Bunday joylashtirishlar sonini toping va qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
