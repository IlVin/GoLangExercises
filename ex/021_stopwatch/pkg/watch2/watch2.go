package watch2

import "time"

type Stopwatch struct {
	start     time.Time
	durations []time.Duration
}

// Сбрасывает таймер и начинает отсчет времени сначала
func (w *Stopwatch) Start() {
	w.start = time.Now()
	w.durations = []time.Duration{}
}

// Запоминает промежуточное время
func (w *Stopwatch) SaveSplit() {
	w.durations = append(w.durations, time.Now().Sub(w.start))
}

// Возвращает слайс с продолжительностями промежуточных результатов относительно стартового времени
func (w *Stopwatch) GetResults() []time.Duration {
	return w.durations
}
