#include <QPushButton>

#include <QString>
#include <QDateTime>

#include <QThread>
#include <QProcess>
#include <QtConcurrent/QtConcurrent>

#include "mainwindow.h"
#include "ui_mainwindow.h"

namespace
{

struct DirHashResult
{
    QString hash;
    QString error;
};

} // <anonymous>

// -----------------------------------------------------------------------

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    ui->plainTextEditLog->setReadOnly(true);

    // ui
    {
        connect(ui->pushBtnRun, &QPushButton::clicked,
                this, &MainWindow::onClicked_pushBtnRun);
    }
}

MainWindow::~MainWindow()
{
    delete ui;
}

// -----------------------------------------------------------------------

void MainWindow::onClicked_pushBtnRun()
{
    auto task = [this](QPromise<DirHashResult>& promise) -> void {

        promise.start();

        // ***

        QProcess dirHashProcess;
        dirHashProcess.start("../../dependencies/builded/DirHash.exe", {
                                 ui->lineEditPath->text(),
                                 "-quiet", "sha1", "-nowait"
                             });

        QThread::msleep(2500); // imitation...

        DirHashResult result;
        if (!dirHashProcess.waitForFinished(5000)) {
            result.error =
                    "no response received";
        }
        if (dirHashProcess.exitCode() != 0) {
            result.error =
                    "error response";
        }

        // ***

        result.hash =
                dirHashProcess.readAllStandardOutput().trimmed();

        promise.addResult(
                    std::move(
                        result));

        // ***

        promise.finish();
    };
    auto future = QtConcurrent::task(std::move(task))
            .spawn();

    // ***

    const auto valWatcher = new QFutureWatcher<DirHashResult>{ this };
    valWatcher->setFuture(future);

    // ***

    connect(valWatcher, &QFutureWatcher<DirHashResult>::finished,
            this, [this]() {

        const auto fw = static_cast<QFutureWatcher<DirHashResult>*>(sender());
        const auto f = fw->future();

        // ***

        const auto result = f.result();
        const QString msg = {
            QDateTime::currentDateTime().toString("hh:mm:ss.z") + " | " +
            (result.error.isEmpty() ? result.hash : result.error)
        };

        // ***

        ui->plainTextEditLog->appendPlainText(msg);
        fw->deleteLater();
    });
}

