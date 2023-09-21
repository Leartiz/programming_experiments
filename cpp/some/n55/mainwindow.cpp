#include <QDebug>

#include <QString>
#include <QTextStream>

#include "mainwindow.h"
#include "ui_mainwindow.h"

const QString MainWindow::ConfigKey::winPos{
    "WinPos"
};

// -----------------------------------------------------------------------

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , m_ui(new Ui::MainWindow)
{
    m_ui->setupUi(this);

    // ui
    {
        auto ok = connect(m_ui->pushBtnPrint, &QPushButton::clicked,
                          this, &MainWindow::onClicked_pushBtnPrint);
        Q_ASSERT(ok);

        ok = connect(m_ui->pushBtnSync, &QPushButton::clicked,
                     this, &MainWindow::onClicked_pushBtnSync);
    }

    // ***

    m_settings = new QSettings{
        "conf.ini", QSettings::IniFormat, this
    };

    moveIfNeed();
}

void MainWindow::onClicked_pushBtnPrint()
{
    QString settingsAsStr;
    QTextStream ts{ &settingsAsStr };

    const auto winPos = m_settings->value(ConfigKey::winPos).toPoint();
    ts << ConfigKey::winPos << ": "
       << "x = " << winPos.x() << ", "
       << "y = " << winPos.y() << "";

    m_ui->plainTextEdit->appendPlainText(
                settingsAsStr);
}

void MainWindow::onClicked_pushBtnSync()
{
    m_settings->sync();
    moveIfNeed();
}

void MainWindow::moveIfNeed()
{
    if (m_settings->contains(ConfigKey::winPos)) {
        move(m_settings->value(ConfigKey::winPos).toPoint());
    }
}

void MainWindow::closeEvent(QCloseEvent *event)
{
    m_settings->setValue(ConfigKey::winPos, pos());
    event->accept();
}

MainWindow::~MainWindow()
{
    delete m_ui;
}

