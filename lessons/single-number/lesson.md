# Single number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [2,2,1]
Output: 1
```

## TASK

Bo'sh bo'lmagan butun sonlar massivi `nums` berilgan. Undagi bitta elementdan tashqari barcha elementlar ikki marta qatnashadi. Faqat bir marta qatnashgan o'sha yolg'iz sonni toping.

Siz `O(n)` vaqt murakkabligida va qo'shimcha xotira ishlatmaydigan (`O(1)`) algoritm yozishingiz kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. O(1) xotira va O(n) vaqt talab qilingani uchun map o'rniga bitli operatsiyalarni o'ylang — XOR ning o'z-o'ziga qarshi xossasidan foydalaning.
2. Massivdagi barcha sonlarni ketma-ket XOR qiling: bir xil ikkita son XOR qilinganda bir-birini yo'q qiladi (natija 0), shuning uchun oxirida faqat yolg'iz son qoladi.
