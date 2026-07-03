#include <atomic>
#include <chrono>
#include <condition_variable>
#include <functional>
#include <iostream>
#include <queue>
#include <thread>

class SimpleIoService {
public:
    SimpleIoService()
        : running_(true), worker_(
              &SimpleIoService::run, this) {}

    ~SimpleIoService() {
        stop();
    }

    void stop() {
        running_ = false;
        cv_.notify_all();
        if (worker_.joinable()) {
            worker_.join();
        }
    }

    void post(const std::function<void()>& func) {
        {
            std::lock_guard<std::mutex> lock(mutex_);
            tasks_.push(func);
        }
        cv_.notify_one();
    }

private:
    void run() {
        while (running_) {
            std::function<void()> task;
            {
                std::unique_lock<std::mutex> lock(mutex_);
                cv_.wait(lock, [this] {
                    std::cout << "awakening cv" << '\n';
                    return !tasks_.empty() || !running_; // <--- exit condition!
                });

                if (!running_ && tasks_.empty()) {
                    return;
                }

                task = tasks_.front();
                tasks_.pop();
            }

            task();
        }
    }

private:
    std::atomic<bool> running_;

    std::thread worker_; // exec tasks!
    std::queue<std::function<void()>> tasks_;

    std::mutex mutex_;
    std::condition_variable cv_;
};

void exampleTask(int id) {
    std::cout << "Executing task " << id << " on thread "
              << std::this_thread::get_id() << '\n';
}

int main() {
    SimpleIoService io_service;

    for (int i = 1; i <= 5; ++i) {
        io_service.post(
            [=] { exampleTask(i); }
        );
    }

    std::this_thread::sleep_for(std::chrono::seconds(1));
    io_service.stop();

    return 0;
}

