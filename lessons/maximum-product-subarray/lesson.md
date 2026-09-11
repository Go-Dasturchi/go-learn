# Maximum Product Subarray

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **25 daqiqa**

## EXAMPLE

```
Input: nums = [2,3,-2,4]
Output: 6
Tushuntirish: [2,3] qism-massivi eng katta ko'paytmani (6) beradi.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Eng katta ko'paytmaga ega bo'lgan uzluksiz qism-massivni (subarray) toping va uning ko'paytmasini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Manfiy sonlar borligi sababli faqat "eng katta joriy ko'paytma"ni kuzatish yetarli emas — manfiy son eng kichik (eng manfiy) ko'paytmani ham eng kattaga aylantirib yuborishi mumkinligini o'ylab ko'ring.
2. Har bir qadamda ham joriy maksimal, ham joriy minimal ko'paytmani saqlab boring; joriy son manfiy bo'lsa ikkalasini almashtirib qo'ying, so'ng har birini (joriy son bilan davom ettirish yoki joriy sondan qaytadan boshlash) orasidan kattasini/kichigini tanlab yangilang va shu jarayondagi eng katta maksimal qiymatni umumiy javob sifatida saqlang.
