#ifndef MYPROGRESSBAR_H
#define MYPROGRESSBAR_H

#include <QWidget>
#include <QProgressBar>

class MyProgressBar :
        public QProgressBar
{
    Q_OBJECT

public:
    explicit MyProgressBar(
            QWidget* wgt = nullptr);

    // QObject interface
protected:
    void customEvent(QEvent *event);
};

#endif // MYPROGRESSBAR_H
