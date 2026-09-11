# Russian Doll Envelopes

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
Output: 3
Tushuntirish: Konvertlarni eng ko'p [2,3] => [5,4] => [6,7] tartibida joylashtirish mumkin.
```

## TASK

Sizga `envelopes` massivi berilgan, bunda `envelopes[i] = [wi, hi]` i-chi konvertning eni va balandligini bildiradi.

Bitta konvertni boshqa bir konvert ichiga joylashtirish uchun, ichkisining eni ham, balandligi ham tashqisidan qat'iy kichik bo'lishi shart.

Bir-birining ichiga joylashtirilishi mumkin bo'lgan maksimal konvertlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(envelopes [][]int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu masalani "Longest Increasing Subsequence" (eng uzun o'suvchi qism-ketma-ketlik) masalasiga aylantirish mumkin — buning uchun konvertlarni to'g'ri tartibda saralab olish kerak.
2. Konvertlarni eni bo'yicha o'sish tartibida, eni teng bo'lganda esa balandligi bo'yicha KAMAYISH tartibida saralang (bu bir xil enli konvertlarning bir-biriga tasodifan "ichma-ich" hisoblanib ketishining oldini oladi). Shundan so'ng faqat balandliklar ketma-ketligida eng uzun QAT'IY o'suvchi qism-ketma-ketlikni O(n log n) da (tails massivi va binar qidiruv bilan) toping.
