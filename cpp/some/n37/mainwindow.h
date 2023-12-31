#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow final :
        public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(
        QWidget *parent = nullptr);
    ~MainWindow() override;

public slots:
    void onClicked_pushBtnRun();

private:
    Ui::MainWindow *ui{
        nullptr
    };
};
#endif // MAINWINDOW_H
