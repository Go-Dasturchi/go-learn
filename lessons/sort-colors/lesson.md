# Sort Colors

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Medium daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **O'rta** (Medium) · Ball: **20** · Taxminiy vaqt: **20 daqiqa**

## EXAMPLE

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

## TASK

Sizga `n` ta elementdan iborat `nums` massivi berilgan bo'lib, ular qizil, oq yoki ko'k ranglarga bo'yalgan. Ular bir xil rangdagi elementlar yonma-yon kelishi uchun (qizil, oq va ko'k ketma-ketligida) ularni in-place (joyida) tartiblang.

Biz ranglarni ifodalash uchun mos ravishda `0`, `1` va `2` butun sonlaridan foydalanamiz.

Siz standart kutubxonadagi tayyor sort funksiyasidan foydalanmasligingiz kerak.
_Eslatma: Bu yerda bizning test tizimimiz talabi bilan funksiya massivni o'zgartiribgina qolmay, o'sha o'zgartirilgan massivni qaytarishi (return) kerak._

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Masala shartini qayta o'qib, kerakli algoritmni aniqlang.
2. Avval eng oddiy (naiv) yechimni yozing, keyin kerak bo'lsa optimallashtiring.
