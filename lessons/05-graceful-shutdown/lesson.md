# 05 — Graceful Shutdown

## THEORY

Serverni to'xtatish kerak bo'lganda (masalan, yangi versiyani joylashtirish uchun), uni **darhol** o'chirib qo'yish — xavfli: agar ayni shu paytda birov so'rov yuborayotgan bo'lsa, yoki baza tranzaksiyasi tugallanmagan bo'lsa, ma'lumot **yo'qolishi** yoki so'rov **yarim bajarilgan** holatda qolishi mumkin.

**Yechim — Graceful Shutdown (silliq to'xtatish).** Server, to'xtash signalini olganda:

1. Yangi so'rovlarni qabul qilishni **to'xtatadi**.
2. Hozir bajarilayotgan so'rovlarning **tugashini kutadi** (ma'lum vaqt ichida).
3. Resurslarni (baza ulanishi, fayllar) **tartibli yopadi**.
4. Shundan keyingina, dastur to'liq to'xtaydi.

**Go'da signal ushlash — `os/signal`:**

```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

<-sigCh // Ctrl+C yoki "kill" buyrug'i kelguncha shu yerda kutadi
fmt.Println("to'xtash signali qabul qilindi, silliq yopilyapti...")
```

`os.Interrupt` — `Ctrl+C` bosilganda, `syscall.SIGTERM` esa `kill <pid>` yoki Docker/Kubernetes konteyner to'xtatilganda yuboriladigan signal.

**`http.Server`'ni silliq to'xtatish:**

```go
server := &http.Server{Addr: ":8080"}

go func() {
	server.ListenAndServe() // alohida goroutine'da ishga tushadi
}()

<-sigCh // to'xtash signalini kutish

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

server.Shutdown(ctx) // yangi so'rov qabul qilmaydi, mavjudlarini kutadi
```

`server.Shutdown(ctx)` — **eng muhim qism**: u darhol serverni o'chirmaydi, balki yangi ulanishlarni rad etib, **hozirgi so'rovlar tugashini kutadi**, `ctx`da berilgan vaqt (bu yerda 10 soniya) ichida. Agar shu vaqt ichida hammasi tugamasa, majburan to'xtatiladi.

**Nega `context.WithTimeout` kerak — "Context" darsini eslang.** U yerda context'ning **bekor qilish** (cancellation) va **muddat** (deadline) g'oyasini ko'rgan edik — aynan shu g'oya, graceful shutdown'da, "cheksiz kutmaslik" uchun ishlatiladi: server abadiy kutib qolmaydi, balki belgilangan vaqtdan keyin majburan yopiladi.

**Bir nechta resursni tartibli yopish.** Katta dasturda, server, baza ulanishi, keshdagi ulanish kabi bir nechta resurs bo'lishi mumkin — ularning har biri, **teskari tartibda** (oxirgi ochilgani birinchi yopiladi) yopilishi kerak:

```go
func hammasiniYop(ctx context.Context, yopuvchilar []func(context.Context) error) error {
	for i := len(yopuvchilar) - 1; i >= 0; i-- {
		if err := yopuvchilar[i](ctx); err != nil {
			return err
		}
	}
	return nil
}
```

Bu — "Higher-Order Functions" darsida ko'rgan g'oyaning qo'llanilishi: har bir resursning "yopish" mantig'i, `func(context.Context) error` turidagi funksiya sifatida uzatiladi, va `hammasiniYop` ularni **tartib bilan** chaqiradi.

## EXAMPLE

```go
package main

import (
	"context"
	"errors"
	"fmt"
)

func hammasiniYop(ctx context.Context, yopuvchilar []func(context.Context) error) error {
	for i := len(yopuvchilar) - 1; i >= 0; i-- {
		if err := yopuvchilar[i](ctx); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var tartib []string

	yopuvchilar := []func(context.Context) error{
		func(ctx context.Context) error { tartib = append(tartib, "server"); return nil },
		func(ctx context.Context) error { tartib = append(tartib, "baza"); return nil },
		func(ctx context.Context) error { tartib = append(tartib, "kesh"); return nil },
	}

	err := hammasiniYop(context.Background(), yopuvchilar)
	fmt.Println(err, tartib)

	// xatolik bo'lsa, darhol to'xtaydi
	xatoYopuvchilar := []func(context.Context) error{
		func(ctx context.Context) error { return nil },
		func(ctx context.Context) error { return errors.New("yopib bo'lmadi") },
	}
	err = hammasiniYop(context.Background(), xatoYopuvchilar)
	fmt.Println(err)
}
```

Natija:

```
<nil> [kesh baza server]
yopib bo'lmadi
```

Resurslar **teskari tartibda** yopildi (`kesh` → `baza` → `server`), chunki ro'yxatda oxirgi qo'shilgan — birinchi ochilgan resurs deb hisoblanadi (masalan, avval server ochiladi, keyin baza, keyin kesh — yopishda esa aksincha).

## TASK

`hammasiniYop(ctx context.Context, yopuvchilar []func(context.Context) error) error` funksiyasini to'ldiring:

1. `yopuvchilar` ro'yxatini **teskari tartibda** (oxiridan boshigacha) aylanib chiqing.
2. Har birini `ctx` bilan chaqiring; agar xatolik qaytsa, **darhol** shu xatolikni qaytaring (qolganlarini chaqirmang).
3. Hammasi muvaffaqiyatli bo'lsa, `nil` qaytaring.

`main()` funksiyasini o'zgartirish shart emas.

Vim editorini ochish uchun ENTER bosing. Kod tayyor bo'lgach, **`:wqa`** (yoki `:xa`) buyrug'i bilan chiqing — ikkita oyna (dars matni + kod) ochiq bo'lgani uchun oddiy `:q` faqat bittasini yopadi va kodni saqlamaydi.

## HINTS

1. `for i := len(yopuvchilar) - 1; i >= 0; i-- { ... }` — teskari tartibda aylanish uchun.
2. `if err := yopuvchilar[i](ctx); err != nil { return err }`, tsikldan keyin `return nil`.
