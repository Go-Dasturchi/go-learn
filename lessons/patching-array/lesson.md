# Patching Array

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [1,3], n = 6
Output: 1
Tushuntirish: nums = [1,3] bilan quyidagilar ifodalanishi mumkin: 1, 3, 4.
2 ni qo'shsak: [1,2,3] => 1, 2, 3, 4, 5, 6 barchasini ifodalash mumkin.
```

## TASK

Sizga tartiblangan `nums` butun sonlar massivi va butun son `n` berilgan. `nums` massiviga ba'zi sonlar qo'shib (`patch`) `[1, n]` oralig'idagi har bir butun sonni massiv elementlarining quyi to'plami sifatida ifodalash imkoniga ega bo'lmoqchisiz.

Siz qo'shishingiz kerak bo'lgan **minimal** sonlar miqdorini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int, n int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Ochko'z (greedy) yondashuv qo'llang: "hozircha [1, miss-1] oralig'idagi barcha sonlarni ifodalay olamiz" degan invariantni saqlab, miss qiymatini bosqichma-bosqich qanday o'sishini kuzatib boring.
2. `miss` — hozircha kafolatlangan diapazondan tashqaridagi eng kichik ifodalab bo'lmaydigan son. Agar nums dagi keyingi son `miss` dan kichik yoki teng bo'lsa, uni diapazonga qo'shib miss += nums[i] qiling (patch shart emas). Aks holda diapazonni ikki baravar kengaytiruvchi `miss` ning o'zini "sun'iy" qo'shib (miss += miss) patches sonini oshiring — bu ochko'zlik miss <= n bo'lguncha davom etadi.
