package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	// Разделяем строку на части по запятой
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных тренировки")
	}

	// Преобразуем количество шагов в целое число
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, err
	}

	// Проверяем, что количество шагов больше нуля
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	// Получаем тип активности
	activityType := parts[1]

	// Преобразуем продолжительность в тип time.Duration
	durStr := parts[2]
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	// Возвращаем количество шагов, тип активности и продолжительность
	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	// Длина шага считается от роста пользователя.
	stepLength := height * stepLengthCoefficient
	if stepLength <= 0 {
		// Если рост некорректный, используем среднюю длину шага из константы.
		stepLength = lenStep
	}

	// Считаем дистанцию в метрах
	distanceM := float64(steps) * stepLength

	// Возвращаем дистанцию в километрах
	km := distanceM / float64(mInKm)
	return km
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	// Проверяем, что количество шагов и продолжительность больше нуля
	if duration <= 0 {
		return 0
	}

	// Считаем среднюю скорость в км/ч
	distanceKm := distance(steps, height)
	hours := duration.Hours()
	speed := distanceKm / hours

	// Скорость считается как дистанция, деленная на время в часах
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	// Парсим входные данные
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Проверяем, что количество шагов и продолжительность больше нуля
	distanceKm := distance(steps, height)

	// Считаем среднюю скорость
	speed := meanSpeed(steps, height, duration)

	var calories float64

	// Считаем калории в зависимости от типа активности
	switch activityType {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	// Формируем строку с информацией о тренировке
	hours := duration.Hours()
	text := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, hours, distanceKm, speed, calories)
	return text, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверяем, что количество шагов и продолжительность больше нуля
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	// Проверяем, что вес, рост и продолжительность больше нуля
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}

	// Проверяем, что рост больше нуля
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}

	// Проверяем, что продолжительность больше нуля
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	// Считаем калории по формуле: (вес * скорость * время в минутах) / 60
	speed := meanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Считаем калории (для бега коэффициент как в методичке — без множителя ходьбы)
	calories := (weight * speed * durationInMinutes) / float64(minInH)

	// Возвращаем результат расчёта для бега
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	// Проверяем, что количество шагов и продолжительность больше нуля
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}

	// Проверяем, что вес, рост и продолжительность больше нуля
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше нуля")
	}

	// Проверяем, что рост больше нуля
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}

	// Проверяем, что продолжительность больше нуля
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше нуля")
	}

	// Считаем калории по формуле: (вес * скорость * время в минутах) / 60, учитывая коэффициент для ходьбы
	speed := meanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	durationInMinutes := duration.Minutes()

	// Считаем базовое количество калорий
	baseCalories := (weight * speed * durationInMinutes) / float64(minInH)

	// Учитываем коэффициент для ходьбы
	result := baseCalories * walkingCaloriesCoefficient
	return result, nil
}
