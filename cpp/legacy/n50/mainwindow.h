#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

#include "squaresscene.h"

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow() override;

    // QWidget interface
protected:
    void wheelEvent(QWheelEvent *event) override;

private:
    Ui::MainWindow* m_ui;
    SceneWithSquares* m_scene;
};
#endif // MAINWINDOW_H
