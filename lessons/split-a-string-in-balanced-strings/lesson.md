# Split a String in Balanced Strings

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **10 daqiqa**

## EXAMPLE

```
Input: s = "RLRRLLRLRL"
Output: 4
Tushuntirish: s ni quyidagi qismlarga bo'lish mumkin: "RL", "RRLL", "RL", "RL". Ularning har birida L va R lar soni teng.
```

## TASK

Muvozanatlashgan satr (balanced string) shunday satrki, unda `'L'` va `'R'` belgilarining soni bir xil.

Sizga muvozanatlashgan `s` satri berilgan. Uni shunday qilib qism-satrlarga bo'lingki:

- Har bir qism-satr o'z-o'zidan muvozanatlashgan bo'lsin.
- Qism-satrlar soni imkon qadar **maksimal** bo'lsin.

Olinishi mumkin bo'lgan maksimal muvozanatlashgan qism-satrlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(s string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har safar `L` va `R` sonlari tenglashganda, ochko'zlik (greedy) bilan darhol bitta muvozanatli bo'lakni "kesib olish" mumkinligini o'ylang.
2. Satrni chapdan o'ngga aylanib, `L` uchun balansni oshiring, `R` uchun kamaytiring; balans nolga tushgan har safar hisoblagichni oshiring — bu eng maksimal bo'linish soni bo'ladi.
