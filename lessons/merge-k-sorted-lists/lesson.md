# Merge k Sorted Lists

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Tushuntirish: Ularni bitta qilib birlashtirsak:
1->4->5
1->3->4
2->6
Barchasini birlashtirganda: 1->1->2->3->4->4->5->6 hosil bo'ladi.
```

## TASK

_Eslatma: Ushbu masalada bog'langan ro'yxatlar (linked lists) o'rniga massivlar (arrays) beriladi. Har bir massiv o'sish tartibida tartiblangan._

Sizga k ta tartiblangan butun sonli massivlar ro'yxati `lists` berilgan.

Ularning barchasini bitta tartiblangan massivga birlashtiring va uni qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(lists [][]int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
