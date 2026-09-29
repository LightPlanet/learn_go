#include <atomic>
#include <chrono>
#include <condition_variable>
#include <mutex>

// For examples
#include <print>
#include <thread>

// Not buffered channel
template<typename T>
class channel {
public:
    void close() noexcept {
        Close.store(true);
        HasData.notify_all();
    }

    ~channel() noexcept {
        close();
    }

    friend void operator << (channel& self, const T& v) {
        std::unique_lock lock{ self.Mutex };
        self.Value = v;
        self.Updated.store(true);
        self.HasData.notify_one();
    }

    friend bool operator >> (channel& self, T& out) {
        bool cancelled;
        std::unique_lock lock{ self.Mutex };
        self.HasData.wait(lock, [&](){
            cancelled = self.Close.load();
            return cancelled || self.Updated.load();
        });
        if (!cancelled) {
            out = self.Value;
            self.Updated.store(false);
        }
        return !cancelled;
    }

private:
    std::mutex              Mutex;
    std::condition_variable HasData;
    std::atomic<bool>       Close{ false };
    std::atomic<bool>       Updated{ false };
    T                       Value;
};

// Usage examples -----------------------------------------------------------------------

using std::chrono::seconds;
using std::this_thread::sleep_for;

void NotBufferedChannelExample() {
    channel<int> ch;
    
    std::jthread t{[&ch](){
        int value;
        while (true) {
            if (ch >> value)
                std::println("Received: {}", value);
            else {
                std::println("Canceled");
                return;
            }
        }
    }};

    sleep_for(seconds(1));
    ch << 123;
    
    sleep_for(seconds(1));
    ch.close();
}

std::mutex Mutex;

void Example() {
    // locks in constructor
    std::unique_lock lock{ Mutex };
    
    // do stuff, may throw an exception...
    
    // unlocks underlying Mutex
    lock.unlock();

    // Desctructor is always called
    // unique_lock desctructor unlocks Mutex if is locked
}

int main() {
    NotBufferedChannelExample();
}