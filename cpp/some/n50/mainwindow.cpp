#include <QDebug>
#include <QWheelEvent>
#include <QGraphicsRectItem>

#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , m_ui(new Ui::MainWindow)
    , m_scene(new SceneWithSquares(10, 10, 25, this))
{
    m_ui->setupUi(this);
    m_ui->graphicsView
            ->setScene(m_scene);

    m_ui->graphicsView->setVerticalScrollBarPolicy(Qt::ScrollBarAlwaysOff);
    m_ui->graphicsView->setHorizontalScrollBarPolicy(Qt::ScrollBarAlwaysOff);
}

MainWindow::~MainWindow()
{
    delete m_ui;
}

void MainWindow::wheelEvent(QWheelEvent *pWheelEvent)
{
    qDebug() << "angleDelta:" << pWheelEvent->angleDelta();
    qDebug() << "pixelDelta:" << pWheelEvent->pixelDelta();
    qDebug() << "phase:"      << pWheelEvent->phase();
    qDebug() << "position:"   << pWheelEvent->position();
    qDebug() << "viewport width:" << m_ui->graphicsView->viewport()->width() / 2.0;
    qDebug() << "viewport height" << m_ui->graphicsView->viewport()->height() / 2.0;

    // ***

    const double angle = pWheelEvent->angleDelta().y(); // 120, -120
    const double factor = qPow(1.00075, angle);

    const auto targetViewportPos = pWheelEvent->position();
    const auto targetScenePos = m_ui->graphicsView->mapToScene(targetViewportPos.toPoint());

    m_ui->graphicsView->centerOn(targetScenePos);
    m_ui->graphicsView->scale(factor, factor);

    const QPointF deltaViewportPos =
            targetViewportPos -
            QPointF(
                m_ui->graphicsView->viewport()->width() / 2.0,
                m_ui->graphicsView->viewport()->height() / 2.0
                );

    QPointF viewportCenter =
            m_ui->graphicsView->mapFromScene(targetScenePos) - deltaViewportPos;
    m_ui->graphicsView->centerOn(
                m_ui->graphicsView->mapToScene(
                    viewportCenter.toPoint()));

    return;
}

