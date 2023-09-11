#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

#include <QGraphicsView>
#include <QGraphicsScene>
#include <QGraphicsItem>

#include <QGraphicsLineItem>
#include <QGraphicsRectItem>
#include <QGraphicsEllipseItem>
#include <QGraphicsPixmapItem>

#include <QTimer>

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class Spaceship : public QGraphicsPixmapItem
{
public:
    explicit Spaceship(int scHeight);
};

class Asteroid : public QGraphicsPixmapItem
{
public:
    explicit Asteroid(int x);

    // QGraphicsItem interface
public:
    void advance(int phase) override;
    void mousePressEvent(QGraphicsSceneMouseEvent*) override;

private:
    int m_ySpeed{ 2 };
};

// -----------------------------------------------------------------------

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow();

private slots:
    void onTimeout_genTimer();

private:
    Ui::MainWindow *ui;
    QGraphicsScene *m_scene;

private:
    QTimer* m_animTimer;
    QTimer* m_genTimer;
};

#endif // MAINWINDOW_H
