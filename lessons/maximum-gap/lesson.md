# Maximum Gap

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [3,6,9,1]
Output: 3
Tushuntirish: Tartiblangandan so'ng [1,3,6,9]. Qo'shni elementlar o'rtasidagi farq: (3-1)=2, (6-3)=3, (9-6)=3. Maksimal farq 3 ga teng.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Uni tartiblagandan (sorted) so'ng qo'shni elementlar orasidagi maksimal farqni toping.
Agar massivda 2 tadan kam element bo'lsa, `0` qaytaring.

Algoritm chiziqli vaqt (O(n)) va chiziqli qo'shimcha xotira (O(n)) ishlatib ishlashi kerak (masalan, Radix Sort yoki Bucket Sort).

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Solishtirishga asoslangan saralash O(n log n) beradi, lekin bizga O(n) kerak — bu yerda Bucket Sort g'oyasidan foydalaning: agar sonlarni (max-min)/(n-1) hajmidagi "chelaklar"ga taqsimlasangiz, javob hech qachon bitta chelak ICHIDA bo'lmaydi, faqat chelaklar ORASIDA bo'ladi.
2. Har bir chelakda faqat minimum va maksimum qiymatlarni saqlash yetarli (chelak ichidagi tartib ahamiyatsiz, chunki maksimal farq ikki qo'shni to'lgan chelak orasida bo'ladi). Chelaklarni to'ldirib bo'lgach, ketma-ket to'lgan chelaklar bo'ylab yurib, oldingi chelakning maksimumi bilan joriy chelakning minimumi orasidagi farqni hisoblab, eng kattasini toping.
