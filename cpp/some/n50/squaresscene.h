#ifndef SQUARESSCENE_H
#define SQUARESSCENE_H

#include <QObject>
#include <QGraphicsScene>
#include <QVector>

class SceneWithSquares : public QGraphicsScene
{
    Q_OBJECT
public:
    SceneWithSquares(int w, int h, int sideSize, QObject *parent = nullptr);

    // QGraphicsScene interface
protected:
    void mousePressEvent(QGraphicsSceneMouseEvent *event);
    void mouseReleaseEvent(QGraphicsSceneMouseEvent *event);

private:
    using Row = QVector<int>;
    using Matrix = QVector<Row>;

    Matrix m_matrix;
};

#endif // SQUARESSCENE_H
