# Jump Game II

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **40 daqiqa**

## EXAMPLE

```
Input: nums = [2,3,1,1,4]
Output: 2
Tushuntirish: Eng kam sakrash — 0-indeksdan 1-indeksga (1 qadam), keyin 1-indeksdan oxiriga (3 qadam). Jami 2 ta sakrash.
```

## TASK

Sizga n ta musbat sonlardan iborat `nums` massivi berilgan. Siz dastlab massivning boshidagi (0-indeksdagi) pozitsiyasida turibsiz.

Har bir element `nums[i]` siz u joydan maksimal sakrashingiz mumkin bo'lgan uzunlikni bildiradi (ya'ni, agar indeks `i` da bo'lsangiz, ko'pi bilan `nums[i]` qadam oldinga siljishingiz mumkin).

Siz massivning oxirgi indeksiga yetib borishingiz kerak deb kafolatlanadi. Massivning oxirgi indeksiga yetib borish uchun zarur bo'lgan eng kam sakrashlar sonini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Har bir pozitsiyadan qayerga sakrash yaxshiroq ekanini alohida hisoblashga urinmang — buni "ochko'z" (greedy) BFS-ga o'xshash bitta o'tish bilan, joriy sakrashning "qamrov chegarasi" tushunchasi orqali yeching.
2. Ikki chegarani kuzatib boring: curEnd (joriy sakrash bilan yetib borish mumkin bo'lgan eng oxirgi indeks) va farthest (shu vaqtgacha ko'rilgan barcha indekslardan yetib borish mumkin bo'lgan eng uzoq nuqta). Massiv bo'ylab yurib farthest ni yangilab boring; i == curEnd bo'lgan zahoti jumps ni oshiring va curEnd = farthest deb belgilang — bu "keyingi sakrashni boshlash vaqti keldi" degani.
