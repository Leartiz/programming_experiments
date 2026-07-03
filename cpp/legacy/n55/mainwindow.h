#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QSettings>
#include <QMainWindow>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

private:
    struct ConfigKey final
    {
        static const QString winPos;
    };

public:
    MainWindow(QWidget *parent = nullptr);
    ~MainWindow() override;

private slots:
    void onClicked_pushBtnPrint();
    void onClicked_pushBtnSync();

    // QWidget interface
protected:
    void closeEvent(QCloseEvent *event) override;

private:
    void moveIfNeed();

private:
    Ui::MainWindow *m_ui;
    QSettings *m_settings;
};

#endif // MAINWINDOW_H
