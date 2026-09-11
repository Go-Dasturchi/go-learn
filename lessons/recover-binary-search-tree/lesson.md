# Recover Binary Search Tree

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [3,1,4,null,null,2] (Level-order) -> In-order: [1,3,2,4]
Biz sizga in-order beramiz: [1,3,2,4]
Output: [1,2,3,4]
```

## TASK

Ikkilik qidiruv daraxti (BST) ning faqatgina ikkita tuguni qiymatlari o'zaro almashib qolgan. Siz daraxt tuzilishini o'zgartirmasdan, ushbu ikki tugunning o'rnini almashtirishingiz kerak.

_Eslatma: Masala test tizimida oson ishlashi uchun, sizga daraxt tugunlari qiymatlarining pre-order yoki in-order massivi beriladi. Biz soddalik uchun xato joylashgan BST ni massiv ko'rinishida beramiz va siz to'g'rilangan massivni (xuddi shu ko'rinishda) qaytarishingiz kerak._

Massivda `null` o'rniga `nil` yoki `-1` emas, balki faqat tugunlar qiymatlari (Level-order traversal kabi) ketma-ketligi yoziladi. (Soddalik uchun).  
_Masala asl nusxasi massiv emas TreeNode ustida ishlaydi, ammo biz testni array ustida qilamiz: xato BST in-order arrayi beriladi, siz uni to'g'ri in-order ga (ya'ni o'sish tartibiga) keltiring._

Quyidagi funksiyani to'ldiring:

```go
func solve(inorder []int) []int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. To'g'ri BST ning in-order ketma-ketligi doim o'sish tartibida bo'ladi — shu sababli berilgan massivda qoida buzilgan joylarni (qo'shni ikki element noto'g'ri tartibda kelgan joylarni) qidiring.
2. Massiv bo'ylab yurib, nums[i] > nums[i+1] bo'lgan joylarni toping. Agar bunday buzilish faqat bitta joyda uchrasa (ikki qo'shni tugun almashgan), o'sha ikkisini almashtiring; agar ikki alohida joyda uchrasa (uzoqdagi ikki tugun almashgan), birinchi buzilishdagi birinchi elementni va ikkinchi buzilishdagi ikkinchi elementni bir-biri bilan almashtiring.
