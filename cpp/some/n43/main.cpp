#include <QMap>
#include <QMultiMap>
#include <QString>
#include <QDebug>

int main()
{
    QMap<QString, QString> map;
    QMultiMap<QString, QString> mmap;

    {
        map = {
            { "1", "blpAk" },
            { "1", "nJmeg" },
            { "1", "51e4H" },

            { "3", "sk4Z5" },
            { "4", "KLEq3" },
            { "5", "9lWrv" },
        };

        // returns an iterator pointing
        // to the first item with key key in the map.
        // If the map contains no item with key key,
        // the function returns an iterator
        // to the nearest item with a greater key.
        qDebug() << *map.lowerBound("1");

        // returns an iterator pointing to the item that immediately
        // follows the last item with key key in the map.
        // If the map contains no item with key key,
        // the function returns an iterator
        // to the nearest item with a greater key.
        qDebug() << *map.upperBound("1");

        qDebug() << *map.lowerBound("2");
        qDebug() << *map.upperBound("2");

        qDebug() << map.value("1");
    }

    qDebug() << "---";

    {
        mmap = {
            { "1", "blpAk" },
            { "1", "nJmeg" },
            { "1", "51e4H" },

            { "3", "sk4Z5" },
            { "4", "KLEq3" },
            { "5", "9lWrv" },
        };

        qDebug() << *mmap.lowerBound("1");
        qDebug() << *mmap.upperBound("1");

        qDebug() << *mmap.lowerBound("2");
        qDebug() << *mmap.upperBound("2");

        qDebug() << mmap.values("1");
    }

    qDebug() << "---";

    {
        QMultiMap<QString, QString> mapPhonebook;
        mapPhonebook.insert("Kermit", "+49 631322181");
        mapPhonebook.insert("Gonzo",  "+49 631322186");
        mapPhonebook.insert("Gonzo",  "+49 631322000");
        mapPhonebook.insert("Gonzo",  "+49 631322010");
        mapPhonebook.insert("Piggy",  "+49 631322187");
        mapPhonebook.insert("Piggy",  "+49 631322999");

        QMultiMap<QString, QString>::iterator it{
            mapPhonebook.find("Piggy")
        };
        for (; it != mapPhonebook.end() && it.key() == "Piggy"; ++it) {
            qDebug() << it.value();
        }

        qDebug() << "***";

        qDebug() << mapPhonebook.keys();
        qDebug() << mapPhonebook.values();
    }

    return 0;
}
