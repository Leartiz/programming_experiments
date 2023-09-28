#include <QDebug>
#include <QSysInfo>
#include <QHostInfo>
#include <QNetworkInterface>

int main()
{
    qDebug() << "bootUniqueId:" << QSysInfo::bootUniqueId();
    qDebug() << "buildAbi:" << QSysInfo::buildAbi();
    qDebug() << "buildCpuArchitecture:" << QSysInfo::buildCpuArchitecture();
    qDebug() << "currentCpuArchitecture:" << QSysInfo::currentCpuArchitecture();
    qDebug() << "kernelType:" << QSysInfo::kernelType();
    qDebug() << "kernelVersion:" << QSysInfo::kernelVersion();
    qDebug() << "machineHostName:" << QSysInfo::machineHostName();
    qDebug() << "machineUniqueId:" << QSysInfo::machineUniqueId();
    qDebug() << "prettyProductName:" << QSysInfo::prettyProductName();
    qDebug() << "productType:" << QSysInfo::productType();
    qDebug() << "productVersion:" << QSysInfo::productVersion();

    // ***

    qDebug();
    qDebug() << "hostName:" << QHostInfo{}.hostName();
    qDebug() << "localDomainName:" << QHostInfo::localDomainName();
    qDebug() << "localHostName:" << QHostInfo::localHostName();

    // ***

    qDebug();
    qDebug() << "hardwareAddress:" << QNetworkInterface{}.hardwareAddress();
    qDebug() << "allAddresses:" << QNetworkInterface::allAddresses();
    qDebug() << "allInterfaces:" << QNetworkInterface::allInterfaces();

    return 0;
}
