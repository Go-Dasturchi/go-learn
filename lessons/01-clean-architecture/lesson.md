# 01 — Clean Architecture

## THEORY

Kichik dasturda, hamma narsani bitta `main.go` fayliga yozib qo'yish mumkin — HTTP handler ham, SQL so'rov ham, biznes-mantiq ham bitta joyda. Lekin dastur kattalashgani sari, bu **muammoga aylanadi**: agar bazani PostgreSQL'dan boshqasiga almashtirish kerak bo'lsa, yoki biznes-mantiqni HTTP'siz sinash (test) kerak bo'lsa, hammasi bir-biriga **chirmashib ketgani** uchun, kichik o'zgarish ham butun dasturni "sindirib qo'yishi" mumkin.

**Clean Architecture** — dasturni **qatlamlarga** (layers) ajratish g'oyasi, har bir qatlam faqat **o'zidan ichkarida** joylashgan qatlamga bog'liq bo'ladi:

```
┌─────────────────────────────────────┐
│  Infrastructure (HTTP, SQL, Redis)  │  ← tashqi dunyo bilan aloqa
│  ┌─────────────────────────────┐    │
│  │  Interface (interfeyslar)    │    │  ← "nima qila oladi"
│  │  ┌───────────────────────┐  │    │
│  │  │  Use Case (biznes-     │  │    │  ← "nima qilish kerak"
│  │  │  mantiq)               │  │    │
│  │  │  ┌─────────────────┐  │  │    │
│  │  │  │  Entity (domain) │  │  │    │  ← eng markaziy, hech
│  │  │  └─────────────────┘  │  │    │    narsaga bog'liq emas
│  │  └───────────────────────┘  │    │
│  └─────────────────────────────┘    │
└─────────────────────────────────────┘
```

**Dependency Rule (bog'liqlik qoidasi) — eng muhim qoida:** o'q **faqat ichkariga** qarab yo'nalishi kerak. Markazdagi biznes-mantiq (Use Case), tashqaridagi HTTP yoki SQL haqida **hech narsa bilmasligi** kerak — aksincha, tashqi qatlamlar, markazdagi qatlamlarga (interfeyslar orqali) bog'lanadi.

**Bu — "Repository Pattern" darsining davomi.** O'sha darsda ko'rganingiz `FoydalanuvchiRepository` interfeysi — aynan Clean Architecture'dagi "Interface" qatlamining bir misoli edi. Use Case (biznes-mantiq) qatlami, `PostgresRepository`ning o'zi bilan emas, balki **interfeys** bilan ishlaydi:

```go
// Entity — eng markaziy, hech narsaga bog'liq emas
type Foydalanuvchi struct {
	ID  int
	Ism string
}

// Interface — Use Case qaysi metodlarga muhtojligini belgilaydi
type FoydalanuvchiOmbori interface {
	Ol(id int) (Foydalanuvchi, error)
}

// Use Case — biznes-mantiq, faqat interfeys bilan ishlaydi
type FoydalanuvchiUseCase struct {
	ombor FoydalanuvchiOmbori
}

func (u *FoydalanuvchiUseCase) IsmniOl(id int) (string, error) {
	f, err := u.ombor.Ol(id)
	if err != nil {
		return "", err
	}
	return f.Ism, nil
}
```

`FoydalanuvchiUseCase` — `sql.DB`, `http.Request` yoki boshqa tashqi narsalar haqida **umuman bilmaydi**. U faqat `FoydalanuvchiOmbori` interfeysini biladi — bu esa, uni **haqiqiy bazasiz**, oddiy soxta (fake) implementatsiya bilan ham sinash imkonini beradi.

**Nega bu foydali:**

1. **Biznes-mantiq — bazasiz, HTTP'siz test qilinadi.** Eng qimmatli kod (pul o'tkazish, buyurtma hisoblash kabi qoidalar) — eng tez va ishonchli sinaladi.
2. **Almashtirish oson.** PostgreSQL'ni MongoDB'ga, yoki HTTP'ni gRPC'ga almashtirish — faqat tashqi qatlamni o'zgartirishni talab qiladi, markazdagi biznes-mantiqqa tegmasdan.
3. **Katta jamoada ishlash qulay.** Har bir dasturchi, boshqalarning kodini butunlay bilmasdan, o'z qatlamida ishlashi mumkin — chunki qatlamlar orasidagi "shartnoma" interfeyslar orqali aniq belgilangan.

## EXAMPLE

```go
package main

import (
	"errors"
	"fmt"
)

type Foydalanuvchi struct {
	ID  int
	Ism string
}

type FoydalanuvchiOmbori interface {
	Ol(id int) (Foydalanuvchi, error)
}

type XotiraOmbori struct {
	malumotlar map[int]Foydalanuvchi
}

func (o *XotiraOmbori) Ol(id int) (Foydalanuvchi, error) {
	f, bor := o.malumotlar[id]
	if !bor {
		return Foydalanuvchi{}, errors.New("topilmadi")
	}
	return f, nil
}

type FoydalanuvchiUseCase struct {
	ombor FoydalanuvchiOmbori
}

func (u *FoydalanuvchiUseCase) IsmniOl(id int) (string, error) {
	f, err := u.ombor.Ol(id)
	if err != nil {
		return "", err
	}
	return f.Ism, nil
}

func main() {
	ombor := &XotiraOmbori{malumotlar: map[int]Foydalanuvchi{1: {ID: 1, Ism: "Ali"}}}
	useCase := &FoydalanuvchiUseCase{ombor: ombor}

	ism, err := useCase.IsmniOl(1)
	fmt.Println(ism, err)

	_, err = useCase.IsmniOl(99)
	fmt.Println(err)
}
```

Natija:

```
Ali <nil>
topilmadi
```

`FoydalanuvchiUseCase` — `XotiraOmbori` degan **aniq turni bilmaydi ham**, faqat `FoydalanuvchiOmbori` interfeysini biladi — ertaga u `PostgresOmbori` bilan ham, hech qanday o'zgarishsiz ishlayveradi.

## TASK

`FoydalanuvchiUseCase.IsmniOl(id int) (string, error)` methodini to'ldiring:

1. `u.ombor.Ol(id)` ni chaqiring.
2. Agar xatolik bo'lsa, `"", err` qaytaring.
3. Aks holda, topilgan foydalanuvchining `Ism` maydonini, `nil` xatolik bilan qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `f, err := u.ombor.Ol(id)` — natijada ikkita qiymat qaytadi.
2. `if err != nil { return "", err }`, aks holda `return f.Ism, nil`.
