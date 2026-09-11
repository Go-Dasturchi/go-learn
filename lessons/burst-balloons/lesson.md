# Burst Balloons

## THEORY

Bu — amaliy mashq masalasi (LeetCode to'plamidan, Hard daraja). Yangi Go tushunchasi emas, balki oldin o'rgangan bilimlaringizni mustahkamlash uchun mo'ljallangan.

Qiyinlik: **Qiyin** (Hard) · Ball: **30** · Taxminiy vaqt: **45 daqiqa**

## EXAMPLE

```
Input: nums = [3,1,5,8]
Output: 167
Tushuntirish:
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
```

## TASK

Sizga `n` ta sharning ustiga yozilgan raqamlar `nums` massivi berilgan. Siz barcha sharlarni portlatishingiz kerak.

Agar siz `i` indeksdagi sharni portlatsangiz, u holda `nums[i - 1] * nums[i] * nums[i + 1]` miqdorda tanga olasiz. Agarda `i - 1` yoki `i + 1` massiv chegarasidan tashqariga chiqib ketsa, ularning o'rniga `1` raqami bor deb hisoblashingiz mumkin.

Barcha sharlarni portlatib yig'ish mumkin bo'lgan maksimal tangalar (coins) miqdorini qaytaring.

Quyidagi funksiyani to'ldiring:

```go
func solve(nums []int) int {
    // ...
}
```

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. Qaysi balonni birinchi portlatish emas, balki qaysi balonni bir oraliqda ENG OXIRIDA portlatishni tanlash haqida o'ylang — bu oraliq (interval) dinamik dasturlash masalasi. Massiv chetlariga qiymati 1 bo'lgan virtual balonlar qo'shib olish chegaraviy holatlarni soddalashtiradi.
2. dp[left][right] ni faqat (left,right) ochiq oralig'idagi barcha balonlarni portlatgandan keyin olinadigan maksimal tanga miqdori deb belgilang. Har bir k ni shu oraliqda eng oxirida portlaydigan balon deb tanlab, coins = balloons[left]*balloons[k]*balloons[right] + dp[left][k] + dp[k][right] formulasini barcha mumkin bo'lgan k lar bo'yicha hisoblab, eng kattasini tanlang; oraliqlarni kichik uzunlikdan kattaga qarab to'ldiring.
