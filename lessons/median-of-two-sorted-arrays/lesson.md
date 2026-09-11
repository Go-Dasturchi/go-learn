# Median of two sorted arrays

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Tushuntirish: Birlashgan massiv = [1,2,3] va uning mediani 2.
```

## TASK

O'lchamlari mos ravishda `m` va `n` bo'lgan ikkita saralangan `nums1` va `nums2` massivlari berilgan. Ularning umumiy medianini toping.

Umumiy vaqt murakkabligi `O(log(m+n))` bo'lishi talab qilinadi.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums1 []int, nums2 []int) float64 {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. To'liq birlashtirib chiqishning hojati yo'q — kichikroq massivda binar qidiruv qilib, ikkala massivni "chap yarim" va "o'ng yarim"ga bo'luvchi kesish nuqtasini toping, shunda chap yarimdagi elementlar soni umumiy uzunlikning yarmiga teng bo'lsin.
2. `nums1` da `i` ta, `nums2` da esa `j = half - i` ta elementni chap tomonga olib, `nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]` shartini tekshiring (chegara holatlar uchun cheksizliklardan foydalaning). Shart bajarilmasa, qaysi tomonga siljish kerakligini aniqlab (agar nums1[i-1] > nums2[j] bo'lsa i ni kamaytiring, aks holda oshiring) binar qidiruvni davom ettiring; shart bajarilganda chap qismning maksimumi va (juft uzunlikda) o'ng qismning minimumidan medianani hisoblang.
