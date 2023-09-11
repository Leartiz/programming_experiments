#include <QtWidgets>

int main(int argc, char** argv)
{
    QApplication app{ argc, argv };

    QGraphicsScene scene(QRectF(-100, -100, 300, 300));
    QGraphicsView view(&scene);

    QGraphicsRectItem* pRectItem =
            scene.addRect(
                QRectF(-30, -30, 120, 80),
                QPen(Qt::black),
                QBrush(Qt::green)
                );
    pRectItem->setFlags(QGraphicsItem::ItemIsMovable);

    view.show();
    return app.exec();
}
