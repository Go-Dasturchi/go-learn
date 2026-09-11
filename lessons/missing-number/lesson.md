# Missing number

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Easy daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Oson** (Easy) · Ball: **10** · Taxminiy vaqt: **15 daqiqa**

## EXAMPLE

```
Input: nums = [3,0,1]
Output: 2
Tushuntirish: Massiv oralig'i [0, 3] bo'lib, ular orasida faqat 2 yo'q.
```

## TASK

`[0, n]` oralig'idagi `n` ta turli xil butun sonlardan iborat `nums` massivi berilgan. Massivda bitta son tushib qolgan, o'sha sonni toping.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `0` dan `n` gacha bo'lgan sonlar yig'indisi uchun oddiy formula borligini eslang — bu sizga qidiruv qilmasdan turib solishtirish imkonini beradi.
2. `n = len(nums)` uchun `0+1+...+n` formula yig'indisini hisoblang, so'ng massivdagi barcha sonlarni shu yig'indidan ayirib chiqing — qolgan natija tushib qolgan son bo'ladi.
