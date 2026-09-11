# Permutations

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **30 daqiqa**

## EXAMPLE

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

## TASK

Sizga o'zaro farq qiladigan butun sonlardan iborat `nums` massivi berilgan. Uning barcha mumkin bo'lgan o'rin almashtirishlarini (permutations) qaytaring. Javobni istalgan tartibda qaytarishingiz mumkin.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) [][]int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir pozitsiyaga qaysi son qo'yilishini tanlab, ishlatilgan sonlarni belgilab boradigan rekursiv qidiruv (backtracking) haqida o'ylab ko'ring.
2. Joriy tanlangan sonlar ketma-ketligi va qaysi indekslar allaqachon ishlatilganini bildiruvchi bayroq (used) massivini yuriting: har chaqiriqda ishlatilmagan har bir sonni tanlab, uni belgilab rekursiyaga kiring, so'ng orqaga qaytishda belgini olib tashlang; ketma-ketlik uzunligi nums bilan tenglashganda uni natijaga saqlang.
