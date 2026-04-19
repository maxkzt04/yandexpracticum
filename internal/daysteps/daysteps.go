package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	// Разделяем строку на части по запятой
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	// Преобразуем количество шагов в целое число
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, err
	}

	// Проверяем, что количество шагов больше нуля
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	// Преобразуем продолжительность в тип time.Duration
	durStr := parts[1]
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, 0, err
	}

	// Проверяем, что продолжительность больше нуля
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	// Возвращаем количество шагов и продолжительность
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	// Парсим входные данные
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Проверяем, что количество шагов и продолжительность больше нуля
	if steps <= 0 {
		return ""
	}

	// Считаем дистанцию для дневной активности по фиксированной длине шага.
	var distanceM float64
	distanceM = float64(steps) * stepLength
	// Конвертируем дистанцию в километры
	distanceKm := distanceM / float64(mInKm)

	// Считаем количество сожжённых калорий для дневной активности.
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Формируем строку с информацией о дневной активности
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
	return result
}
