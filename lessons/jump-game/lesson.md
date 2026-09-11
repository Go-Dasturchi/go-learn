# Jump Game

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [2,3,1,1,4]
Output: true
Tushuntirish: 0-indeksdan 1 qadam sakrab 1-indeksga o'tamiz, u yerdan esa 3 qadam sakrab oxirgi indeksga yetib boramiz.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Siz dastlab massivning birinchi indeksida joylashgansiz va har bir element sizning o'sha joydan maksimal qancha sakrashingiz mumkinligini bildiradi.

Agar oxirgi indeksga yetib bora olsangiz `true` qaytaring, aks holda `false`.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) bool {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir sakrash yo'lini alohida sinab ko'rish shart emas — har bir pozitsiyada "shu yergacha yetib borib, keyin qay yergacha uzoqlasha olishimiz mumkin" degan greedy (ochko'z) fikrlashni qo'llang.
2. "Hozirgacha yeta oladigan eng uzoq indeks" (maxReach) degan o'zgaruvchini saqlab, massiv bo'ylab yuring: agar joriy indeks maxReach dan katta bo'lib qolsa, demak bu yerga yetib bo'lmaydi — false qaytaring; aks holda maxReach ni `i + nums[i]` bilan yangilab boring va oxirigacha yetib borsangiz true qaytaring.
