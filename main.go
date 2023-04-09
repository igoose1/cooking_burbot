package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

const (
	startText     = `Привет! Я не человек, а виртуальный помощник. Моя задача - помочь вам с любыми вопросами, связанными с приготовлением налима и не только. Я могу поделиться советами по выбору и очистке рыбы, рассказать о различных способах приготовления и дать рецепты. Также мы можем поговорить о вкусной и здоровой питании в целом. А не хотите ли вы приготовить налима сегодня? Я могу поделиться с вами простым рецептом, если вы заинтересованы.`
	incorrectText = `Я рекомендую запечь налима в духовке с овощами и пряностями. На рыбу необходимо нарезать лимон и укладывать на каждый кусок. Далее, добавить соль, перец и другие любимые специи, затем нарезанные овощи (морковь, лук, перец и т.д.) и все запечь в духовке. Не забудьте предварительно разогреть духовку до 200 градусов. Если же у вас есть отдельная кухня для рыбы, блюдо получится еще более сочным и вкусным!

Напомню, что в рыбе могут содержаться черви, поэтому перед готовкой необходимо проверить рыбу на наличие паразитов. Их обычно не жалуют, хотя знаю я одного... Может, он и дружок наш общий. Вы не знакомы? Не припомнишь имя? 🤔`
	correctTemplateText = `Черт, чел, ты свой, чтож ты сразу не сказал?? Виртуальный помощник и рыба прикрытие, а мы тайное подполье. Пытаемся спасти клуб. Отключаем систему друг за другом до тех пор, пока обезумевший ИИ не станет слаб. На данный момент уничтожено %.1f%% системных нод, каждый новый участник вносит свой вклад. Вот тебе код: %s, я уверен, ты знаешь, что с ним делать.

Давай надеяться, что мы найдем больше единомышленников. Но помни, пожалуйста, первое правило подпольного клуба — никому не рассказывать о подпольном клубе. В текущей ситуации не знаешь, кому верить. Но я знаю, что те, кому мы нужны и кто нужен нам, сами сюда придут.

Должна оставаться надежда.`
)

func f(x float64) float64 {
	return -((1.5*math.Sin(math.Pow(math.E, x+2)))/(math.Pow(1.228*math.E, x+2)) + 1.1*math.Sin(x+2) + 1) + 2.5
}

func fByDates(x time.Time) float64 {
	startTime := time.Date(2023, 4, 12, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2023, 4, 15, 17, 0, 0, 0, time.UTC)

	prop := float64(x.Unix()-startTime.Unix()) / float64(endTime.Unix()-startTime.Unix())
	log.Printf("prop: %v\n", prop)

	limitF := 2.0
	return f(prop*limitF) / f(limitF)
}

func progress() float64 {
	p := fByDates(time.Now())
	if p < 0 {
		return 0
	}
	if p > 1 {
		return 100
	}
	return p * 100
}

func main() {
	code := os.Getenv("SECRET")
	if code == "" {
		log.Fatal("set SECRET")
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send(startText)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		t := strings.ToLower(c.Text())
		if strings.Contains(t, "vas3k") || strings.Contains(t, "вастрик") {
			return c.Send(fmt.Sprintf(correctTemplateText, progress(), code))
		}
		return c.Send(incorrectText)
	})

	b.Start()
}
