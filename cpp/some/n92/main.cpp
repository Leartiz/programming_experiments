#include <QDebug>
#include <QtConcurrent/QtConcurrent>
#include <QFutureWatcher>

QFuture<QString> foo() {
    QPromise<QString> p;
    p.start();
    p.addResult("123");
    p.finish();
    return p.future();
}

int main(int argc, char** argv) {
    QCoreApplication app(argc, argv);

    auto fw = new QFutureWatcher<QString>();

    fw->setFuture(foo());
    QObject::connect(fw, &QFutureWatcher<QString>::finished,
                     fw, [fw](){
        qDebug() << ".";
    });
    return app.exec();
}
