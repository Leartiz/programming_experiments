#include <stdexcept>
#include <algorithm>

#include <QThread>
#include <QWebSocket>

#include "worker.h"
#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , m_ui(new Ui::MainWindow)
    , m_svr{
          new QWebSocketServer(
              svrName, QWebSocketServer::NonSecureMode, this
              ),
      }
{
    m_ui->setupUi(this);

    // svr
    {
        connect(m_svr, &QWebSocketServer::newConnection,
                this, &MainWindow::onNewConnection_svr);
    }
    initWorkers();

    if (!m_svr->listen(QHostAddress::Any, port)) {
        throw std::runtime_error("server start failed");
    }
}

void MainWindow::initWorkers()
{
    for (int i = 0; i < thCount; ++i) {
        auto th = new QThread{ this };
        auto worker = new Worker{ i, nullptr };
        {
            QObject::connect(th, &QThread::started, worker, &Worker::almostStart);
            QObject::connect(th, &QThread::finished, worker, &Worker::almostFinish);

            // ***

            QObject::connect(this, &MainWindow::newSock,
                             worker, &Worker::pushSock,
                             Qt::QueuedConnection);
            QObject::connect(worker, &Worker::messageProcessed,
                             this, &MainWindow::onMessageProcessed_worker,
                             Qt::QueuedConnection);
        }

        // ***

        m_workerThs.push_back(th);
        worker->moveToThread(th);
        th->start();
    }
}

void MainWindow::onNewConnection_svr()
{
    QWebSocket* wsk = m_svr->nextPendingConnection();

    wsk->moveToThread(nullptr);
    wsk->setParent(nullptr);

    const auto index = std::rand() % m_workerThs.size();
    emit newSock(index, wsk);
}

void MainWindow::onMessageProcessed_worker(const QString& message)
{
    m_ui->plainTextLog->appendPlainText(message);
}

MainWindow::~MainWindow()
{
    std::for_each(m_workerThs.begin(), m_workerThs.end(), [](QThread* th) {
        th->quit();
        th->wait();
    });
    delete m_ui;
}

