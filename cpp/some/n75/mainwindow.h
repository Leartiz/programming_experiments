#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QVector>
#include <QMainWindow>
#include <QWebSocketServer>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

    static constexpr char svrName[] = "test_svr";
    static const quint16 port = 34343;
    static const int thCount = 2;

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow() override;

signals:
    void newSock(int, QWebSocket*);

private slots:
    void onNewConnection_svr();
    void onMessageProcessed_worker(const QString&);

private:
    void initWorkers();

private:
    Ui::MainWindow *m_ui;

private:
    QWebSocketServer* m_svr;
    QVector<QThread*> m_workerThs;
};

#endif // MAINWINDOW_H
