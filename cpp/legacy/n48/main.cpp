#include <QtWidgets>

int main(int argc, char** argv)
{
    QApplication app{ argc, argv };

    QGraphicsScene scene(QRectF(-100, -100, 300, 300));
    QGraphicsView view(&scene);

    // ***

    QGraphicsRectItem* pRectItem =
            scene.addRect(
                QRectF(-30, -30, 120, 80),
                QPen(Qt::black),
                QBrush(Qt::green)
                );
    pRectItem->setFlag(QGraphicsItem::ItemIsMovable);

    // *** parent-child ***

    QGraphicsTextItem* pTextItem =
            scene.addText("Moveee");
    pTextItem->setFlag(QGraphicsItem::ItemIsMovable);


    QGraphicsLineItem* pLineItem =
            scene.addLine(-10, -10, 80, 80, QPen(Qt::red, 2));
    pLineItem->setFlag(QGraphicsItem::ItemIsMovable);

    pTextItem->setParentItem(pLineItem);

    // *** transform ***

    QGraphicsTextItem* pTextItem2 = scene.addText("Shear");
    pTextItem2->setTransform(
                QTransform().shear(2, 0.0), true);
    pTextItem2->setFlag(QGraphicsItem::ItemIsMovable);

    // ***

    view.show();
    return app.exec();
}
