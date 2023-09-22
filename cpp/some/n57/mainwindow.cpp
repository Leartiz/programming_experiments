#include <QRegularExpression>
#include <QRegularExpressionMatch>

#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow)
{
    ui->setupUi(this);

    // ui
    {
        connect(ui->pushBtnCheck, &QPushButton::clicked,
                this, &MainWindow::onClicked_pushBtnCheck);
    }
}

void MainWindow::onClicked_pushBtnCheck()
{
    QRegularExpression regExp{ ui->lineEditRegExp->text() };
    auto result = regExp.match(ui->lineEditValue->text());

    // ***

    QString buffer = ui->plainTextEditLog->toPlainText();
    ui->plainTextEditLog->clear();

    if (result.hasMatch()) {
        buffer.append(".");
    }
    else {
        buffer.append("F");
    }

    ui->plainTextEditLog
            ->appendPlainText(buffer);
}

MainWindow::~MainWindow()
{
    delete ui;
}

