# Maximum Number of Words Found in Sentences

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **5 daqiqa**

## EXAMPLE

```
Input: sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
Output: 6
Tushuntirish:
1-gap: 5 so'z
2-gap: 4 so'z
3-gap: 6 so'z. Maksimal so'zlar soni 6 ga teng.
```

## TASK

Sizga satrlar massivi `sentences` berilgan. Har bir element so'zlari orasida bo'sh joy bo'lgan bitta gapni bildiradi.

Jamlangan so'zlar bo'yicha eng ko'p so'z qatnashgan gapni topib, uning ichidagi so'zlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(sentences []string) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir gapdagi so'zlar sonini bo'sh joylar bo'yicha ajratib sanashni o'ylang.
2. Har bir gapni bo'sh joy bo'yicha bo'laklarga ajrating (masalan, `strings.Fields`), hosil bo'lgan bo'laklar sonini hisoblang va barcha gaplar ichidan eng kattasini saqlab boring.
