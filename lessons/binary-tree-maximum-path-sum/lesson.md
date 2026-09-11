# Binary Tree Maximum Path Sum

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: root = [-10,9,20,null,null,15,7] (biz uni massiv sifatida [-10, 9, 20, -999, -999, 15, 7] deb qabul qilamiz)
Output: 42
Tushuntirish: Optimal yo'l: 15 -> 20 -> 7, ularning yig'indisi: 15 + 20 + 7 = 42.
```

## TASK

Ikkilik daraxtda "yo'l" (path) deb — har qanday tugundan boshlab, ulangan tugunlar bo'ylab o'tadigan ketma-ketlikka aytiladi. Yo'l kamida bitta tugundan iborat bo'lishi va ildiz (root) orqali o'tishi shart emas.

Daraxtdagi har qanday yo'lning maksimal yig'indisini qaytaring.

_Eslatma: Masala asl nusxasida TreeNode ishlatiladi. Test tizimimizda massiv yordamida (level-order traversal, `null` = `-1` yoki qandaydir belgi bilan emas, maxsus format) ifodalanadigan darxt uchun yechish qiyin bo'lgani sababli biz uni soddalashtirilgan massivdagi har xil tugun qiymatlar va yo'llar sifatida baholaymiz. Quyidagi testda sizga ro'yxat beriladi va maksimal subarray / daraxt yig'indisi test qilinadi._

Quyidagi funksiyani to'ldiring:

```go
func solve(nodes []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir tugun uchun ikkita tushunchani ajrating: "shu tugundan faqat bitta shoxga pastga tushuvchi eng yaxshi yo'l" va "shu tugun orqali o'tuvchi, ikkala shoxni ham birlashtiruvchi eng yaxshi to'liq yo'l". Rekursiya faqat birinchisini qaytarishi, ikkinchisi esa global javobni yangilash uchun ishlatilishi kerak.
2. DFS funksiyasi har bir tugunda "shu tugundan pastga tushadigan eng katta yig'indi"ni qaytarsin (agar chap yoki o'ng gain manfiy chiqsa, uni 0 deb hisoblang — manfiy shoxni yo'lga qo'shmang). Har safar node.Val+leftGain+rightGain qiymatini global eng yaxshi natija bilan solishtirib yangilang, lekin funksiyaning o'zi faqat node.Val+max(leftGain, rightGain) ni qaytarishi kerak.
