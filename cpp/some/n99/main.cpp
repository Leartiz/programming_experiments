#include <iostream>
#include <windows.h>

std::string getSystemCodepage() // 1251
{
#ifdef Q_OS_WIN32
    return "CP" + QString::number(GetACP());
#endif
    return "UTF-8";
}

std::string getConsoleCodepage() // 886
{
#ifdef Q_OS_WIN32
    return "CP" + QString::number(GetOEMCP());
#endif
    return "UTF-8";
}

int main()
{
    std::cout << GetConsoleCP() << "\n";
    std::cout << GetACP() << "\n";
    std::cout << GetOEMCP() << "\n";

    std::cout << "Console Codepage: " << getConsoleCodepage() << "\n";
    std::cout << "System Codepage: " << getSystemCodepage() << "\n";
}
