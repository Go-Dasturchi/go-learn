# Count of Smaller Numbers After Self

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Tushuntirish:
5 dan o'ng tomonda: 2 va 1 kichik (2 ta)
2 dan o'ng tomonda: 1 kichik (1 ta)
6 dan o'ng tomonda: 1 kichik (1 ta)
1 dan o'ng tomonda: hech narsa yo'q (0 ta)
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. `counts` ro'yxatini qaytaring, bunda `counts[i]` — bu `nums[i]` dan o'ng tomonda turuvchi va `nums[i]` dan qat'iy kichik sonlar soni.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu masala inversiyalarni sanashga o'xshaydi — qiymatlarning o'zini emas, balki ularning boshlang'ich indekslarini kuzatib turadigan modifikatsiyalangan merge sort (birlashtirib saralash) algoritmini qo'llashni o'ylab ko'ring.
2. Qiymatlar o'rniga indekslar massivini rekursiv ravishda saralang. Merge (birlashtirish) bosqichida, agar o'ng yarimdan bir element chap yarimdagi elementdan oldin natijaga qo'yilsa, bu o'sha paytgacha o'ng tarafdan ko'chirilgan barcha elementlar chap elementdan kichik ekanini bildiradi — shu sonni counts[chap element indeksi] ga qo'shib boring.
