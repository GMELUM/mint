package mysql

import (
	"errors"
	"sync"
	"time"
)

// Структура, которая будет хранить мьютекс и время последней блокировки
type MutexData struct {
	mu       *sync.RWMutex
	lastLock time.Time
	timer    *time.Timer
}

// Основная структура, которая управляет мьютексами
type LocalMutex struct {
	mx   map[string]*MutexData
	lock sync.Mutex // для синхронизации доступа к карте
}

// Конструктор для LocalMutex
func NewLocalMutex() *LocalMutex {
	return &LocalMutex{
		mx: make(map[string]*MutexData),
	}
}

// Получение или создание мьютекса для конкретного ключа
func (m *LocalMutex) getMutexForKey(key string) *MutexData {
	m.lock.Lock()
	defer m.lock.Unlock()

	// Если мьютекса для ключа ещё нет, создаём новый
	if _, exists := m.mx[key]; !exists {
		m.mx[key] = &MutexData{
			mu:       &sync.RWMutex{},
			lastLock: time.Now(),
		}
	}
	return m.mx[key]
}

// Запуск таймера для удаления мьютекса через 10 секунд
func (m *LocalMutex) startTimeoutTimer(key string) {
	// Запускаем новый таймер для ключа
	m.lock.Lock()
	defer m.lock.Unlock()

	// Если таймер уже существует, остановим его
	if data, exists := m.mx[key]; exists && data.timer != nil {
		data.timer.Stop()
	}

	// Создаём новый таймер, который удаляет мьютекс через 10 секунд
	m.mx[key].timer = time.AfterFunc(10*time.Second, func() {
		m.lock.Lock()
		defer m.lock.Unlock()

		// Если блокировка не была снята, удаляем мьютекс
		if data, exists := m.mx[key]; exists && time.Since(data.lastLock) >= 10*time.Second {
			delete(m.mx, key)
			println("Мьютекс для ключа", key, "удален из-за тайм-аута.")
		}
	})
}

// Блокировка для конкретного ключа
func (m *LocalMutex) Lock(key string) error {
	data := m.getMutexForKey(key)
	data.mu.Lock()

	// Обновляем время последней блокировки и запускаем таймер
	data.lastLock = time.Now()
	m.startTimeoutTimer(key)

	return nil
}

// Разблокировка для конкретного ключа
func (m *LocalMutex) Unlock(key string) error {
	data := m.getMutexForKey(key)
	data.mu.Unlock()

	// Останавливаем таймер и удаляем мьютекс
	m.lock.Lock()
	defer m.lock.Unlock()

	if data.timer != nil {
		data.timer.Stop()
		delete(m.mx, key)
	}
	return nil
}

// Блокировка для чтения для конкретного ключа
func (m *LocalMutex) RLock(key string) error {
	data := m.getMutexForKey(key)
	data.mu.RLock()

	// Обновляем время последней блокировки и запускаем таймер
	data.lastLock = time.Now()
	m.startTimeoutTimer(key)

	return nil
}

// Разблокировка для чтения для конкретного ключа
func (m *LocalMutex) RUnlock(key string) error {
	data := m.getMutexForKey(key)
	data.mu.RUnlock()

	// Останавливаем таймер и удаляем мьютекс
	m.lock.Lock()
	defer m.lock.Unlock()

	if data.timer != nil {
		data.timer.Stop()
		delete(m.mx, key)
		println("Мьютекс для ключа", key, "разблокирован для чтения и удален.")
	}
	return nil
}

// Удаление мьютекса для ключа (по необходимости)
func (m *LocalMutex) DeleteKey(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, exists := m.mx[key]; !exists {
		return errors.New("mutex for key does not exist")
	}

	delete(m.mx, key)
	return nil
}
