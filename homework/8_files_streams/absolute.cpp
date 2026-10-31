#include <filesystem>
#include <iostream>

namespace fs = std::filesystem;

int main() {
    // Does not skip "newDirectory/"
    for (auto entry : fs::directory_iterator("newDirectory"))
        std::cout << fs::absolute(entry.path()) << '\n';
}
