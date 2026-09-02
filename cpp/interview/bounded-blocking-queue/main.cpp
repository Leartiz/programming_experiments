#include "bounded_blocking_queue.hpp"

#include <iostream>
#include <string>

int main()
{
    BoundedBlockingQueue<std::string> q(2);

    q.push("a");
    q.push("b");

    std::string item;
    q.pop(item);
    std::cout << item << '\n';
}
