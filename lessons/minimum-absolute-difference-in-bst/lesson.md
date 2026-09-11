# Minimum Absolute Difference in BST

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [1,3,6,9,12]
Output: 2
Tushuntirish: 3 - 1 = 2 (minimal farq).
```

## TASK

Sizga tartiblangan `nums` massivi berilgan (bu BST ning in-order traversal natijasidir). Ushbu massivdagi istalgan ikki element o'rtasidagi minimal mutlaq farqni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Massiv allaqachon tartiblangan bo'lgani uchun eng kichik farqni izlash uchun barcha juftliklarni emas, faqat qo'shni elementlarni tekshirish kifoya.
2. Massiv bo'ylab yurib, har bir qo'shni juftlik orasidagi farqni (`nums[i]-nums[i-1]`) hisoblang va shu farqlarning eng kichigini saqlab boring — javob shu bo'ladi.
