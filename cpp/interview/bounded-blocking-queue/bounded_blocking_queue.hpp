#pragma once

#include <chrono>
#include <condition_variable>
#include <mutex>
#include <queue>
#include <stdexcept>
#include <stdexcept>

template<typename T>
class BoundedBlockingQueue {
public:
    explicit BoundedBlockingQueue(size_t capacity)
        : capacity_(capacity) {}

    // Добавить элемент; если очередь полна, ждать (или таймаут)
    bool push(const T& item, std::chrono::milliseconds timeout = std::chrono::milliseconds(0)) {
        std::unique_lock<std::mutex> lock(mutex_);
        if (timeout.count() > 0) {
            if (!not_full_.wait_for(lock, timeout, [this] { return queue_.size() < capacity_; })) {
                return false; // таймаут
            }
        } else {
            not_full_.wait(lock, [this] { return queue_.size() < capacity_; });
        }

        queue_.push(item);
        not_empty_.notify_one();
        return true;
    }

    // Извлечь элемент; если очередь пуста, ждать (или таймаут)
    bool pop(T& item, std::chrono::milliseconds timeout = std::chrono::milliseconds(0)) {
        std::unique_lock<std::mutex> lock(mutex_);
        if (timeout.count() > 0) {
            if (!not_empty_.wait_for(lock, timeout, [this] { return !queue_.empty(); })) {
                return false; // таймаут
            }
        } else {
            not_empty_.wait(lock, [this] { return !queue_.empty(); });
        }

        item = std::move(queue_.front());
        queue_.pop();
        not_full_.notify_one();
        return true;
    }

    size_t size() const {
        std::lock_guard<std::mutex> lock(mutex_);
        return queue_.size();
    }

private:
    size_t capacity_;
    std::queue<T> queue_;
    mutable std::mutex mutex_;
    std::condition_variable not_full_;
    std::condition_variable not_empty_;
};
