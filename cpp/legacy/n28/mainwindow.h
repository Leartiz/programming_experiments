#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QTimer>
#include <QProcess>
#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

    static const char* cmdCodecName;

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow() override;

private slots:
    void onClicked_pushBtnGo();

private slots:
    void onReadyReadStandardOutput_process();
    void onReadyReadStandardError_process();
    void onTimeout_timer();

private:
    Ui::MainWindow *m_ui{ nullptr };
    QProcess* m_process{ nullptr };
    QTimer* m_timer{ nullptr };
};

#endif // MAINWINDOW_H
