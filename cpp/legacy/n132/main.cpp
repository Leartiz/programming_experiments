#include <iostream>
#include <string>
#include <functional>
#include <cstdint>

struct Message final
{   
public:
    int command = 0;
    std::string header;
    std::string body;

public:
    void println() {
        const auto indent = "   ";
        std::cout << "msg: {\n";
        std::cout << indent << "command: " << command << ",\n";
        std::cout << indent << "header: "  << header << ",\n";
        std::cout << indent << "body: "    << body << "\n";
        std::cout << "}" << std::endl;
    }
};

using CommandId = uint8_t;
using MessageHandler = std::function<void(Message &&)>;
using CommandHandlers = std::unordered_map<CommandId, MessageHandler>;

// -----------------------------------------------------------------------

void handler(Message && msg) 
{
    std::cout << "*** free function ***" << std::endl;
    msg.println();
    std::cout << std::endl;
}

class Handler final
{
public:
    void method(Message && msg) {
        std::cout << "*** member function ***" << std::endl;
        msg.println();
        std::cout << std::endl;
    }
};

// -----------------------------------------------------------------------

int main() 
{
    Handler h;
    const auto ch = CommandHandlers{
        { 
            1, 
            [](Message && msg){
                std::cout << "*** lambda function ***" << std::endl;
                msg.println();
                std::cout << std::endl;
            }
        },
        {
            2,
            std::bind(&Handler::method, h, std::placeholders::_1)
        },
        {
            3,
            handler
        },
        //...
    };

    {
        Message msg{ 1, "111", "111" };
        ch.at(msg.command)(std::move(msg));
    }    
    {
        Message msg{ 2, "222", "222" };
        ch.at(msg.command)(std::move(msg));
    }
        {
        Message msg{ 3, "333", "333" };
        ch.at(msg.command)(std::move(msg));
    }
}