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

1. Bu klassik backtracking (orqaga qaytish) masalasi — ferzilarni qatorma-qator joylashtirib boring, har bir qatorda faqat bitta ferzi bo'ladi, shuning uchun faqat ustunlar va ikkala diagonal bo'yicha to'qnashuvni tekshirish kifoya.
2. Har bir ustun uchun `cols[]`, har bir "/" diagonal uchun `row+col`, har bir "\" diagonal uchun `row-col+n` (manfiy bo'lmasligi uchun n qo'shiladi) qiymatlaridan foydalanib band joylarni belgilovchi uchta bool massiv yuriting. Qatorni to'ldirishda mos ustun/diagonal band bo'lmasa uni band deb belgilab keyingi qatorga o'ting, qaytishda esa belgini qaytadan bo'shating (backtrack); barcha n qator to'ldirilganda hisoblagichni oshiring.
