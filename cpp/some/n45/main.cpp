#include <algorithm>

#include <QDebug>
#include <QString>

#include "hashmap.h"

bool tests()
{
    return true;
}

int main()
{
    if (!tests()) {
        return 1;
    }

    qDebug() << "---";

    qDebug() << lez::hashFromKey("123");
    qDebug() << lez::hashFromKey("1");
    qDebug() << lez::hashFromKey("2");
    qDebug() << lez::hashFromKey("3");

    qDebug() << "--- insert ---";

    lez::HashMap<int, QString> hh;
    for (int i = 0; i < 25; ++i) {
        hh.insert(i, "123");
    }

    qDebug() << "---";

    const auto lines = hh.toString().split('\n', Qt::SkipEmptyParts);
    for (auto it = lines.begin(); it != lines.end(); ++it) {
        qDebug() << *it;
    }

    qDebug() << "---";

    auto keys = hh.keys();
    qDebug() << keys;
    std::sort(keys.begin(),
              keys.end());
    qDebug() << keys;

    qDebug() << "---";

    qDebug() << hh.vals();
    qDebug() << hh.hashVals();

    return 0;
}
