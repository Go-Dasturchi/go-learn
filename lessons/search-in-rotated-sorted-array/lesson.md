# Search in rotated sorted array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rtacha** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

## TASK

O'sish tartibida saralangan va o'z o'qida aylanib (rotated) berilgan `nums` massivi va `target` soni berilgan. Agar `target` mavjud bo'lsa indeksini qaytaring, aks holda `-1`. `O(log n)` vaqtda ishlashi kerak.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, target int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massiv to'liq saralanmagan bo'lsa ham, mid nuqtasi har doim ikki yarimning kamida bittasini saralangan holda qoldiradi — shu saralangan yarimni aniqlab, target o'sha yarimda bor-yo'qligini tekshiring.
2. Har qadamda nums[lo] <= nums[mid] shartini tekshiring: agar chap yarim saralangan bo'lsa va target shu yarim oralig'ida bo'lsa hi=mid-1, aks holda lo=mid+1; aks holda (o'ng yarim saralangan) target o'ng yarim oralig'ida bo'lsa lo=mid+1, aks holda hi=mid-1 qiling.
