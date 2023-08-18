#include <QDebug>
#include <QApplication>

#include <QWidget>
#include <QVBoxLayout>
#include <QPushButton>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    auto btn = new QPushButton{ "Click Me" };
    QObject::connect(btn, &QPushButton::clicked, []() {
        qFatal("something broke...");
    });

    // ***

    auto vl = new QVBoxLayout{};
    vl->addWidget(btn);

    QWidget wgt;
    wgt.setLayout(vl);
    wgt.resize(250, 100);

    // ***

    wgt.show();
    return a.exec();
}
