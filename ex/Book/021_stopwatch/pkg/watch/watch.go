package watch

import "time"

type Stopwatch struct {
	times []time.Time
}

// Сбрасывает таймер и начинает отсчет времени сначала
func (w *Stopwatch) Start() {
	w.times = []time.Time{time.Now()}
}

// Запоминает промежуточное время
func (w *Stopwatch) SaveSplit() {
	w.times = append(w.times, time.Now())
}

// Возвращает слайс с продолжительностями промежуточных результатов относительно стартового времени
func (w *Stopwatch) GetResults() []time.Duration {
	result := []time.Duration{}
	for i := 1; i < len(w.times); i++ {
		result = append(result, w.times[i].Sub(w.times[0]))
	}
	return result
}
