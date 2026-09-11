# Find Minimum in Rotated Sorted Array II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **35 daqiqa**

## EXAMPLE

```
Input: nums = [2,2,2,0,1]
Output: 0
```

## TASK

Faraz qiling, dastlab o'sish tartibida joylashgan `nums` massivi (takrorlanuvchi elementlari bo'lishi mumkin) oldindan ma'lum bo'lmagan biror indeks atrofida aylantirilgan (rotated). Masalan, `[0,1,2,4,4,4,5,6,6,7]` massivi aylantirilib `[4,5,6,6,7,0,1,2,4,4]` ko'rinishiga kelishi mumkin.

Sizga shunday aylantirilgan `nums` massivi berilgan. Undagi eng kichik elementni qaytaring.

Yechimingiz iloji boricha o'rtacha holatda `O(log n)` vaqt murakkabligida bo'lishiga harakat qiling, ammo eng yomon holatda `O(n)` bo'lishi mumkinligini inobatga oling.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Odatiy binar qidiruv mantig'ini ishlating, lekin takrorlanuvchi elementlar mavjudligi sababli nums[mid] ni nums[hi] bilan solishtirganda uchinchi holat — ular teng bo'lgan holatni alohida ko'rib chiqishingiz kerak bo'ladi.
2. lo va hi ko'rsatkichlari bilan ishlang: nums[mid] > nums[hi] bo'lsa minimum mid dan o'ngda, lo=mid+1; nums[mid] < nums[hi] bo'lsa minimum mid yoki undan chapda, hi=mid. Ular teng bo'lsa qaysi tomonda ekanini aniqlab bo'lmaydi, shu sababli xavfsiz variant sifatida hi ni bittaga kamaytiring (hi--), bu minimum elementni yo'qotmaydi.
