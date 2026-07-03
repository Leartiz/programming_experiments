#include <QDebug>
#include <QEvent>

#include "myprogressbar.h"
#include "progressevent.h"

MyProgressBar::MyProgressBar(QWidget* wgt)
    : QProgressBar{ wgt } {}

void MyProgressBar::customEvent(QEvent *event)
{
    if (event->type() == ProgressEvent::Progress) {
        /*
        qDebug() << "MyProgressBar, customEvent, "
                    "event type is progress";
        */

        // ***

        const auto pe =
                static_cast<ProgressEvent*>(event);
        setValue(pe->value());
    }
    QWidget::customEvent(event);
}
