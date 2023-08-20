#include <QLineEdit>
#include <QTextEdit>
#include <QPushButton>

#include <QThread>
#include <QMetaEnum>

#include <QTextCodec>
#include <QMessageBox>
#include <QDebug>

#include "mainwindow.h"
#include "ui_mainwindow.h"

namespace
{

// ...

} // <anonymous>

const char* MainWindow::cmdCodecName{
    "IBM 866"
};

// -----------------------------------------------------------------------

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow{ parent }
    , m_ui{ new Ui::MainWindow }

    , m_process{ new QProcess{ this } }
    , m_timer{ new QTimer{ this } }
{
    m_ui->setupUi(this);

    // ui
    {
        connect(m_ui->pushBtnGo, &QPushButton::clicked,
                this, &MainWindow::onClicked_pushBtnGo);
        connect(m_ui->lineEditCmd, &QLineEdit::returnPressed,
                this, &MainWindow::onClicked_pushBtnGo); // same action!

        connect(m_ui->lineEditCmd, &QLineEdit::returnPressed,
                m_ui->textEditResult, &QTextEdit::clear);
        connect(m_ui->pushBtnGo, &QPushButton::clicked,
                m_ui->textEditResult, &QTextEdit::clear);
    }

    // process
    {
        connect(m_process, &QProcess::readyReadStandardOutput,
                this, &MainWindow::onReadyReadStandardOutput_process);
        connect(m_process, &QProcess::readyReadStandardError,
                this, &MainWindow::onReadyReadStandardError_process);
    }

    // timer
    m_timer->setInterval(1000);
    m_timer->setSingleShot(true);
    {
        connect(m_timer, &QTimer::timeout,
                this, &MainWindow::onTimeout_timer);
    }

    // *** works?

    QThread::currentThread()->setPriority(
        QThread::HighPriority);
    qDebug() << "Current thread priority:"
             << QThread::currentThread()
                    ->priority();
}

MainWindow::~MainWindow()
{
    delete m_ui;
}

// -----------------------------------------------------------------------

void MainWindow::onClicked_pushBtnGo()
{
    if (m_timer->isActive()) {
        QMessageBox::warning(this, "Warn", "Still in progress...");
        return;
    }

    m_timer->start();
    m_process->start("cmd", {
                                "/C",
                                m_ui->lineEditCmd->text() // as args
                     });
}

// -----------------------------------------------------------------------

void MainWindow::onReadyReadStandardOutput_process()
{
    m_timer->stop();
    const QByteArray bytes = m_process->readAllStandardOutput();
    const auto textCodec = QTextCodec::codecForName(cmdCodecName);
    if (!textCodec) {
        qWarning() << "MainWindow, onReadyReadStandardOutput_process, "
                      "text codec is null.";
        return;
    }

    const QString message = {
        "Output:\n" + textCodec->toUnicode(bytes)
    }; m_ui->textEditResult->append(
                message);
}

void MainWindow::onReadyReadStandardError_process()
{
    m_timer->stop();
    const QByteArray bytes = m_process->readAllStandardError();
    const auto textCodec = QTextCodec::codecForName(cmdCodecName);
    if (!textCodec) {
        qWarning() << "MainWindow, onReadyReadStandardError_process, "
                      "text codec is null.";
        return;
    }

    const QString message = {
        "Error:\n" + textCodec->toUnicode(bytes)
    }; m_ui->textEditResult->append(
                message);
}

// -----------------------------------------------------------------------

void MainWindow::onTimeout_timer()
{
    m_process->kill();
    const QString message = {
        "No response for a long time, "
        "the process is killed..."
    };

    // ***

    m_ui->textEditResult->clear();
    m_ui->textEditResult->append(
                message);
}
