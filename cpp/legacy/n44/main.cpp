#include <QJsonObject>
#include <QJsonDocument>

#include <QDebug>

#include <QHash>
#include <QMultiHash>
#include <QString>

struct MyClass
{
    QString firstName;
    QString secondName;
};

QDebug operator<<(QDebug dbg, const MyClass& mc)
{
    return dbg << QJsonObject{
        { "firstName", mc.firstName },
        { "secondName", mc.secondName },
    };
}

bool operator==(const MyClass& mc1, const MyClass& mc2)
{
    return
            mc1.firstName == mc2.firstName &&
            mc1.secondName == mc2.secondName;
}

uint qHash(const MyClass& mc)
{
    return
            qHash(mc.firstName) ^
            qHash(mc.secondName);
}

// -----------------------------------------------------------------------

int main()
{
    {
        QHash<QString, MyClass> hash;
        hash.insert("0", MyClass{ "first", "second" });
        hash.insert("1", MyClass{ "first1", "second1" });
        hash.insert("2", MyClass{ "first2", "second2" });
        qDebug() << hash["0"];
        qDebug() << hash["1"];
        qDebug() << hash["3"];
    }

    qDebug() << "---";

    {
        QHash<MyClass, QString> hash;
        hash[MyClass{ "first", "second" }]   = "412453";
        hash[MyClass{ "first1", "second1" }] = "756341";

        qDebug() << hash[MyClass{ "first", "second" }];
        qDebug() << hash[MyClass{ "first1", "second1" }];
    }

    qDebug() << "---";

    {
        QMultiHash<MyClass, QString> mhash;
        mhash.insert(MyClass{ "first", "second" }, "412453");
        mhash.insert(MyClass{ "first", "second" }, "241424");
        mhash.insert(MyClass{ "first", "second" }, "532563");

        // doesn't work as expected!
        mhash[MyClass{ "first1", "second1" }] = "923919";
        mhash[MyClass{ "first1", "second1" }] = "988888";

        mhash[MyClass{ "first2", "second2" }] = "111111";

        qDebug() << mhash[MyClass{ "first", "second" }];
        qDebug() << mhash[MyClass{ "first1", "second1" }];
        qDebug() << mhash[MyClass{ "first2", "second2" }];

        qDebug() << "***";

        qDebug() << mhash.values(MyClass{ "first", "second" });
        qDebug() << mhash.values(MyClass{ "first1", "second1" });
        qDebug() << mhash.values(MyClass{ "first2", "second2" });

        qDebug() << "***";

        for (
            auto it = mhash.begin();
            it != mhash.end();
            ++it) {

            if (it.key() == MyClass{ "first", "second" }) {
                qDebug() << it.value();
            }
        }
    }

    qDebug() << "---";

    // use multi hash as multi map...
    {
        QMultiHash<QString, QString> mapPhonebook;
        mapPhonebook.insert("Kermit", "+49 631322181");
        mapPhonebook.insert("Gonzo",  "+49 631322186");
        mapPhonebook.insert("Gonzo",  "+49 631322000");
        mapPhonebook.insert("Gonzo",  "+49 631322010");
        mapPhonebook.insert("Piggy",  "+49 631322187");
        mapPhonebook.insert("Piggy",  "+49 631322999");

        QMultiHash<QString, QString>::iterator it{
            mapPhonebook.find("Piggy")
        };
        for (; it != mapPhonebook.end(); ++it) {
            if (it.key() == "Piggy")
                qDebug() << it.value();
        }

        qDebug() << "***";

        qDebug() << mapPhonebook.keys();
        qDebug() << mapPhonebook.values();
    }

    return 0;
}
