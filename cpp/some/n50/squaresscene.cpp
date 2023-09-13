#include "squaresscene.h"

SceneWithSquares::SceneWithSquares(int w, int h, int sideSize, QObject *parent)
    : QGraphicsScene(parent)
{
    m_matrix.resize(w, Row(h, 0));

    // ***

    setSceneRect({ 0, 0,
                   qreal(w * sideSize),  qreal(h * sideSize) });

    // ***

    int x = 0, y = 0;
    for (int i = 0; i < m_matrix.size(); ++i) {
        for (int j = 0; j < m_matrix[0].size(); ++j) {
            addRect(x, y, sideSize, sideSize);
            x += sideSize;
        }

        y += sideSize;
        x = 0;
    }
}

void SceneWithSquares::mousePressEvent(QGraphicsSceneMouseEvent *event)
{

}

void SceneWithSquares::mouseReleaseEvent(QGraphicsSceneMouseEvent *event)
{

}
