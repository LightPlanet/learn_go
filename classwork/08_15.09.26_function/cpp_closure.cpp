#include <memory>
#include <print>
#include <string>

struct defer {
    std::string& S;
    defer(std::string& s) : S{s} {}
    ~defer() { S = "На самом деле"; } // change on destruction
};

// Go: local variable is copied (returns: "Казалось бы")
// C++: undefined behaviour (without NRVO: "Казалось бы", with NRVO: "На самом деле")
std::string f1() {
    std::string value{ "Казалось бы" };
    defer d{value};
    return value;
}

// Go-like named return behaviour (returns: "На самом деле")
// Reference counter is used to emulate automatic memory management
std::shared_ptr<std::string> f2() {
    auto value = std::make_shared<std::string>("Казалось бы");
    defer d{*value};
    return value;
}

int main() {
    std::println("f1(): {}", f1());
    std::println("f2(): {}", *f2());
    return 0;
}