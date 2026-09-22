// Interestingly, Go interface in C++ is more similar to
// concept + template generic code, rather than
// virtual base class + override successor

#include <concepts>
#include <format>
#include <print>
#include <string>

template<typename Animal>
concept animal = requires(Animal a) {
    { a.Speak() } -> std::same_as<std::string>;
};

std::string SaySomething(animal auto a) {
    return a.Speak();
}

struct Cat {
    std::string Name;
    std::string Voice;

    std::string Speak() const {
        return std::format("The cat {} meows {}", Name, Voice);
    }
};

struct Dog {
    std::string Name;
    std::string Voice;
    
    std::string Speak() const {
        return std::format("The dog {} barks {}", Name, Voice);
    }
};

int main() {
    Cat cat{"Fluffy", "MEOWWWW!"};
    Dog dog{"Zoe",    "GRWWR!"};
    std::println("{} {}", SaySomething(cat), SaySomething(dog));
    return 0;
}