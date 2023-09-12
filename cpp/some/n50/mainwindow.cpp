#include <QGraphicsRectItem>
#include <QWheelEvent>

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
}

MainWindow::~MainWindow()
{
    delete m_ui;
}

void MainWindow::wheelEvent(QWheelEvent *pWheelEvent)
{
    // Do a wheel-based zoom about the cursor position
    double angle = pWheelEvent->angleDelta().y();
    double factor = qPow(1.0015, angle);

    auto targetViewportPos = pWheelEvent->position();
    auto targetScenePos = m_ui->graphicsView->mapToScene(
                pWheelEvent->position().toPoint());

    m_ui->graphicsView->scale(factor, factor);
    m_ui->graphicsView->centerOn(targetScenePos);
    QPointF deltaViewportPos = targetViewportPos - QPointF(m_ui->graphicsView->viewport()->width() / 2.0, m_ui->graphicsView->viewport()->height() / 2.0);
    QPointF viewportCenter = m_ui->graphicsView->mapFromScene(targetScenePos) - deltaViewportPos;
    m_ui->graphicsView->centerOn(m_ui->graphicsView->mapToScene(viewportCenter.toPoint()));

    return;
}

