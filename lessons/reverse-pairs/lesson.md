# Reverse Pairs

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [1,3,2,3,1]
Output: 2
Tushuntirish: Shartni qanoatlantiradigan juftliklar: (1, 4) chunki nums[1]=3 va 3 > 2 * nums[4]=2, va (3, 4) chunki nums[3]=3 va 3 > 2 * nums[4]=2.
```

## TASK

Sizga butun sonlardan iborat `nums` massivi berilgan. Qachonki `i < j` va `nums[i] > 2 * nums[j]` o'rinli bo'lsa, bu juftlik **Muhim Teskari Juftlik** (Important Reverse Pair) deb ataladi.

Massiv ichidagi jami shunday juftliklar sonini toping.

Sizning yechimingiz `O(n log n)` vaqt murakkabligida ishlashi tavsiya qilinadi (masalan, Merge Sort yordamida).

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Bu ham inversiyalarni sanashga o'xshaydi, faqat shart nums[i] > nums[j] emas, balki nums[i] > 2*nums[j] — shu sababli modifikatsiyalangan merge sort (birlashtirib saralash) qo'llang, lekin sanashni oddiy merge bosqichidan alohida bajaring.
2. Har bir merge chaqiruvida avval ikkala yarim ichida (hali birlashtirilmasdan turib, ular allaqachon saralangan holda) chap yarimning har bir elementi uchun o'ng yarimda nums[i] > 2*nums[j] shartini qanoatlantiruvchi elementlar sonini ikkita ko'rsatkich yordamida sanab chiqing (ikkala ko'rsatkich faqat oldinga siljiydi, orqaga qaytmaydi). Shundan keyingina odatiy merge (birlashtirish) bosqichini bajarib massivni saralangan holga keltiring.
