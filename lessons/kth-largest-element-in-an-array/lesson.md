# Kth Largest Element in an Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Tushuntirish: Massivni kamayish bo'yicha tartiblaganda: [6, 5, 4, 3, 2, 1]. Uning ichidan 2-kattasi bu 5.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi va bitta `k` butun soni berilgan. Massivdagi `k`-eng katta elementni qaytaring.

E'tibor bering, bu tartiblangan massivdagi har xil (o'zaro farqli) elementlar ichidan emas, balki massivdagi umumiy `k`-eng katta element hisoblanadi.

Sizning yechimingiz `O(n)` vaqt murakkabligida ishlashi tavsiya qilinadi (tayyor tartiblash - sort() funksiyasisiz).

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, k int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. To'liq saralashdan tashqari, "k-eng katta element"ni topish uchun heap (uyum) yoki quickselect kabi qisman tartiblashga asoslangan usullar borligini o'ylab ko'ring — ammo eng sodda va tez tushuniladigan yo'l ham chiqishga imkon beradi.
2. Eng sodda yondashuv sifatida massivning nusxasini kamayish tartibida to'liq saralang va (k-1)-indeksdagi elementni qaytaring; O(n) talab qilinsa, minimal hajmli max-heap yoki quickselect (tez saralashdagi bo'lish qadamiga o'xshash) orqali faqat kerakli qismni tartiblang.
