#include <QColor>

#include "mainwindow.h"
#include "ui_mainwindow.h"

Spaceship::Spaceship(int scHeight)
    : QGraphicsPixmapItem(0)
{
    setPixmap(QPixmap(":/resources/spaceship.png"));
    setScale(0.15);

    // ***

    setPos(0, scHeight -
                  (pixmap().height() * scale()));
}

// -----------------------------------------------------------------------

Asteroid::Asteroid(int x)
    : QGraphicsPixmapItem(0)
{
    setPixmap(QPixmap(":/resources/asteroid.png"));
    setPos(rand() % (x - pixmap().width()), 0);
    setScale(0.15);

    // ***

    setFlag(QGraphicsItem::ItemIsMovable, true);
}

void Asteroid::advance(int phase)
{
    if (phase == 0) {
        return;
    }

    moveBy(0, m_ySpeed);
    //setRotation(rotation() + 10);
}

void Asteroid::mousePressEvent(QGraphicsSceneMouseEvent*)
{
    delete this;
}

// -----------------------------------------------------------------------

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    m_scene = new QGraphicsScene{ { 0, 0, 800, 600 }, this };

    m_scene->addRect(m_scene->sceneRect());
    ui->graphicsView->setScene(m_scene);

    // ***

    m_scene->addItem(
        new Spaceship(
            m_scene->height()));

    // ***

    m_animTimer = new QTimer(this);
    connect(m_animTimer, &QTimer::timeout,
            m_scene, &QGraphicsScene::advance);

    m_animTimer->start(1000 / 60);

    // ***

    m_genTimer = new QTimer(this);
    connect(m_genTimer, &QTimer::timeout,
            this, &MainWindow::onTimeout_genTimer);
    m_genTimer->start(1000);
}

void MainWindow::onTimeout_genTimer()
{
    m_scene->addItem(
        new Asteroid(
            m_scene->width()));
}

MainWindow::~MainWindow()
{
    delete ui;
}
