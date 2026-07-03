#include <iostream>

#include <QCoreApplication>

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    auto result = a.exec();
    std::cout << result << std::endl;

    return 0;
}
